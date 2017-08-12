package main 

import "fmt"
//类型 [n]T 是一个有 n 个类型为 T 的值的数组
func main() {
	var a [2]string

	a[0] = "hello"
	a[1] = "world"
	
	fmt.Println(a[0],a[1])
	fmt.Println(a)
	fmt.Println("&a",&a)
	fmt.Println(len(a))
}