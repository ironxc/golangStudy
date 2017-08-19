
package main 

import "fmt"

//自定义一个person类型的struct
type person struct{
	name string
	age int
}


func Older(p1 , p2 person) (person ,int ){
	if p1.age > p2.age {
		return p1 , p1.age - p2.age
	}

	return p2 , p2.age - p1.age	
}


func main() {

	// var p person

	// p.name = "小明"
	// p.age = 18
	// p := person{ "小明" ,18 }

	// fmt.Printf("name : %s  , age : %d",p.name ,p.age)
	var tom person
	tom.name , tom.age = "Tom" , 22

	bob := person{ age: 18 ,name: "Bob" }

	paul := person{ age: 20 ,name: "Paul"}


	tb_Older, tb_diff := Older(tom ,bob)
	tp_Older, tp_diff := Older(tom ,paul)
	pb_Older, pb_diff := Older(paul ,bob)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
	        tom.name, bob.name, tb_Older.name, tb_diff)

    fmt.Printf("Of %s and %s, %s is older by %d years\n",
        tom.name, paul.name, tp_Older.name, tp_diff)

    fmt.Printf("Of %s and %s, %s is older by %d years\n",
        bob.name, paul.name, pb_Older.name, pb_diff)


}
