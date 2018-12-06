package main
import (
    "fmt"
    "time"
)
//chan会阻塞来保证携程的同步
func fixed_shooting(msg_chan chan string) {
    for {
        msg_chan <- "fixed shooting"//输入数据必须输出，不然会阻塞
        fmt.Println("continue fixed shooting...")
    }
}

func count(msg_chan chan string) {
    for {
        msg := <-msg_chan
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}
func main() {
    var c chan string
    c = make(chan string)
    go fixed_shooting(c)
    go count(c)
    var input string
    fmt.Scanln(&input)
}