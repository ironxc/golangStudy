
//主入口,main
package main

//导入包,打包导入方式
import (
	"fmt"
	"math/rand"
	"math"
)
//一个个导入,包名与路径最后一个目录相同
// import "fmt"
// import "math/rand"
// import "math"
type Client struct{
	a string
}
type per struct{
	clients map[*Client]bool
}
func main() {
	fmt.Println("My favorite number is", rand.Intn(20))
	fmt.Println(math.Pi)
	aa := Client{a: "huhuhu"}

	var p = per{} 
	p.clients = map[*Client]bool{
		&aa: true,
	}

	fmt.Println(aa)
	fmt.Println(p)
	fmt.Println(p.clients[&aa])
}