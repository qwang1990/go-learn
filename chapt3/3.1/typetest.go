package main

import "fmt"

//为类型添加方法
type Integer int
func (a Integer) Less(b Integer) bool {
	return a > b
}
//因为需要改变a的值了,所以用指针当参数
func (a *Integer) Add(b Integer) {
	*a += b
}

func main()  {
	var a Integer  = 1
	if a.Less(2) {
		fmt.Println(a,"Less 2")
	}

	a.Add(2)
	fmt.Println("a = ",a)
}
