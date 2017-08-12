package main

import (
	"fmt"
	"math"
)

const Pi = 3.14 //定义常量

//类型转换
//var u uint = uint(f)
// u := unit(f)
func main() {
	var x, y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x,y,z)
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	t := 42.1 // 根据变量值的精度确定类型
	fmt.Printf("v is of type %T\n", t)

	// Pi = 100;//常量不能变化
	//常量不能使用:=声明符号
	fmt.Println("常量:", Pi)
}