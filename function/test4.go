package main

import "fmt"

var i, j int = 1, 2
// k := 3//不能在函数外使用 :=

func main() {
	var c , python ,java = true, false, "no!"
	// var i , j int = 100, 1000
	// k := 3

	//如果初始化直接给值，则不用给类型，会自动在值中获得类型
	fmt.Println(c, python ,java ,i , j)
	fmt.Println(":=可以替代var声明变量,--但是在函数外不能用:",k)

}