package main 

import "fmt"


// Go程序设计的一些规:==>

/*
Go之所以会那么简洁，是因为它有一些默认的行为：
大写字母开头的变量是可导出的，也就是其它包可以读取的，是公有变量；
小写字母开头的就是不可导出的，是私有变量。
大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；
小写字母开头的就是有private关键词的私有函数

*/
//iota默认开始值是0，const中每增加一行加1
const (
	x = iota
	y = iota
	z = iota
)
type M map[string] interface {}

func main() {
	bb := M{ "a": "c"}
	bb["aa"] = "ccc"
	fmt.Println("bb" != "")
	cc, _ := bb["aa"]
	fmt.Println("bb", bb ,cc)
	fmt.Println("x ",x)
	fmt.Println("y ",y)
	fmt.Println("z ",z)
}