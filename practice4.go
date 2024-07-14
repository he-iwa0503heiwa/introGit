package main
import(
	"fmt"
)
// type Student struct{
// 	name string
// 	math, english float64
// }
// func (s Student)avg(math, english float64)(avgrlt float64){
// 	avgrlt = (math + english) / 2
// 	return 
// }
type User struct{
	name string
}
func (u User)calc(weight,height float64)(result float64){
	result = weight / height / height * 10000
	return
}
func main(){
	u := User{name:"tanaka"}
	fmt.Println(u.name,u.calc(55,175))
	// s := Student{name:"tanaka"}
	// fmt.Println(s.name,s.avg(98,11))
}