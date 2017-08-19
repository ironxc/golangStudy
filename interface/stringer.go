package main



//可以通过interface让函数接受任意类型的参数
//其实是在调用自己的method


//fmt源码中定义了
/*
	type Stringer interface {
	     String() string
	}

*/
import (
    "fmt"
    "strconv"
)

type Human struct {
    name string
    age int
    phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
    return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
    Bob := Human{"Bob", 39, "000-7777-XXX"}
    fmt.Println("This Human is : ", Bob)
}