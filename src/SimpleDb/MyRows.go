package SimpleDb

import (
	"database/sql"
	"errors"
	"reflect"
	"time"
)

type MyRows struct {
	*sql.Rows
	values     map[string]interface{}
	columnName []string
}

var ErrNoRows = errors.New("sql: no rows in result set")

/*根据字段名来取字段的值
name:字段名，区分大小写
value:用来接收字段值的变量，需传变量的地址,如&a
*/
func (this *MyRows) GetValue(name string, value interface{}) error {
	if this.values == nil || len(this.values) == 0 {
		return errors.New("没有调用Next,或没有可用的行")
	}
	i, ok := this.values[name]
	if ok {

		err := ConvertAssign(value, i)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("字段不存在，请注意大小写")
}
func (this *MyRows) Scan(dest ...interface{}) error {
	if this.values == nil || len(this.values) == 0 {
		return errors.New("没有调用Next,或没有可用的行")
	}

	for i := 0; i < len(dest); i++ {
		err := ConvertAssign(dest[i], this.values[this.columnName[i]])
		if err != nil {
			return err
		}
	}
	return nil

}
func (this *MyRows) Next() bool {
	bResult := this.Rows.Next()
	if bResult {
		//如果成功，取所有列的数据到values里
		if this.columnName == nil || len(this.columnName) == 0 {
			this.columnName, _ = this.Rows.Columns()
		}
		if this.values == nil {
			this.values = make(map[string]interface{})
		}
		var arr []interface{}
		for i := 0; i < len(this.columnName); i++ {
			var inf interface{}
			arr = append(arr, &inf)
		}
		//将数据接收到interface{}变量里
		this.Rows.Scan(arr...)

		for i := 0; i < len(this.columnName); i++ {
			this.values[this.columnName[i]] = reflect.ValueOf(arr[i]).Elem().Interface()
		}
	}
	return bResult
}

//取int类型的字段
func (this *MyRows) Int(strFieldName string) int {
	var intValue int
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取int8类型的字段
func (this *MyRows) Int8(strFieldName string) int8 {
	var intValue int8
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取int16类型的字段
func (this *MyRows) Int16(strFieldName string) int16 {
	var intValue int16
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取int32类型的字段
func (this *MyRows) Int32(strFieldName string) int32 {
	var intValue int32
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取int64类型的字段
func (this *MyRows) Int64(strFieldName string) int64 {
	var intValue int64
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint类型的字段
func (this *MyRows) Uint(strFieldName string) uint {
	var intValue uint
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint8类型的字段
func (this *MyRows) Uint8(strFieldName string) uint8 {
	var intValue uint8
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint16类型的字段
func (this *MyRows) Uint16(strFieldName string) uint16 {
	var intValue uint16
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint32类型的字段
func (this *MyRows) Uint32(strFieldName string) uint32 {
	var intValue uint32
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint64类型的字段
func (this *MyRows) Uint64(strFieldName string) uint64 {
	var intValue uint64
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取float32类型的字段
func (this *MyRows) Float32(strFieldName string) float32 {
	var floatValue float32
	this.GetValue(strFieldName, &floatValue)
	return floatValue
}

//取float32类型的字段
func (this *MyRows) Float(strFieldName string) float32 {
	return this.Float32(strFieldName)
}

//取float64类型的字段
func (this *MyRows) Float64(strFieldName string) float64 {
	var floatValue float64
	this.GetValue(strFieldName, &floatValue)
	return floatValue
}

//取string类型的字段
func (this *MyRows) String(strFieldName string) string {
	var strValue string
	this.GetValue(strFieldName, &strValue)
	return strValue
}

//取[]byte类型的字段
func (this *MyRows) Bytes(strFieldName string) []byte {
	var byteValue []byte
	this.GetValue(strFieldName, &byteValue)
	return byteValue
}

//取bool类型的字段
func (this *MyRows) Bool(strFieldName string) bool {
	var boolValue bool
	this.GetValue(strFieldName, &boolValue)
	return boolValue
}

//取Time类型的字段
func (this *MyRows) Time(strFieldName string) time.Time {
	var timeValue time.Time
	this.GetValue(strFieldName, &timeValue)
	return timeValue
}
