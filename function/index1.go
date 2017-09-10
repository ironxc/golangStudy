package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Ver struct {
	X, Y float64
}

//传值是copy了一份并不会改变原数据
func (v Ver) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Ver) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// v := &Vertex{3, 4}
	v := Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())

	ve := Ver{3, 4}
	ve.Scale(5)
	fmt.Println(ve, ve.Abs())

}