package main

import "fmt"

func main()  {
	fmt.Printf("hello world\n")
	_,_,name:=GetName();
	fmt.Printf("name=%s\n",name);
	//常量
	const (
		a = iota
		b = iota
	)
	fmt.Printf("%d  %d\n",a,b)

	//数组&切片
	var myArray [10] int = [10] int {1,2,3,4,5,6,7,8,9,10}

	var mySlice [] int = myArray[:5]

	fmt.Println("Elements of myArray")
	for _,v:= range myArray {
		fmt.Print(v," ")
	}

	fmt.Println("\nElements of mySlice: ")

	for _,v := range mySlice {
		fmt.Print(v," ")
	}

	fmt.Println()

	//第二种创建slice的方法 并且长度和容量不一定一样
	mySlice2 := make([]int,5,10)
	fmt.Println("len(mySlice2)",len(mySlice2))
	fmt.Println("cap(mySlice2)",cap(mySlice2))

	//map
	var personDB map[string] PersonInfo
	    personDB = make(map[string] PersonInfo)

	personDB["12345"] = PersonInfo{"12345","Tom","Room 203,..."}
	personDB["1"] = PersonInfo{"1","Jack","Room 101,..."}

	person, ok := personDB["12345"]
	if ok {
		fmt.Println("found person",person.Name)
	} else {
		fmt.Println("do not find")
	}

	i := example(10)
	fmt.Println(i)

	//不定参数
	myfunc(1,2,3,4,5,6)

	//任意类型的不定参数
	var v1 int  = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234

	MyPrintf(v1,v2,v3,v4)

	//闭包(保证了i的安全性)
	var j int = 5

	closure := func() (func()) {
		var i int = 10
		return func() {
			fmt.Printf("i,j: %d, %d\n",i,j)
		}
	}()

	closure()

	j *= 2

	closure()


}

type PersonInfo struct {
	ID string
	Name string
	Address string
}

func GetName() (firstname,lastname,nickname string) {
	return "may","chan","wang"
}

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func myfunc(args ...int)  {
	for _,arg := range args {
		fmt.Println(arg)
	}
}

func MyPrintf(args ...interface{})  {
	for _,arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int value")
		case string:
			fmt.Println(arg, "is a string value")
		case int64:
			fmt.Println(arg, "is a int64 value")
		default:
			fmt.Println(arg, "is a defult value")
		}
	}
}









