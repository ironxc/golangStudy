package main 
import (
	"fmt"
)
func main() {
	//这样会阻塞导致死锁，因为channel没缓存
    // c := make(chan int)
    // c <- 1
    // x := <-c
    // fmt.Println(x)
	
	c := make(chan int, 1)//添加缓存
	c <- 1
	x := <-c
	fmt.Println(x)

	
	
	ch := make(chan int)
	go func(){
		ch <- 1
	}()

	fmt.Println( <- ch)
}