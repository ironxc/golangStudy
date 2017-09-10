/*
实现 `WordCount`。
它应当返回一个含有 s 中每个 “词” 个数的 map。
函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败
*/

package main

import (
	"./wc"
	"strings"
)

func WordCount(s string) map[string]int {

	var result = make(map[string]int)
	sarr := strings.Fields(s) 
	for _, key := range sarr {
		if _, ok := result[key]; ok{
			result[key]++
		}else{
			result[key] = 1
		}

	}
	return result
}

func main() {
	
	// fmt.Println(WordCount("hello world world"))
	wc.Test(WordCount)
}