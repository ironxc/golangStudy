package main 

import ( 

	"fmt"
 )
 

type Person struct {
	x, y string
}

type MyInt int

//可以在任意类型上定义方法
func (v *Person) Say() {
	fmt.Println(v.x)
	fmt.Println(v.y)
}
func (v MyInt) Say() {
	fmt.Println(v)
}

type Location struct {
	x, y int
}
// func (l *Location) Transform() (int, int) {
// 	temp := l.x
// 	l.x = l.y
// 	l.y = temp
// 	return  l.x, l.y
// }

func (l Location) Transform() (int, int) {
	temp := l.x
	l.x = l.y
	l.y = temp
	return  l.x, l.y
}
func main() {
	
	p1 := &Person{x:"hello", y:"Hi"}

	p1.Say()
	p2 := Person{x:"ll", y:"mm"}
	p2.Say()
	int1 := MyInt(10)
	fmt.Println(int1)
	// int1.Say()
	l1 := &Location{12,13}
	x,y := l1.Transform()
	fmt.Println("x", x, "y", y, l1)
}