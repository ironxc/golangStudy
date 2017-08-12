package main 

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for index, value := range pow {
		fmt.Println(index)
		fmt.Printf("%d\n", value)
	}
}