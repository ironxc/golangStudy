//嵌入interface

/*
type Interface interface {
    sort.Interface //嵌入字段sort.Interface
    Push(x interface{}) //a Push method to push elements into the heap
    Pop() interface{} //a Pop elements that pops elements from the heap
}
*/

package main 

import "fmt"

type interface1 interface{

	Add(a int, b int) int
}

//interface2拥有有Add何Print的方法
type interface2 interface{
	interface1//嵌入interface1
	Print(a string) string
}

type Opp struct {
	s string
	a int
	b int
}

//添加Print方法
func (p Opp) Print(a string) string{
		return a+"Method"
}

//添加Add方法
func (p Opp) Add(a int, b int) int{
		return a + b
}

func main() {

    var b interface1
    var c interface2
	a := Opp{ "abcd", 10, 20}

	b = a 
	fmt.Println(b)

	c = a
	fmt.Println(c)

	// fmt.Println(a.Add(a.a, a.b), a .Print(a.s))
	// fmt.Println(b.Add(b.a, b.b), b .Print(b.s))

}