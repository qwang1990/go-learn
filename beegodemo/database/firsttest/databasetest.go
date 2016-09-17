package	main

import (
"fmt"
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {

	// create table
	orm.RunSyncdb("default", false, true)

	// set default database
	orm.RegisterDataBase("default", "mysql", "root@/my_db?charset=utf8", 30)

}

func main() {
	orm.Debug = true
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// insert
	user2 := User{Name: "slene2"}
	 o.Insert(&user2)

	// insert
	user3 := User{Name: "slene3"}
	 o.Insert(&user3)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)

	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM user").Values(&maps)
	fmt.Println("一共",num,"数据")
	for _,term := range maps{
		fmt.Println(term["id"],":",term["name"])
	}



	// update
	//user.Name = "astaxie"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//// read one
	//u := User{Id: user.Id}
	//err = o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)
	//
	//// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

