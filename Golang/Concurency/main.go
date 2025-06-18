package main

import (
	"fmt"
)

func SayHello() {
	fmt.Println("Hello from Go routine!")
}

func Greet() {
	fmt.Println("Hello world")
}

func Counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	go SayHello()
	 Counter()
	

}
