package main

import "fmt"

func main(){
	fruits := [] int {2,4,6,8,10}
	for i, fruit := range fruits{
		fmt.Println(i, fruit * fruit)

	}

}
