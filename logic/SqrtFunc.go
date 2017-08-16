package main 

import "fmt"

func Sqrt(x float64) float64 {
	z := float64(1)
	fmt.Println("z的值",z - z)
}

func main() {
	fmt.Println(Sqrt(2))
}