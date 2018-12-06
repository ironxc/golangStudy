package main

import (  
    "fmt"
    "os"
)

// func main() {  
//     f, err := os.Open("/test.txt")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     fmt.Println(f.Name(), "opened successfully")
// }
type Describer interface{
    Describe() int
}
func main() {  
    var d1 Describer
    if d1 == nil {
        fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
    }
		f, err := os.Open("/test.txt")
    if err, ok := err.(*os.PathError); ok {
				// fmt.Println(err.Err)
        fmt.Println("File at path", err.Path, "failed to open")
        return
    }
    fmt.Println(f.Name(), "opened successfully")
    
}