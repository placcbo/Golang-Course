package main
import "fmt"


func main(){
	

	scores := [] float64 {50.5,20.5,20,30.56,10}
	sum := 0
	for _,score := range scores{
		sum += score
	}

	average := sum / len(scores)
	fmt.Println(average)
	

 }