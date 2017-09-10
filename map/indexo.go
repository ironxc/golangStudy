
// map[keyType]valueType,
//map使用前必须make
package main 

import "fmt"

type Vertex struct {
	Lart, Long float64
}

type Person struct {
	name string
	age int
}
var m map[string]Vertex

func main() {

	var numbers = make(map[string]int)
	numbers["one"] = 100

	nums := make(map[string]int)
	nums ["one"] = 1
	nums ["two"] = 2
	nums ["three"] = 3
	fmt.Println(nums)
	fmt.Println(numbers)

	m = make(map[string]Vertex)
	m["Bell Lab"] = Vertex{
		40.68433, -74.39967,
		//最后一个必须加逗号，或者让}在同一行
	}

	//通过["key"] 取值
	fmt.Println(m["Bell Lab"])

	//声明并初始化值
	var persons = map[string]Person{
		"p1": Person{
			"小明",
			18,
		},
		"p2": Person{
			"小王",
			20,
		},
		// 如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。
		"p3": {
			"小李",
			19,
		},  
	}
	for i,v := range persons {
		fmt.Println(i, v)
	}
	fmt.Printf("%s的年龄是%d\n", persons["p1"].name, persons["p1"].age)
	fmt.Printf("%s的年龄是%d\n", persons["p2"].name, persons["p2"].age)
	fmt.Printf("%s的年龄是%d\n", persons["p3"].name, persons["p3"].age)

}