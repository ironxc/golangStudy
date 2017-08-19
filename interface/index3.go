
//value, ok = element.(T)获取宾凉element是否是T类型的 
//element是interface变量
package main
 import (
     "fmt"
     "strconv"
 )

 type Element interface{}
 type List [] Element

 type Person struct {
     name string
     age int
 }

 //定义了String方法，实现了fmt.Stringer
 func (p Person) String() string {
     return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
 }

 func main() {
     list := make(List, 3)
     list[0] = 1 // an int
     list[1] = "Hello" // a string
     list[2] = Person{"Dennis", 70}

     for index, element := range list {
         if value, ok := element.(int); ok {
             fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
         } else if value, ok := element.(string); ok {
             fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
         } else if value, ok := element.(Person); ok {
            //所有变量传入了String,而Person调用了自己的String方法   
             fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
         } else {
             fmt.Printf("list[%d] is of a different type\n", index)
         }
     }

     for index, element := range list{
           switch value := element.(type) {
               case int:
                   fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
               case string:
                   fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
               case Person:
                   fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
               default:
                   fmt.Println("list[%d] is of a different type", index)
           }
       }
 }