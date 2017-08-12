package main 

import "fmt"

//指针相当于一个变量，但是他存的是值的地址，当然也有自己值的地址（就是地址的地址）
func main() {
	i , j := 42, 2701

	var x *int

	fmt.Println("指针的初始值",x)    //<nil>
	fmt.Println(&x)                  //

	p := &i                  //指向i
	fmt.Println("p指针的值",p)
	fmt.Println("p指针的地址",&p)
	fmt.Println("p指针的值",*&p)
	
	fmt.Println(*p)          // 通过p指针读取
	*p = 21                  //通过指针设置值
	fmt.Println(i)           //查看新值

	p = &j                   //指向j
    *p = *p / 37               //通过指针划分值给j
    fmt.Println(j)            //查看j的新值
}