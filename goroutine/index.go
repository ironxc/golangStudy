package main

import (
	"fmt"
	"runtime"
)

//runtime.Gosched(),把CPU时间片让给别人下次继续恢复执行
// runtime.Goexit() 停止
func say(s string) {

	
	fmt.Println(s)
    // runtime.Gosched()
    // runtime.Goexit()
    fmt.Println(runtime.NumCPU())
    fmt.Println("换goroutine")
    
}
func main() {
	say("hello")
	// go say("world")
}