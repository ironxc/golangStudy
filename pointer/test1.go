package main 

import "fmt"

// func aa(a int) {
// 	fmt.Println(a)
// }
func main() {
	// var a = aa(10) //函数未返回值，就不能赋值给变量
	// fmt.Println(a)
	 a := []int{1 ,3}
	var b [2]int

	fmt.Println(a)
	fmt.Println(a , len(a), cap(a))
	fmt.Println(b)

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}