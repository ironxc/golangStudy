/*
channel 可以是 _带缓冲的_。
为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：

ch := make(chan int, 100)
向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。
当缓冲区清空的时候接受的时候会阻塞。

修改例子使得缓冲区被填满，然后看看会发生什么。
*/

package main

import "fmt"

func prt(c chan int, num int) {
	c <- num
}
func main() {
	c := make(chan int, 100)
	var a , b int
	a = 1001
	b = 12
	go prt(c, a)
	go prt(c, b)
	v := <-c
	fmt.Println(<-c)
	fmt.Println(v)
}