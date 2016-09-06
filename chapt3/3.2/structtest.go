package main

import "fmt"

type Rect struct {
	W,Y float64
	width,height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}


func main() {
	//几种创建对象的方法
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{1,2,100,200}
	rect4 := &Rect{width:100,height:200}

	fmt.Println(rect1)
	fmt.Println(rect2)
	fmt.Println(rect3)
	fmt.Println(rect4)
}