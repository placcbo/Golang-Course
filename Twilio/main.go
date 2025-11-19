package main

import "fmt"

func main(){
	messagelength := 10

	maxMessageLength := 20

	fmt.Println("tyring to send a message of length:", messagelength)

	if messagelength <maxMessageLength{
		fmt.Println("Message sent")
	} else{
		fmt.Println("Message not sent")
	}

}