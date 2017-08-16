package main 

import (
	"fmt"
	"math"
)
//go没有类，但是可以在struct上定义方法，就像是某个类型的实例附带的函数

type Cat struct {
	name string
	weight int
}

type Vertex struct {
	X, Y float64
}



func (v *Vertex) Abs() float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}



func (v Cat) say() int{
	return v.weight
}

type MyFloat float64

//任意类型都可以定义方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {

	c := Cat{"kit" ,2}
	fmt.Println(c.say())

	//传指针更快，因为不用copy值
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}