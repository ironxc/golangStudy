package main

import "fmt"

/*
一个普遍存在的接口是 fmt 包中定义的 Stringer。

Stringer 是一个可以用字符串描述自己的类型。
`fmt`包 （还有许多其他包）使用这个来进行输出。

可以通过interface让函数接受任意类型的参数
其实是在调用自己的method


fmt源码中定义了

	type Stringer interface {
	     String() string
	}


*/
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	//这里调用了Person自己的String()方法
	fmt.Println(a, z)

	bird_a := bird{"鸟叫","叽叽喳喳~~~~~~~~~~~~"}
	duck_a := duck{"鸭叫","嘎嘎嘎嘎~~~~~~~~~~~~"}
	Say(bird_a)
	Say(duck_a)
}

type animal interface{
	Say() 
}

type bird struct {
	Name string
	Str string
}

type duck struct{
	Name string
	Str string
}
func (v bird)Say() {
	fmt.Println(v.Name,":",v.Str)	
}
func (v duck)Say() {
	fmt.Println(v.Name,":",v.Str)	
}

func Say(v animal) {
	v.Say()
}