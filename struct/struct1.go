
package main 

import "fmt"

//定义Human和Student类型

type Human struct {

	name string
	age int
}

type Student struct {

	Human
	speciality string
	age int
}

func print(name string, age int, spe string) {

	fmt.Printf("name is %s, age is %d, speciality is %s\n", name, age, spe)	
}

func main() {

	nick := Student{Human{"Nick",18}, "Computer",100}

	//字段和struct重复，由外到内，优先访问外层的
	print(nick.name, nick.age, nick.speciality)

	//访问内层的字段可以通过匿名字段来访问
	fmt.Println("nick.Human.age is" , nick.Human.age)
	nick.speciality = "Music Guitar"
	nick.age = 20
	fmt.Println("修改信息后")
	
	print(nick.name, nick.age, nick.speciality)

}