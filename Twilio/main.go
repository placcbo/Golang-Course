package main

import "fmt"

func main(){
	messagesFromDoris := [] string{
		"you doing anything later?",
		"Did you get my last message",
		"Don't leave me hunging..",
		"please respond I'm lonely",
	}
	numMessages := float64(len(messagesFromDoris))
	costPerMessage := 0.02
	totalCost := numMessages * costPerMessage
	fmt.Println("you sent a total of", numMessages, "your total cost is",totalCost)
}