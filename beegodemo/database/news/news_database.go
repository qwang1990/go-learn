package	main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"time"
)



type News struct {
	Id int
	Title string
	Source string
	Author string
	Publishtime time.Time
	Imageurl string
	Summary string
	Content string
	Count int
	Name int
	Deleted bool
}


func init() {

	// create table
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", "root@/chang?charset=utf8", 3)

	orm.RegisterModel(new(News))
	// create table
	orm.RunSyncdb("default", false, true)



}

func main() {
	orm.Debug = true
	o := orm.NewOrm()

	new  := News{Id:1}
	err := o.Read(&new)
	fmt.Printf("ERR: %v\n", err)
	fmt.Println(new)

	//
	var news []*News
	qs := o.QueryTable("news")
	num,err := qs.Filter("id__in", 1,2,3,4,5,6,7,8).All(&news)
	fmt.Printf("Returned Rows Num: %d, %s", num, err)

	//
	var maps []orm.Params
	num1, err1 := o.QueryTable("news").Values(&maps)
	if err1 == nil {
		fmt.Printf("Result Nums: %d\n", num1)
		for _, m := range maps {
			fmt.Println(m["Id"], m["Title"])
		}
	}

	//
	var lists []orm.ParamsList
	num2, err2 := o.QueryTable("news").ValuesList(&lists)
	if err2 == nil {
		fmt.Printf("Result Nums: %d\n", num2)
		for _, row := range lists {
			fmt.Println(row)
		}
	}

	//
	var list orm.ParamsList
	num3, err3 := o.QueryTable("news").ValuesFlat(&list, "Summary")
	if err3 == nil {
		fmt.Printf("Result Nums: %d\n", num3)
		//fmt.Printf("All User Names: %s", strings.Join(list, ", "))
		fmt.Println(list)
	}





}

