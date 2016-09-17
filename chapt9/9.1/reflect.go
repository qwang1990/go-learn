package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	p := reflect.ValueOf(&x)
	fmt.Println("type: ",p.Type())
	fmt.Println(p.CanSet())

	v := p.Elem()
	fmt.Println(v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
