
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

func main() {
	fmt.Println("My favorite number is", rand.Intn(20))
	fmt.Println(math.Pi)
}