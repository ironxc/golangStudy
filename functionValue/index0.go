package main 

import "fmt"
//函数可以作为值
//函数也可以是闭包，返回一个或多个绑定了内部值的函数
 func adder() (func(int) int ,func(int) int ) {
 	suma, sumb := 0,0
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