package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	// for ; sum < 1000; {
	// 	sum += sum
	// }

	//当while来用
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// fmt.Println("go语言");

	for {
		fmt.Println("aa")
	}
}