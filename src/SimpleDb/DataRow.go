package SimpleDb

import (
	"errors"
	"time"
)

type DataRow struct {
	column map[string]interface{}
}

func (this *DataRow) copyMap(dest, src map[string]interface{}) {
	for k, v := range src {
		dest[k] = v
	}
}

/*根据字段名来取字段的值
name:字段名，区分大小写
value:用来接收字段值的变量，需传变量的地址,如&a
*/
func (this *DataRow) GetValue(name string, value interface{}) error {
	i, ok := this.column[name]
	if ok {
		err := ConvertAssign(value, i)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("字段不存在，请注意大小写")
}

//取int类型的字段
func (this *DataRow) Int(strFieldName string) int {
	var intValue int
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取int8类型的字段
func (this *DataRow) Int8(strFieldName string) int8 {
	var intValue int8
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取int16类型的字段
func (this *DataRow) Int16(strFieldName string) int16 {
	var intValue int16
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取int32类型的字段
func (this *DataRow) Int32(strFieldName string) int32 {
	var intValue int32
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取int64类型的字段
func (this *DataRow) Int64(strFieldName string) int64 {
	var intValue int64
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint类型的字段
func (this *DataRow) Uint(strFieldName string) uint {
	var intValue uint
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint8类型的字段
func (this *DataRow) Uint8(strFieldName string) uint8 {
	var intValue uint8
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint16类型的字段
func (this *DataRow) Uint16(strFieldName string) uint16 {
	var intValue uint16
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint32类型的字段
func (this *DataRow) Uint32(strFieldName string) uint32 {
	var intValue uint32
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取uint64类型的字段
func (this *DataRow) Uint64(strFieldName string) uint64 {
	var intValue uint64
	this.GetValue(strFieldName, &intValue)
	return intValue
}

//取float32类型的字段
func (this *DataRow) Float32(strFieldName string) float32 {
	var floatValue float32
	this.GetValue(strFieldName, &floatValue)
	return floatValue
}

//取float32类型的字段
func (this *DataRow) Float(strFieldName string) float32 {
	return this.Float32(strFieldName)
}

//取float64类型的字段
func (this *DataRow) Float64(strFieldName string) float64 {
	var floatValue float64
	this.GetValue(strFieldName, &floatValue)
	return floatValue
}

//取string类型的字段
func (this *DataRow) String(strFieldName string) string {
	var strValue string
	this.GetValue(strFieldName, &strValue)
	return strValue
}

//取[]byte类型的字段
func (this *DataRow) Bytes(strFieldName string) []byte {
	var byteValue []byte
	this.GetValue(strFieldName, &byteValue)
	return byteValue
}

//取bool类型的字段
func (this *DataRow) Bool(strFieldName string) bool {
	var boolValue bool
	this.GetValue(strFieldName, &boolValue)
	return boolValue
}

//取Time类型的字段
func (this *DataRow) Time(strFieldName string) time.Time {
	var timeValue time.Time
	this.GetValue(strFieldName, &timeValue)
	return timeValue
}
