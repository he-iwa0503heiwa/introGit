package main 
import (
	"fmt"
)
type Student struct {
	name string
}

func (s Student) calcAvg(data []float64) (avgRlt float64){
	sum := 0.0
	for i := 0; i < len(data); i++{
		sum += data[i]
	}
	avgRlt = sum / float64(len(data))
	return
}
func (s Student) judge(avgRlt float64)(judgeRlt string){
	if avgRlt < 60{
		judgeRlt = "failed"
	}else {
		judgeRlt = "passed"
	}
	return
}

func main(){
	s := Student{name:"suzuki"}
	data := []float64{90,40,45,67,87}
	avgrlt := s.calcAvg(data)
	judgeRlt := s.judge(avgrlt)
	fmt.Println(avgrlt)
	fmt.Println(s.name,judgeRlt)
}