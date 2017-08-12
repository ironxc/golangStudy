package main 
import "fmt"

func main() {
	c := 257 
	var d int64 = 123242412312324243 
	a := &c
	var b = &d  

	fmt.Println(a,*a)
	fmt.Println(b,*b)

	var n = 12;
	var m = n;//值的复制
	fmt.Println(n)
	fmt.Println("&n",&n)
	fmt.Println(m)
	fmt.Println("&m",&m)
}