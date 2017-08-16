package main 

import (
  "fmt"
  "math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt( -x ) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x,n,lim float64) float64 {
	
	fmt.Println("x",x)

	//在条件前定义的数据只在if的范围有效
	if v := math.Pow(x,n); v < lim{
		return v
	}

	if y := x; y < 1 {
		fmt.Println("aaaaa")
	} else{
		fmt.Println("y",y)
	}	

	return lim
}
func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}
