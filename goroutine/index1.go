package main

/*
	channel用于goroutine间的通信
	必须使用make 创建channel
	ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})

	ch <- v //发送到channel ch
	v := <-ch //从ch中接收数据并赋值给v
*/
/*
默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，
这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。
所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。
其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。
无缓冲channel是在多个goroutine之间同步很棒的工具。
*/
import "fmt"

func sum(a []int, c chan int) {
    total := 0
    for _, v := range a {
        total += v
    }
    c <- total  // send total to c
}

func main() {
    a := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int) //创建一个值类型为Int的channel
    go sum(a[:len(a)/2], c) //开启一个协程
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c  // receive from c

    z := <-c
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
    fmt.Println(x, y, x + y)
}
