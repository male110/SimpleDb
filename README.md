<p>
&nbsp;&nbsp;&nbsp;&nbsp; 这个是我用Go写的第一个东东，可能还存在些BUG没有测试到。这里主要是提供一个参考。各位可以改写成自己的风格。<br />
</p>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 在命令行下输入如下两条命令，进行安装</p>
<div style="background-color: F8F8F8">
<pre>
		go get github.com/male110/SimpleDb
		go install github.com/male110/SimpleDb

</pre>
    </div>
    
<p>
    Go语言的数据库操作，只能用Rows.Scan来一次性读取所有列。感觉很不习惯，我还是习惯按照列名来一列列的取数据。所以我自己封装了一个数据结构MyRows,MyRows实现了一个函数，<span
        style="color: #000000;">GetValue(name<span style="color: #c0c0c0;"> </span>
    <span style="color: #000080;">string</span><span style="color: #000000;">,</span><span
        style="color: #c0c0c0;"> </span>value<span style="color: #c0c0c0;"> </span><span
            style="font-weight: 600; color: #000080;">interface</span><span style="color: #000000;">{})可以按列名来取数据。如下所示：</span></p>
<div style="background-color: F8F8F8">
<pre>
		err = rows.GetValue("IsBoy", &isBoy)
		if err != nil {
			fmt.Println(err)
			return
		}
</pre>
    </div>
<p>
    &nbsp;&nbsp;&nbsp; 为了操作方便，还定义了其它的结构体，如MyDb，其Query函数可以直接返回&nbsp;MyRows。NewDb用来创建MyDb结构，其参数与sql.Open一至，怎么传取决于你所使用的驱动程序。</p>
<div style="background-color: F8F8F8">
<pre>
    db, err := SimpleDb.NewDb("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return
	}
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
		//按字段名取数据,也可以用rows.Scan(&id,&name,&age),来取
		rows.GetValue("id", &id)
		rows.GetValue("name", &name)

		rows.GetValue("age", &age)
		//可以根据返回值，判断是否成功
		err = rows.GetValue("IsBoy", &isBoy)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, "\t", name, "\t", age, "\t", isBoy)
	}
</pre>
    </div>
<p>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    同时还时现了一个简单的ORM，实现了最基本的插入数据，修改数据，删除数据。 
    我一般使用ORM只用这么几个方法，其它的都是写SQL语句。这里只是一个参考，大家可以根据自己的需要，自己习惯，进行修改。改成自己需要的格式。数据结构的定义格式如下：</p>
<div style="background-color: F8F8F8">
<pre>
type Person struct {
	/*TableName类型只是用来设置表名。如果结构体名跟表名相同可以省略*/
	TableName SimpleDb.TableName "person"

	/*name是表名,PK用来设置是否主键，true主键，false非主键*/
	Id int `name:"id"PK:"true"Auto:"true"`

	Name   string "name" //tag里的name表是对应的字段名
	Age    int    "age"  //tag里的age表是对应的字段名
	IsBoy  bool
	NotUse string "-" //-不会保存到数据库中
}
</pre>
</div>
<p>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
上面的说明已经很详细了，
SimpleDb.TableName类型的字段，只用来在tag中定义结构体对应的表名，如果没有该字段，认为表名就是结构体名相同。PK:&quot;true&quot;表示是主键，Auto:&quot;true&quot;表示该字段是自动增长的列，name:&quot;id&quot;,来指定该字段对应的数据表中的列名，如不指定认为跟字段名相同。当只需要指定列名时，可以直接写在tag中，如：<span
        style="color: #000000;">&quot;name&quot;、&quot;age&quot;</span>。tag为&quot;-&quot;表示不对应数据表中的任何列。</p>

<div style="background-color: F8F8F8">
<pre>
// 插入数据
p := &Person{Name: "张三丰", Age: 500, IsBoy: true}
db.Insert(p)
//修改数据
db.Update(p)
//删除数据
db.Delete(p)
</pre>
</div>
<p>
   下面来看一个完整的例子，首先他创表:</p>
<div style="background-color: F8F8F8">
<pre>
CREATE TABLE `person` (
	`id` INT(50) NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(50) NULL DEFAULT NULL,
	`age` INT(50) NULL DEFAULT NULL,
	`IsBoy` SMALLINT(10) NULL,
	PRIMARY KEY (`id`)
)
COLLATE='utf8_general_ci';

insert into `person` (name,age,IsBoy) values('张三',20,0);
insert into `person` (name,age,IsBoy) values('王五',19,1);
</pre>
</div>
</span>
<p>
    下面是完整的代码</p>
<div style="background-color: F8F8F8">
<pre>
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/male110/SimpleDb"
)

type Person struct {
	/*TableName类型只是用来设置表名。如果结构体名跟表名相同可以省略*/
	TableName SimpleDb.TableName "person"

	/*name是表名,PK用来设置是否主键，true主键，false非主键*/
	Id int `name:"id"PK:"true"Auto:"true"`

	Name   string "name" //tag里的name表是对应的字段名
	Age    int    "age"  //tag里的age表是对应的字段名
	IsBoy  bool
	NotUse string "-" //-不会保存到数据库中
}

func main() {
	db, err := SimpleDb.NewDb("mysql", "root:123@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("打开SQL时出错:", err.Error())
		return
	}
	defer db.Close()
	p := &Person{Name: "张三丰", Age: 500, IsBoy: true}
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
		//按字段名取数据,也可以用rows.Scan(&id,&name,&age),来取
		rows.GetValue("id", &id)
		rows.GetValue("name", &name)

		rows.GetValue("age", &age)
		//可以根据返回值，判断是否成功
		err = rows.GetValue("IsBoy", &isBoy)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, "\t", name, "\t", age, "\t", isBoy)
	}
	//输出分割线
	fmt.Println("==========割割割割割割割割============")
	p.Name = "彭祖"
	p.Age = 800
	//修改数据
	_, err = db.Update(p)
	if err != nil {
		fmt.Println(err, "xxxx")
		return
	}
	//QueryDataRows返回一个DataRow数组，DataRow中有一map来存放行中的数据
	var arrRow []SimpleDb.DataRow
	arrRow, err = db.QueryDataRows("select * from person")
	if err != nil {
		fmt.Println(err, "zzzzz")
		return
	}
	for _, row := range arrRow {
		var id, age int
		var name string
		var isBoy bool
		//只能按字段名取数据
		row.GetValue("id", &id)
		row.GetValue("name", &name)
		row.GetValue("age", &age)
		//可以根据返回值，判断是否成功
		err = rows.GetValue("IsBoy", &isBoy)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id, "\t", name, "\t", age, isBoy)
	}
	var p2 Person
	p2.Id = p.Id
	//根据主键从数据库中取单条数据
	err = db.Load(&p2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)
	//根据主键删除一条数据
	db.Delete(p2)
}
</pre>
</div>

