// MyStmt
package SimpleDb

import (
	"database/sql"
)

type MyStmt struct {
	*sql.Stmt
}

func (this *MyStmt) Query(args ...interface{}) (*MyRows, error) {
	r, err := this.Stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	myrows := &MyRows{Rows: r, values: make(map[string]interface{})}
	return myrows, nil
}
func (this *MyStmt) QueryRow(args ...interface{}) (*MyRow, error) {
	//这里调用了query来实现
	r, err := this.Query(args...)
	if err != nil {
		return nil, err
	}
	myrow := &MyRow{rows: r}
	myrow.readData()
	return myrow, nil
}
