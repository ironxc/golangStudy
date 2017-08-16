//append() and copy

// 注：append函数会改变slice所引用的数组的内容，从而影响到引用同一数组的其它slice。 
//但当slice中没有剩余空间（即(cap-len) == 0）时，此时将动态分配新的数组空间。
//返回的slice数组指针将指向这个空间，而原数组的内容将保持不变；
//其它引用此数组的slice则不受影响。
package main 

import "fmt"

func main() {

	slice1 := make([]int,2,10)

	fmt.Println("slice1",slice1)

	slice2 := slice1[:]
	slice2 = append(slice1,1,2,100)

	slice3 := slice1

	fmt.Println("slice2",slice2)
	fmt.Println("slice1",slice1)

	fmt.Println("&slice1",&slice1)
	fmt.Println("&slice3",&slice3)

	// slice3 = append(slice3,1,100,123)
	slice3 = append(slice1,1,12,123,12,12,3123,123,123,123,123,123,1)

	
	fmt.Println("slice3",slice3)
	fmt.Println("slice1",slice1)
}
