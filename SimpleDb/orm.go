package SimpleDb

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
)

type TableName string

var typeTableName TableName
var tableNameType reflect.Type = reflect.TypeOf(typeTableName)

type TableInfo struct {
	Name   string      //表名
	Fields []FieldInfo //字段
}

type FieldInfo struct {
	Name           string //字段名
	IsPrimaryKey   bool   //是否主键
	IsAutoGenerate bool   //是否自动生成(增长)
	Value          reflect.Value
}

func getTableInfo(model interface{}) (tbinfo *TableInfo, err error) {
	defer func() {
		if e := recover(); e != nil {
			tbinfo = nil
			err = e.(error)
		}
	}()
	err = nil
	tbinfo = &TableInfo{}
	rt := reflect.TypeOf(model)
	rv := reflect.ValueOf(model)
	//默认是结构体名
	tbinfo.Name = rt.Name()
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		rv = rv.Elem()
	}
	for i, j := 0, rt.NumField(); i < j; i++ {
		rtf := rt.Field(i)
		rvf := rv.Field(i)
		//tableNameType类型，只是用该字段的tag来设置表名
		if rtf.Type == tableNameType {
			tbinfo.Name = string(rtf.Tag)
			continue
		}
		if rtf.Tag == "-" {
			continue
		}
		//如果字段没有tag
		var f FieldInfo
		if rtf.Tag == "" {
			f = FieldInfo{Name: rtf.Name, IsPrimaryKey: false, IsAutoGenerate: false, Value: rvf}

		} else {
			//判断tag中有没有:有的话，说明设置了主键自增等参数，否则认为tag中存的是数据库的字段名
			strTag := string(rtf.Tag)

			if strings.Index(strTag, ":") == -1 {
				f = FieldInfo{Name: strings.TrimSpace(strTag), IsPrimaryKey: false, IsAutoGenerate: false, Value: rvf}
			} else {
				//取字段名
				strName := rtf.Tag.Get("name")
				if strName == "" {
					strName = rtf.Name
				}
				//取主键
				isPk := false
				str := rtf.Tag.Get("PK")
				if str == "true" {
					isPk = true
				}
				str = rtf.Tag.Get("Auto")
				//获取是否自增的值
				isAuto := false
				if str == "true" {
					isAuto = true
				}
				f = FieldInfo{Name: strName, IsPrimaryKey: isPk, IsAutoGenerate: isAuto, Value: rvf}
			}
		}
		tbinfo.Fields = append(tbinfo.Fields, f)
	}
	return
}

//生成插入的SQL语句，和对应的参数
func generateInsertSql(model interface{}) (string, []interface{}, *TableInfo, error) {
	tbinfo, err := getTableInfo(model)
	if err != nil {
		return "", nil, nil, err
	}

	//如果结构体中没有字段，抛出异常
	if len(tbinfo.Fields) == 0 {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有字段")
	}
	strSql := "insert into " + tbinfo.Name
	strField := ""
	strValue := ""
	var param []interface{}
	for _, v := range tbinfo.Fields {
		if v.IsAutoGenerate { //跳过自动增长的自段
			continue
		}
		strField += v.Name + ","
		strValue += "?,"
		param = append(param, v.Value.Interface())
	}
	if strField == "" {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有字段,或只有自增自段")
	}
	strField = strings.TrimRight(strField, ",")
	strValue = strings.TrimRight(strValue, ",")
	strSql += " (" + strField + ") values(" + strValue + ")"
	return strSql, param, tbinfo, nil
}

//生成修改的Sql语句,以主键做为修改的条件
func generateUpdateSql(model interface{}) (string, []interface{}, *TableInfo, error) {
	tbinfo, err := getTableInfo(model)
	if err != nil {
		return "", nil, nil, err
	}

	//如果结构体中没有字段，抛出异常
	if len(tbinfo.Fields) == 0 {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有字段")
	}
	strSql := "update " + tbinfo.Name
	strField := ""
	strWhere := ""
	var param []interface{}
	var whereParam []interface{}
	for _, v := range tbinfo.Fields {
		if v.IsPrimaryKey { //主键，做为修改的条件
			strWhere += v.Name + "=? and"
			whereParam = append(whereParam, v.Value.Interface())
			continue
		}
		if v.IsAutoGenerate { //如果是自动增长的，跳过
			continue
		}
		strField += v.Name + "=?,"
		param = append(param, v.Value.Interface())
	}

	if strField == "" {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有字段,或只有自增自段")
	}
	if strWhere == "" {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有设置主键，不能使用此方式进行修改")
	}
	param = append(param, whereParam...)
	strField = strings.TrimRight(strField, ",")
	strWhere = strings.TrimRight(strWhere, "and")
	strSql += " set " + strField + " where " + strWhere
	return strSql, param, tbinfo, nil
}

