package main

import "fmt"

//变量在未赋值时，会自动给值，
//bool ==> false , string ==> ""
func main() {
		
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}	