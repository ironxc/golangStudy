package main 

import "fmt"

// []T 是一个元素类型为 T 的 slice
// [n]T 是一个有 n 个类型为 T 的值的数组
//数组必须长度，长度是它的类型的一部分，因此数组不能改变长度
// slice的长度是可以改变的，可以认为是动态数组

/*
双引号用来创建可解析的字符串字面量(支持转义，但不能用来引用多行)

反引号用来创建原生的字符串字面量，
这些字符串可能由多行组成(不支持任何转义序列)，
原生的字符串字面量多用于书写多行消息、HTML以及正则表达式

而单引号则用于表示Golang的一个特殊类型：rune，类似其他语言的byte但又不完全一样，
是指：码点字面量（Unicode code point），不做任何转义的原始内容
*/

/*
	slice总是引用了底层的数组
*/
func main() {
	var arr1 = [2]int{1,2}

	var arr2 = [...]byte{'a' ,'e' ,'d' ,'c' ,'B' }
	// var arr2 = [...]byte{"a" ,"e" ,"d" ,"c" ,"B" }//使用双引号会报类型错 

	var arr3 = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	fmt.Println("arr1",arr1)

	fmt.Println("arr2",arr2)
	fmt.Println("arr3",arr3)


	var slice = []int{1,2,3,4,5}
	slice1 := slice

	fmt.Println("slice",slice)
	fmt.Println("slice1",slice1)
	slice[0] = 100;

	//slice是引用类型，slice会指向一个底层的array
	fmt.Println("slice",slice)
	fmt.Println("slice1",slice1)
	fmt.Println("len(slice)",len(slice))	 
	fmt.Println("cap(slice)",cap(slice))	
	
	//array不是引用类型，它的拷贝就是值的拷贝
	arra := [2]int{100, 101}
	arrb := arra
	arra[0] = 99
	fmt.Println("arra", arra)
	fmt.Println("arrb", arrb)
	
	slicem := make([]int,10,20)

 	 fmt.Println("slicem",len(slicem),cap(slicem)) 
}