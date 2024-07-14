package main
import (
	"fmt"
)
type Student struct{
	name string
	math,english float64
}

type User struct{
	gender string
	age int
}
func main(){
	var s Student
	s.name = "sato"
	s.math = 90
	s.english = 24
	//↓初期化を上書きすることはできない　新しい変数を宣言
	s1 := Student{name:"fujino"}

	fmt.Println(s)
	fmt.Println(s1)

	var u User
	u.gender = "male"
	u.age = 24
	fmt.Println(u)

	u2 := User{gender:"male",age:24}
	fmt.Println(u2)
}