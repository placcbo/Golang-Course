package main

import "fmt"

func main(){
	// ==============================================
	// 1. BASIC ARRAY DECLARATION
	// ==============================================
	
	// Declare an array of 5 integers (zero-valued)

	var numbers [5] int
	fmt.Println("Empty array", numbers)

	// Initialize an array with values

	primes := [6] int {2,3,5,7,11,13}
	fmt.Println("Prime array", primes)

	// Array with mixed types (using interface{})

	mixed := [4] interface{}{"apple",5,true,3.14}

	fmt.Println("mixed array", mixed)
}