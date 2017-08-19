
//value, ok = element.(T)获取宾凉element是否是T类型的
//element是interface变量
package main 

import "fmt"

type Element interface{

}

func main() {
	var a interface{}
	b := 100
	a = b
	fmt.Println(a)
	value, ok := a.(int)

	fmt.Printf("value:%d,ok:%t\n", value, ok)

	val, is := a.(string)		
	fmt.Println(val)	//娶不到值
	fmt.Println(is)	//false
}