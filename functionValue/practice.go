package main 

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func(int) []int {

	return func (x int) []int {

		var fbnBox  []int
		for i := 0; i <= x; i++ {

			if i == 0 {
				fbnBox = append(fbnBox,0)
			}else if i == 1 {
				fbnBox = append(fbnBox,1)
			}else{
		        fbnBox = append(fbnBox ,fbnBox[i-1] + fbnBox[i-2])
			}			
		}
		return fbnBox
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}