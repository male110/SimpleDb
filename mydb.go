package SimpleDb

import (
	"database/sql"
)

type MyDb struct {
	*sql.DB
}

func NewDb(driverName, dataSourceName string) (*MyDb, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	mydb := &MyDb{DB: db}
	return mydb, nil
}

func (this *MyDb) Query(query string, args ...interface{}) (*MyRows, error) {
	r, err := this.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}
	myrows := &MyRows{Rows: r, values: make(map[string]interface{})}
	return myrows, nil
}

func (this *MyDb) QueryRow(query string, args ...interface{}) (*MyRow, error) {
	//这里调用了query来实现
	r, err := this.Query(query, args...)
	if err != nil {
		return nil, err
	}
	myrow := &MyRow{rows: r}
	myrow.readData()
	return myrow, nil
}

func (this *MyDb) Prepare(query string) (*MyStmt, error) {
	s, err := this.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	mystmt := &MyStmt{Stmt: s}
	return mystmt, nil
}

func (this *MyDb) Begin() (*MyTx, error) {
	t, err := this.DB.Begin()
	if err != nil {
		return nil, err
	}
	mytx := &MyTx{Tx: t}
	return mytx, nil
}

func (this *MyDb) QueryDataRows(query string, args ...interface{}) ([]DataRow, error) {
	rows, err := this.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []DataRow
	for rows.Next() {
		datarow := DataRow{column: make(map[string]interface{})}
		datarow.copyMap(datarow.column, rows.values)
		result = append(result, datarow)
	}
	return result, nil
}

/*将一个Struct结构体插入到数据库,如有自增自段，
会把自增值赋给model中的对应字段，model必须是可修改的，即传地址如，&m*/
func (this *MyDb) Insert(model interface{}) error {
	strSql, param, tabinfo, err := generateInsertSql(model)
	if err != nil {
		return err
	}
	var result sql.Result
	result, err = this.Exec(strSql, param...)
	if err != nil {
		return err
	}
	setAuto(result, tabinfo)
	return nil
}

//根据主键更新数据记录,返回所影响的行数
func (this *MyDb) Update(model interface{}) (int64, error) {
	strSql, param, _, err := generateUpdateSql(model)
	if err != nil {
		return 0, err
	}
	var result sql.Result

	result, err = this.Exec(strSql, param...)
	if err != nil {
		return 0, err
	}
	var afect int64
	afect, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return afect, nil
}

//根据主键返回一条数据，并赋值到model中，model必须是可修改的，即传地址如，&m
func (this *MyDb) Load(model interface{}) error {
	strSql, param, tabinfo, err := generateLoadSql(model)
	if err != nil {
		return err
	}
	var row *MyRow
	row, err = this.QueryRow(strSql, param...)
	if err != nil {
		return err
	}
	err = SetFieldValue(row, tabinfo)
	if err != nil {
		return err
	}
	return nil
}

//根据主键删除一条数据,返回所影响的行
func (this *MyDb) Delete(model interface{}) (int64, error) {
	strSql, param, _, err := generateDeleteSql(model)
	if err != nil {
		return 0, err
	}
	var result sql.Result
	result, err = this.Exec(strSql, param...)
	if err != nil {
		return 0, err
	}
	var effect int64
	effect, err = result.LastInsertId()
	return effect, err
}
