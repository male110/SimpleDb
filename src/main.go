package main

import (
	"SimpleDb"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Person struct {
	/*TableName类型只是用来设置表名。如果结构体名跟表名相同可以省略*/
	TableName SimpleDb.TableName "person"

	/*name是表名,PK用来设置是否主键，true主键，false非主键*/
	Id int `name:"id"PK:"true"Auto:"true"`

	Name    string "name" //tag里的name表是对应的字段名
	Age     int    "age"  //tag里的age表是对应的字段名
	IsBoy   bool
	NotUse  string "-" //-不会保存到数据库中
	AddDate time.Time
}

func main() {
	db, err := SimpleDb.NewDb("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return
	}
	defer db.Close()
	p := &Person{Name: "张三丰", Age: 500, IsBoy: true}
	p.AddDate = time.Now()
	//fmt.Println(p.AddDate)
	//fmt.Println(p.AddDate.Unix())
	//插入一条数据
	err = db.Insert(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("新插入数据的ID:", p.Id)
	var rows *SimpleDb.MyRows
	//从数据库中取数据
	rows, err = db.Query("select * from person")
	if err != nil {
		fmt.Println(err)
		return
	}
	//显示数据
	for rows.Next() {
		var id, age int
		var name string
		var isBoy bool
		var addDate time.Time
		//var a time.Time
		//按字段名取数据,也可以用rows.Scan(&id,&name,&age),来取
		rows.GetValue("id", &id)
		rows.GetValue("name", &name)

		rows.GetValue("age", &age)
		rows.GetValue("IsBoy", &isBoy)
		//可以根据返回值，判断是否成功
		err = rows.GetValue("AddDate", &addDate)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, "\t", name, "\t", age, "\t", isBoy, "\t", addDate)
		fmt.Println(rows.Int("id"), "\t", rows.String("name"), "\t", rows.Int("age"), "\t", rows.Bool("IsBoy"), "\t", rows.Time("AddDate"))
	}
	rows.Close()
	//输出分割线
	fmt.Println("==========割割割割割割割割============")
	p.Name = "彭祖"
	p.Age = 800
	//修改数据
	_, err = db.Update(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	//QueryDataRows返回一个DataRow数组，DataRow中有一map来存放行中的数据
	var arrRow []SimpleDb.DataRow
	arrRow, err = db.QueryDataRows("select * from person")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range arrRow {
		//可以用String取所有字段的值用来显示
		fmt.Println(row.String("id"), "\t", row.String("name"), "\t", row.String("age"), "\t", row.String("IsBoy"), "\t", row.String("AddDate"))
	}
	var p2 Person
	p2.Id = p.Id
	//根据主键从数据库中取单条数据
	err = db.Load(&p2)
	if err != nil {
		fmt.Println(p2)
		fmt.Println("xxxxxx:", err)
		return
	}
	fmt.Println(p2)
	//根据主键删除一条数据
	db.Delete(p2)
}
