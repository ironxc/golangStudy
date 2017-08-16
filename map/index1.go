package main 

import "fmt"
/*
在 map m 中插入或修改一个元素：
m[key] = elem

获得元素：
elem = m[key]

删除元素：
delete(m, key)
*/

func main() {
	// m := make(map[string]int)
	var m = map[string]int{
		"one": 100,
		"two": 200,
	}
	m["one"] = 1;

	fmt.Println("m[\"one\"]",m["one"])
	fmt.Println("m",m)

	delete(m,"one")
	
	fmt.Println("m",m)

	v, ok := m["two"]
	fmt.Println("值：", v, "键是否存在?", ok)

	va, oka := m["t"]
	//不存在值为空，返回false
	fmt.Println("值：", va, "键是否存在?", oka)

}