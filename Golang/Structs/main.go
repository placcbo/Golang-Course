package main
import "fmt"

func main(){

		fmt.Println("=== Complete Go Structs Tutorial ===")
	fmt.Println("From Basics to Advanced Operations\n")

	// ================================================================
	// PART 1: INTRODUCTION TO STRUCTS
	// ================================================================
	fmt.Println("🟢 PART 1: INTRODUCTION TO STRUCTS")

	fmt.Println("\n1.1 What are Structs?")
	fmt.Println("• Structs are user-defined types that group together variables (fields)")
	fmt.Println("• Each field has a name and a type")
	fmt.Println("• Structs are value types (copied when assigned or passed)")
	fmt.Println("• Used to represent real-world entities with multiple properties")

	// 1.2 DECLARING STRUCTS
	fmt.Println("\n 1.2 Different Ways to Declare Structs:")

	// Method 1: Basic struct definition

	type Person struct{
		name string
		age int
	}

	// Method 2: Struct with mixed field types

	type Employee struct{
		ID int
		Name string
		Position string
		salary float64
		isManager bool
	}

	// Method 3: Nested struct

	type Address struct{
		Street string
		Country string
		Zipcode string
		City string
	}

	type Customer struct{
		ID string
		Name string
		address address
	}




























}