package SimpleDb

import (
	"errors"
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
