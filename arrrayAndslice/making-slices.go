package main 

import "fmt"

// b = b[:cap(b)] // len(b)=5, cap(b)=5
// b = b[1:]      // len(b)=4, cap(b)=4
func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]

	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	for i := 0; i < cap(d); i++ {
		d[i] = i
	}

	fmt.Println(d)
	fmt.Println(c)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s ,len(x) ,cap(x) ,x)	
}