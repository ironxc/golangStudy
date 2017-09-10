package main 

import "fmt"
//函数可以作为值
//函数也可以是闭包，返回一个或多个绑定了内部值的函数
/*
Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。 
函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。
*/
 func adder() (func(int) int ,func(int) int ) {
 	suma, sumb := 10,10
 	return func (x int) int {
 			suma += x
 			return suma },
 		   func (x int) int {
 			sumb -=x
 			return sumb }
 }
func main() {
	sum := func (a int, b int) int {
		c := a + b
		return c
	}
	fmt.Println(sum(1,2))

	
	a , b := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i),b(i)) 
	}
}