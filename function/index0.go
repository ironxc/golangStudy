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
func main() {
	
	p1 := &Person{x:"hello", y:"Hi"}

	p1.Say()

	int1 := MyInt(10)
	fmt.Println(int1)
	// int1.Say()

}