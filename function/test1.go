package main

import "fmt"
func add(x int, y int) int{
	return x + y
}

func addMore(x,y ,z int) int{
	return x + y + z
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(addMore(42, 13, 100))

	a, b := swap("hello", "world")

	fmt.Println(a , b)
	fmt.Println(swap("hello", "world"))
	
}

func swap(a, b string) (string, string){
	return a+b, a
}