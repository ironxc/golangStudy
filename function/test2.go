package main

import . "fmt" //点操作可以让你调用这个包的函数时不需要前面的包名

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func add(a int) int {
	a = 100
	return a
}
func main() {
	a , b := split(17);
	Println(a,b)
	Println(split(17))

	c := add(12)
	Println(c)
}