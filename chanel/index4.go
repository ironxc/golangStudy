package main
import (
    "fmt"
    "time"
)
func fixed_shooting(msg_chan chan string) {
    for {
				msg_chan <- "A"
				fmt.Println("A输入")
    }
}
func three_point_shooting(msg_chan chan string) {
    for {
				msg_chan <- "B"
				fmt.Println("B输入")
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
    go three_point_shooting(c)
    go count(c)
    var input string
    fmt.Scanln(&input)
}