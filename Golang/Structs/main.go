package main
import "fmt"

func main(){

		fmt.Println("=== Complete Go Structs Tutorial ===")
	fmt.Printf("From Basics to Advanced Operations\n")

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
		Age int
	}

	// Method 2: Struct with mixed field types

	type Employee struct{
		ID int
		Name string
		Position string
		salary float64
		isManAger bool
	}

	// Method 3: Nested struct

	type Address struct{
		Street string
		Country string
		Zipcode string
		City string
	}

	type Customer struct{
		ID int
		Name string
		Address Address
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
		Age:29,
	}
	fmt.Printf("Person with named fields: %+v\n", p2)

	// Method 3: Without field names (must be in order)
	p3 := Person{"Bob",30}
	fmt.Printf("Person without field names: %+v\n", p3)

	// Method 4: Using new() - returns pointer

	p4 := new(Person)
	p4.name = "kamau"
	p4.Age = 31
	fmt.Printf("Person created with new(): %+v\n", *p4)

	// Method 5: Partial initialization

	p5 := Person{
		name: "diana",

	}

	fmt.Printf("Partially initialized Person: %+v\n", p5)


	// 2.2 NESTED STRUCT INITIALIZATION
	fmt.Println("\n2.2 Nested Struct Initialization:")


	C1 := Customer{
		ID:1001 ,
		Name: "Ndirangu",
		Address: Address{
			Street: "kasarani",
		Country: "kenya",
		Zipcode: "0001-500",
		City: "nairobi",

		},	
	}
	
		fmt.Printf("Customer with nested Address: %+v\n", C1)

		// Initialize nested struct separately
	addr := Address{
		Street:  "456 Oak Ave",
		City:    "Gotham",
		Country: "USA",
	}
	c2 := Customer{ID: 1002, Name: "Wayne Enterprises", Address: addr}
	fmt.Printf("Another customer: %+v\n", c2)


	// ================================================================
	// PART 3: ACCESSING AND MODIFYING STRUCT FIELDS
	// =================================================================
	fmt.Println("\n\n🟡 PART 3: ACCESSING AND MODIFYING STRUCT FIELDS")

	// 3.1 ACCESSING FIELDS
	fmt.Println("\n3.1 Accessing Struct Fields:")

	fmt.Printf("p2.name: %s, p2.Age: %d\n",p2.name,p2.Age)

	// Access nested struct fields
	fmt.Printf("c1.Address.City: %s\n", C1.Address.City)

	// Modify fields
	p2.Age = 31
	fmt.Printf("After modifying p2.Age: %+v\n", p2)

	// Modify nested fields
	C1.Address.Zipcode = "54321"
	fmt.Printf("After modifying c1.Address.ZipCode: %+v\n", C1)

	// 3.3 POINTERS TO STRUCTS
	fmt.Println("\n3.3 Working with Pointers to Structs:")

	// 3.3 POINTERS TO STRUCTS
	fmt.Println("\n3.3 Working with Pointers to Structs:")

	// Create pointer to struct
	p6 := &Person{name: "Eve", Age: 28}
	fmt.Printf("Pointer to struct: %+v\n", *p6)

	// Access fields through pointer (no dereferencing needed)
	p6.Age = 29
	fmt.Printf("After modification through pointer: %+v\n", *p6)


	// ================================================================
	// PART 4: STRUCT METHODS
	// ================================================================




























}