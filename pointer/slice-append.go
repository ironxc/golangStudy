package main 

import "fmt"

func main() {
	var a[]int

	//添加到空的slice里边去
	a = append(a,0)
	printSlice("a", a)

	a = append(a ,0)
	printSlice("a", a)

	a = append(a, 2,3,4)
	printSlice("a", a)
}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}