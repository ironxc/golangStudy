package main 

import "fmt"


//struct ,结构体，一个字段的集合
type Vertex struct{
	X int
	Y int
}

type student struct{
	age  int8
	className , name string
}
func main() {

	//像js中的对象，点取值，同时也会改变原struct变量
	i := 100
	fmt.Println(Vertex{1, 2})
	v := Vertex{
		1, 2,
	}
	v.X = 4
	fmt.Println(v.X)

	fmt.Println("改变后的v",v)

	// ..结构体指针
	p := &v
	m := &i
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(m)
	p.X = 1e9
	fmt.Println(v)

	fmt.Println("------------------")

	var(

		v1 = student{18, "三年级", "小王"}
		v2 = student{}
		v3 = student{name:"小张", age:19, className:"四年级"}
		p1 = &student{name:"小张", age:19 }
	)

	fmt.Println(v1, v2, v3, p1 )

	var xx = 100;

	fmt.Println(&xx)
	fmt.Println(bool(*&xx==xx))
}
