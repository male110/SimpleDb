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
