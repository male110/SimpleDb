// MyTx
package SimpleDb

import (
	"database/sql"
)

type MyTx struct {
	*sql.Tx
}

func (this *MyTx) Prepare(query string) (*MyStmt, error) {
	t, err := this.Tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	mysmt := &MyStmt{Stmt: t}
	return mysmt, nil
}

func (this *MyTx) Query(query string, args ...interface{}) (*MyRows, error) {
	r, err := this.Tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	rows := &MyRows{Rows: r, values: make(map[string]interface{})}
	return rows, nil
}

func (this *MyTx) QueryRow(query string, args ...interface{}) (*MyRow, error) {
	r, err := this.Query(query, args...)
	if err != nil {
		return nil, err
	}
	row := &MyRow{rows: r, IsHasRows: false}
	row.readData()
	return row, nil
}
