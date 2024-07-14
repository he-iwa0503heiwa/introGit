package main
import (
	"fmt"
)

func main(){
	num := 0
	for i := 0; i < 2; i++{
		fmt.Println(i)
	}
	
	for {
		if num < 2 {
			fmt.Println(num)
			num++
		}else {
			break
		}
	}
	result := calc(10,5)
	fmt.Println(result)
}

func calc(x,y int)(sum int){
	sum = x + y
	return
}