package main 

import "fmt"
//类型 [n]T 是一个有 n 个类型为 T 的值的数组
//数组必须长度，长度是它的类型的一部分，因此数组不能改变长度
// slice的长度是可以改变的，可以认为是动态数组
func main() {
	var a [2]string

	a[0] = "hello"
	a[1] = "world"
	
	fmt.Println(a[0],a[1])
	fmt.Println(a)
	fmt.Println("&a",&a)
	fmt.Println(len(a))
}