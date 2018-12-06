package main

import "fmt"

const (
	Big = 1 << 100 //1乘以2的10次方
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x*0.1	
}

func main() {
	fmt.Println("---------------")
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))
	fmt.Println( 4 % 2)
}