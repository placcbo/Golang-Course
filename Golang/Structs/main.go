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
		address Address
	}


	fmt.Println("Defined Person, Employee, Address, and Customer struct types")

	// ================================================================
	// PART 2: CREATING AND INITIALIZING STRUCTS
	// ================================================================
	fmt.Println("\n\n🔵 PART 2: CREATING AND INITIALIZING STRUCTS")

	// Method 1: Zero-value struct
	var p1 Person
	fmt.Printf("Zero-value Person: %+v\n", p1)

	// Method 2: Using field names

	p2 := Person{
		name:"kevin",
		age:29,
	}
	fmt.Printf("Person with named fields: %+v\n", p2)

	// Method 3: Without field names (must be in order)
	p3 := Person{"Bob",30}
	fmt.Printf("Person without field names: %+v\n", p3)

	// Method 4: Using new() - returns pointer

	p4 := new(Person)
	p4.name = "kamau"
	p4.age = 31
	fmt.Printf("Person created with new(): %+v\n", *p4)

	// Method 5: Partial initialization

	p5 := Person{
		name: "diana",

	}

	fmt.Printf("Partially initialized Person: %+v\n", p5)



























}