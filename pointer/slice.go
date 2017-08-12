// []T 是一个元素类型为 T 的 slice。

package main 

import "fmt"

func main() {
	p := []int{2 ,3 ,5 ,11 ,23 ,1 }
	fmt.Println("p==",p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] ==%d\n ", i, p[i])	
	}

	fmt.Println("p[1:4] ==", p[1:4])//[3,5,11]

	fmt.Println("p[:4] ==", p[:4]) //[2,3,5,11]
	fmt.Println("p[4:] ==", p[4:]) //[23,1]

	fmt.Println("cap(p)",cap(p))
}