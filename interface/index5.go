package main

import (  
    "fmt"
)

type Test interface {  
    Tester()
}

type Person struct {
    name string
}


func (m Person) Tester() {  
    fmt.Println(m)
}

func describe(t Test) {  
    fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {  
    var t Test
    f := Person{"nnn"}
    t = f
    describe(t)
    t.Tester()
}