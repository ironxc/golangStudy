package main 

//switch和大多数语言一样
import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Print("Go runs on")
	// fmt.Println(runtime.GOOS)
	switch os := runtime.GOOS; os{
		case "darwin":
			fmt.Println("os x")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s.", os)		
	}

	os := runtime.GOOS
	fmt.Println(os)
	switch {
		case "darwin" == os:
			fmt.Println("os x")
		case "linux" == os:
			fmt.Println("Linux")
		default:
			fmt.Printf("%s.", os)
	}
}