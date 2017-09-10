/*

练习：Stringers
让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址。

例如，`IPAddr{1,`2,`3,`4}` 应当输出 `"1.2.3.4"`。
*/
package main


import (
 "fmt"
"strconv"
)
type IPAddr [4]byte



func (m IPAddr) String() string {
	str := ""
	for i, v := range m {

		if(i >= len(m)-1){
			str = str+ strconv.Itoa(int(v))
		}else{
			str = str+ strconv.Itoa(int(v)) +"."

		}
	}

	return str
}
// TODO: Add a "String() string" method to IPAddr.

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	fmt.Println(addrs)

	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
		
	}
}