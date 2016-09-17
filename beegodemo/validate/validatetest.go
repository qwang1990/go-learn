package main
import (
	"log"
	"strings"

	"github.com/astaxie/beego/validation"
)

// 验证函数写在 "valid" tag 的标签里
// 各个函数之间用分号 ";" 分隔，分号后面可以有空格
// 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
// 正则函数(Match)的匹配模式用两斜杠 "/" 括起来
// 各个函数的结果的key值为字段名.验证函数名
type user struct {
	Id     int
	Name   string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以Bee开头
	Age    int    `valid:"Range(1, 140)"` // 1 <= Age <= 140，超出此范围即为不合法
	Email  string `valid:"Email; MaxSize(100)"` // Email字段需要符合邮箱格式，并且最大长度不能大于100个字符
	Mobile string `valid:"Mobile"` // Mobile必须为正确的手机号
	IP     string `valid:"IP"` // IP必须为一个正确的IPv4地址
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}

func main() {
	valid := validation.Validation{}
	u := user{Name: "Beego", Age: 2, Email: "dev@beego.me",Mobile:"13011112222",IP:"127.256.0.1"}
	b, err := valid.Valid(&u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
}
