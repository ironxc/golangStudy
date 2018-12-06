package main

import (  
    "fmt"
)

// func assert(i interface{}) {  
//     s := i.(int) //get the underlying int value from i
//     fmt.Println(s)
// }
// func main() {  
//     var s interface{} = 56
//     assert(s)
// }

// func assert(i interface{}) {  
//     s := i.(int) 
//     fmt.Println(s)
// }
// func main() {  
//     var s interface{} = "Steven Paul"
//     assert(s)
// }

type Person struct{
    name string
}

func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
    if(!ok){
      v, ok := i.(Person) 
      fmt.Println(v, ok)             
    }
}
func main() {  
    var p interface{} = Person{"asf"}
    assert(p)
    var s interface{} = 56
    assert(s)
    var i interface{} = "Steven Paul"
    assert(i)
}