//生成加载Sql语句,以主键做为查询的条件
func generateLoadSql(model interface{}) (string, []interface{}, *TableInfo, error) {
	tbinfo, err := getTableInfo(model)
	if err != nil {
		return "", nil, nil, err
	}

	//如果结构体中没有字段，抛出异常
	if len(tbinfo.Fields) == 0 {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有字段")
	}
	strSql := "select * from " + tbinfo.Name
	strWhere := ""
	var param []interface{}
	for _, v := range tbinfo.Fields {
		if !v.IsPrimaryKey { //主键，做为查询的条件，如果不是主键继续
			continue
		}
		strWhere += v.Name + "=? and"
		param = append(param, v.Value.Interface())
	}
	if strWhere == "" {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有设置主键，不能使用此方式进行查询")
	}
	strWhere = strings.TrimRight(strWhere, "and")
	strSql += " where " + strWhere + " limit 1"
	return strSql, param, tbinfo, nil
}

//生成删除Sql语句,以主键做为删除的条件
func generateDeleteSql(model interface{}) (string, []interface{}, *TableInfo, error) {
	tbinfo, err := getTableInfo(model)
	if err != nil {
		return "", nil, nil, err
	}

	//如果结构体中没有字段，抛出异常
	if len(tbinfo.Fields) == 0 {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有字段")
	}
	strSql := "Delete from " + tbinfo.Name
	strWhere := ""
	var param []interface{}
	for _, v := range tbinfo.Fields {
		if !v.IsPrimaryKey { //主键，做为查询的条件，如果不是主键继续
			continue
		}
		strWhere += v.Name + "=? and"
		param = append(param, v.Value.Interface())
	}
	if strWhere == "" {
		return "", nil, nil, errors.New(tbinfo.Name + "结构体中没有设置主键，不能使用此方式进行删除")
	}
	strWhere = strings.TrimRight(strWhere, "and")
	strSql += " where " + strWhere
	return strSql, param, tbinfo, nil
}

//将插入时得到自动增长的ID赋值给Struct对应字段
func setAuto(result sql.Result, tabinfo *TableInfo) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	id, err := result.LastInsertId()
	if id == 0 {
		return
	}
	if err != nil {
		return
	}
	for _, v := range tabinfo.Fields {
		if v.IsAutoGenerate {
			v.Value.SetInt(id)
			break
		}
	}
	return
}

type IGetData interface {
	GetValue(name string, value interface{}) error
}

//取数据赋给对应的struct里的字段
func SetFieldValue(row IGetData, tabinfo *TableInfo) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	for _, field := range tabinfo.Fields {
		value := field.Value.Interface()
		row.GetValue(field.Name, &value)
		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		fieldValue := field.Value
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}
		switch value.(type) {
		case int, int8, int16, int32, int64:
			intValue := rv.Int()
			//取到的是int,struct字段类型是bool
			if field.Value.Kind() == reflect.Bool {
				//对于bool类型0是false，非0是true
				fieldValue.SetBool(intValue != 0)
			} else {
				fieldValue.SetInt(intValue)
			}
		case uint, uint8, uint16, uint32, uint64:
			intValue := rv.Uint()
			//取到的是int,struct字段类型是bool
			if fieldValue.Kind() == reflect.Bool {
				fieldValue.SetBool(intValue != 0)
			} else {
				fieldValue.SetUint(intValue)
			}
		case float32, float64:
			fieldValue.SetFloat(rv.Float())
		case string:
			switch fieldValue.Interface().(type) {
			case []byte: //取到的是string,struct字段类型是[]byte
				fieldValue.SetBytes([]byte(rv.String()))
			default:
				fieldValue.SetString(rv.String())
			}
		case []byte:
			switch fieldValue.Interface().(type) {
			//取到的是[]byte,字段类型是string
			case string:
				fieldValue.SetString(string(rv.Bytes()))
			default:
				fieldValue.SetBytes(rv.Bytes())
			}
		default:
			fieldValue.Set(rv)
		}
		//field.Value.Set(rv)
	}
	return
}

//设置model的值
func SetValue(row IGetData, model interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	var tabinfo *TableInfo
	tabinfo, err = getTableInfo(model)
	if err != nil {
		return
	}
	err = SetFieldValue(row, tabinfo)
	return
}
