package SimpleDb

type MyRow struct {
	rows      *MyRows
	isHasRows bool
}

func (this *MyRow) readData() {
	defer this.rows.Close()
	this.isHasRows = this.rows.Next() //在Next中会把数据读到rows.values里
}

/*根据字段名来取字段的值
name:字段名，区分大小写
value:用来接收字段值的变量，需传变量的地址,如&a
*/
func (this *MyRow) GetValue(name string, value interface{}) error {
	if this.isHasRows == false {
		return ErrNoRows
	}
	return this.rows.GetValue(name, value)
}
func (this *MyRow) Scan(dest ...interface{}) error {
	if this.isHasRows == false {
		return ErrNoRows
	}
	return this.rows.Scan(dest...)
}
