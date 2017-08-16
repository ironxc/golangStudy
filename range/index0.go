package main 

import "fmt"

var pow  = []int{1, 2, 4, 8, 16, 32, 64, 128}
var	pow2 = make([]int, 10)

//for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。
// 通过_符号来忽略index或者value
func main() {
	for i , v := range pow {
		fmt.Printf("pow[%d] = %d\n", i, v)
	}

	for i := range pow2{
		pow2[i] = 1 << uint(i)
	}
	// _ 符号把序号忽略了
	for _, value := range pow2 {
			fmt.Printf("%d\n", value)
	}

	// _ 符号把value值给忽略了
	for index, _ := range pow2 {
			fmt.Printf("%d\n", index)
	}
	fmt.Println("pow2",pow2)
}