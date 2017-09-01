package main

import (
	"fmt"
	"runtime"
)

//runtime.Gosched(),把CPU时间片让给别人下次继续恢复执行
func say(s string) {

	fmt.Println("aa")
	fmt.Println(s)
    runtime.Gosched()
    // runtime.Goexit()
    // fmt.Println(runtime.NumCPU())
    fmt.Println("换goroutine")
    
}
func main() {
	go say("world")
	say("hello")
}