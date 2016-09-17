package main
import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"io"
)


func init() {
	fmt.Println("main init")
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root@/orm_test?charset=utf8")

	orm.RegisterModel(new(User), new(Profile),new(Post),new(Tag))
	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	var w io.Writer
	orm.DebugLog = orm.NewLog(w)
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
	orm.ResetModelCache()

	//orm.RegisterDataBase("db1", "mysql", "root@/orm_test?charset=utf8")
	//
	//
	//o1 := orm.NewOrm()
	//o1.Using("db1")
	//dr := o1.Driver()
	//fmt.Println(dr.Name() == "db1") // true
	//fmt.Println(dr.Type()) // true



}

