package main

import "fmt"


func main(){

    fmt.Println("=== Complete Go Pointers Tutorial ===")
    fmt.Println("From Basics to Advanced Pointer Operations\n")

	 // ================================================================
    // PART 1: INTRODUCTION TO POINTERS
    // ================================================================
  fmt.Println("🟢 PART 1: INTRODUCTION TO POINTERS")
    fmt.Println("========================================")
	fmt.Println("\n1.1 What are pointers?")
	 fmt.Println("• A pointer holds the memory address of a variable")
	   fmt.Println("• Use '&' to get the address of a variable")
	    fmt.Println("• Use '*' to dereference (get the value at the address)")
		fmt.Println("• Useful for sharing data, modifying values, and efficient memory usage")
  // 1.2 BASIC POINTER DECLARATION
    fmt.Println("\n1.2 Declaring and Using Pointers:")

	var x int = 10
	var ptr *int = &x
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %d\n", &x)
	fmt.Printf("Pointer ptr: %p\n", ptr)
	fmt.Printf("Value of ptr (*ptr): %d\n", *ptr)

	    // ================================================================
    // PART 2: POINTER OPERATIONS
    //

	   fmt.Println("\n\n🔵 PART 2: POINTER OPERATIONS")
    fmt.Println("========================================")

	  // 2.1 MODIFYING VALUES THROUGH POINTERS
    fmt.Println("\n2.1 Modifying Values via Pointers:")
	fmt.Printf("Before: x = %d\n", x)
	*ptr = 20

	fmt.Printf("After: x = %d (via pointer)\n",x)

	  // 2.2 POINTERS AND FUNCTIONS
    fmt.Println("\n2.2 Passing Pointers to Functions:")




	


}