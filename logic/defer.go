
package main 

import "fmt"


func fun1() {
	defer fmt.Println("等上层函数结束我在跑")

	fmt.Println("先跑一个")

	fmt.Println("再跑一个")

	for v := 1 ; v <10 ; v++{
		fmt.Println(v,"run")
	}
}
func main() {

	//defer的延迟函数调用会被压入一个栈，先进后出
	defer fmt.Println("第一个进去的","world")

	fmt.Println("hello")
	fun1()

	for i := 2; i<10 ; i++ {
		defer fmt.Println("第",i,"个进去的")
	}

}