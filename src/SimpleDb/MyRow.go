package SimpleDb

import (
	"time"
)

type MyRow struct {
	rows      *MyRows
	IsHasRows bool //true表示有数据，false没数据
}

func (this *MyRow) readData() {
	defer this.rows.Close()
	this.IsHasRows = this.rows.Next() //在Next中会把数据读到rows.values里
}

/*根据字段名来取字段的值
name:字段名，区分大小写
value:用来接收字段值的变量，需传变量的地址,如&a
*/
func (this *MyRow) GetValue(name string, value interface{}) error {
	if this.IsHasRows == false {
		return ErrNoRows
	}
	return this.rows.GetValue(name, value)
}

//取int类型的字段
func (this *MyRow) Int(strFieldName string) int {
	return this.rows.Int(strFieldName)
}

//取int8类型的字段
func (this *MyRow) Int8(strFieldName string) int8 {
	return this.rows.Int8(strFieldName)
}

//取int16类型的字段
func (this *MyRow) Int16(strFieldName string) int16 {
	return this.rows.Int16(strFieldName)
}

//取int32类型的字段
func (this *MyRow) Int32(strFieldName string) int32 {
	return this.rows.Int32(strFieldName)
}

//取int64类型的字段
func (this *MyRow) Int64(strFieldName string) int64 {
	return this.rows.Int64(strFieldName)
}

//取uint类型的字段
func (this *MyRow) Uint(strFieldName string) uint {
	return this.rows.Uint(strFieldName)
}

//取uint8类型的字段
func (this *MyRow) Uint8(strFieldName string) uint8 {
	return this.rows.Uint8(strFieldName)
}

//取uint16类型的字段
func (this *MyRow) Uint16(strFieldName string) uint16 {
	return this.rows.Uint16(strFieldName)
}

//取uint32类型的字段
func (this *MyRow) Uint32(strFieldName string) uint32 {
	return this.rows.Uint32(strFieldName)
}

//取uint64类型的字段
func (this *MyRow) Uint64(strFieldName string) uint64 {
	return this.rows.Uint64(strFieldName)
}

//取float32类型的字段
func (this *MyRow) Float32(strFieldName string) float32 {
	return this.rows.Float32(strFieldName)
}

//取float32类型的字段
func (this *MyRow) Float(strFieldName string) float32 {
	return this.rows.Float(strFieldName)
}

//取float64类型的字段
func (this *MyRow) Float64(strFieldName string) float64 {
	return this.rows.Float64(strFieldName)
}

//取string类型的字段
func (this *MyRow) String(strFieldName string) string {
	return this.rows.String(strFieldName)
}

//取[]byte类型的字段
func (this *MyRow) Bytes(strFieldName string) []byte {
	return this.rows.Bytes(strFieldName)
}

//取bool类型的字段
func (this *MyRow) Bool(strFieldName string) bool {
	return this.rows.Bool(strFieldName)
}

//取Time类型的字段
func (this *MyRow) Time(strFieldName string) time.Time {
	return this.rows.Time(strFieldName)
}
func (this *MyRow) Scan(dest ...interface{}) error {
	if this.IsHasRows == false {
		return ErrNoRows
	}
	return this.rows.Scan(dest...)
}
