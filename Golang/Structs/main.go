package main

import (
	"fmt"
	"sort"
)

func main() {
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
	fmt.Println("\n1.2 Different Ways to Declare Structs:")

	// Method 1: Basic struct definition
	type Person struct {
		Name string
		Age  int
	}

	// Method 2: Struct with mixed field types
	type Employee struct {
		ID        int
		Name      string
		Position  string
		Salary    float64
		IsManager bool
	}

	// Method 3: Nested struct
	type Address struct {
		Street  string
		City    string
		Country string
		ZipCode string
	}

	type Customer struct {
		ID      int
		Name    string
		Address Address
	}

	fmt.Println("Defined Person, Employee, Address, and Customer struct types")

	// ================================================================
	// PART 2: CREATING AND INITIALIZING STRUCTS
	// ================================================================
	fmt.Println("\n\n🔵 PART 2: CREATING AND INITIALIZING STRUCTS")

	// 2.1 CREATING STRUCT INSTANCES
	fmt.Println("\n2.1 Creating Struct Instances:")

	// Method 1: Zero-value struct
	var p1 Person
	fmt.Printf("Zero-value Person: %+v\n", p1)

	// Method 2: Using field names
	p2 := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("Person with named fields: %+v\n", p2)

	// Method 3: Without field names (must be in order)
	p3 := Person{"Bob", 25}
	fmt.Printf("Person without field names: %+v\n", p3)

	// Method 4: Using new() - returns pointer
	p4 := new(Person)
	p4.Name = "Charlie"
	p4.Age = 35
	fmt.Printf("Person created with new(): %+v\n", *p4)

	// Method 5: Partial initialization
	p5 := Person{Name: "Diana"}
	fmt.Printf("Partially initialized Person: %+v\n", p5)

	// 2.2 NESTED STRUCT INITIALIZATION
	fmt.Println("\n2.2 Nested Struct Initialization:")

	// Initialize nested struct
	c1 := Customer{
		ID:   1001,
		Name: "Acme Corp",
		Address: Address{
			Street:  "123 Main St",
			City:    "Metropolis",
			Country: "USA",
			ZipCode: "12345",
		},
	}
	fmt.Printf("Customer with nested Address: %+v\n", c1)

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
	// ================================================================
	fmt.Println("\n\n🟡 PART 3: ACCESSING AND MODIFYING STRUCT FIELDS")

	// 3.1 ACCESSING FIELDS
	fmt.Println("\n3.1 Accessing Struct Fields:")

	// Access fields with dot notation
	fmt.Printf("p2.Name: %s, p2.Age: %d\n", p2.Name, p2.Age)

	// Access nested struct fields
	fmt.Printf("c1.Address.City: %s\n", c1.Address.City)

	// 3.2 MODIFYING FIELDS
	fmt.Println("\n3.2 Modifying Struct Fields:")

	// Modify fields
	p2.Age = 31
	fmt.Printf("After modifying p2.Age: %+v\n", p2)

	// Modify nested fields
	c1.Address.ZipCode = "54321"
	fmt.Printf("After modifying c1.Address.ZipCode: %+v\n", c1)

	// 3.3 POINTERS TO STRUCTS
	fmt.Println("\n3.3 Working with Pointers to Structs:")

	// Create pointer to struct
	p6 := &Person{Name: "Eve", Age: 28}
	fmt.Printf("Pointer to struct: %+v\n", *p6)

	// Access fields through pointer (no dereferencing needed)
	p6.Age = 29
	fmt.Printf("After modification through pointer: %+v\n", *p6)

	// ================================================================
	// PART 4: STRUCT METHODS
	// ================================================================
	fmt.Println("\n\n🟠 PART 4: STRUCT METHODS")

	// 4.1 DEFINING METHODS
	fmt.Println("\n4.1 Defining Methods on Structs:")

	// Method with value receiver
	greet := func(p Person) string {
		return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
	}

	// Method with pointer receiver
	birthday := func(p *Person) {
		p.Age++
		fmt.Printf("Happy Birthday %s! You're now %d\n", p.Name, p.Age)
	}

	// Test methods
	fmt.Println(greet(p2))
	birthday(&p2)
	fmt.Println(greet(p2))

	// 4.2 METHODS ON NESTED STRUCTS
	fmt.Println("\n4.2 Methods on Nested Structs:")

	// Method on Address struct
	formatAddress := func(a Address) string {
		return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.Country, a.ZipCode)
	}

	// Method on Customer that uses Address method
	customerInfo := func(c Customer) string {
		return fmt.Sprintf("%s (%d)\nAddress: %s", c.Name, c.ID, formatAddress(c.Address))
	}

	fmt.Println(customerInfo(c1))

	// ================================================================
	// PART 5: ADVANCED STRUCT OPERATIONS
	// ================================================================
	fmt.Println("\n\n🟣 PART 5: ADVANCED STRUCT OPERATIONS")

	// 5.1 EMBEDDED STRUCTS (COMPOSITION)
	fmt.Println("\n5.1 Embedded Structs (Composition):")

	type ContactInfo struct {
		Email   string
		Phone   string
		Website string
	}

	// Business embeds ContactInfo
	type Business struct {
		Name string
		ContactInfo // Embedded struct
		Revenue     float64
	}

	// Create business
	b1 := Business{
		Name: "Tech Solutions Inc",
		ContactInfo: ContactInfo{
			Email:   "info@techsolutions.com",
			Phone:   "555-123-4567",
			Website: "www.techsolutions.com",
		},
		Revenue: 1000000,
	}

	// Access embedded fields directly
	fmt.Printf("Business email: %s\n", b1.Email)
	fmt.Printf("Full business info: %+v\n", b1)

	// 5.2 STRUCT TAGS
	fmt.Println("\n5.2 Struct Tags:")

	type User struct {
		ID        int    `json:"id" db:"user_id"`
		Username  string `json:"username" db:"username"`
		Password  string `json:"-" db:"password_hash"` // - means omit in JSON
		CreatedAt string `json:"created_at" db:"created_at"`
	}

	u1 := User{
		ID:        1,
		Username:  "johndoe",
		Password:  "secret",
		CreatedAt: "2023-01-01",
	}

	fmt.Printf("User with tags: %+v\n", u1)
	fmt.Println("Note: Tags are accessed via reflection, not shown here")

	// 5.3 ANONYMOUS STRUCTS
	fmt.Println("\n5.3 Anonymous Structs:")

	// One-time use struct
	temp := struct {
		ID   int
		Name string
	}{
		ID:   999,
		Name: "Temporary",
	}

	fmt.Printf("Anonymous struct: %+v\n", temp)

	// ================================================================
	// PART 6: WORKING WITH STRUCT COLLECTIONS
	// ================================================================
	fmt.Println("\n\n🔴 PART 6: WORKING WITH STRUCT COLLECTIONS")

	// 6.1 SLICES OF STRUCTS
	fmt.Println("\n6.1 Slices of Structs:")

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Diana", 28},
	}

	fmt.Println("People slice:")
	for i, p := range people {
		fmt.Printf("%d: %+v\n", i, p)
	}

	// 6.2 SORTING STRUCT SLICES
	fmt.Println("\n6.2 Sorting Struct Slices:")

	// Sort by age (ascending)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by age (ascending):")
	for _, p := range people {
		fmt.Printf("%s: %d\n", p.Name, p.Age)
	}

	// Sort by name (descending)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name > people[j].Name
	})
	fmt.Println("\nSorted by name (descending):")
	for _, p := range people {
		fmt.Printf("%s: %d\n", p.Name, p.Age)
	}

	// 6.3 MAPS WITH STRUCT VALUES
	fmt.Println("\n6.3 Maps with Struct Values:")

	employeeDB := map[int]Employee{
		101: {101, "John Smith", "Developer", 85000, false},
		102: {102, "Sarah Lee", "Manager", 120000, true},
		103: {103, "Mike Brown", "Designer", 75000, false},
	}

	fmt.Println("Employee database:")
	for id, emp := range employeeDB {
		fmt.Printf("%d: %s (%s) - $%.2f\n", id, emp.Name, emp.Position, emp.Salary)
	}

	// 6.4 FILTERING STRUCT SLICES
	fmt.Println("\n6.4 Filtering Struct Slices:")

	// Filter people older than 30
	var olderThan30 []Person
	for _, p := range people {
		if p.Age > 30 {
			olderThan30 = append(olderThan30, p)
		}
	}
	fmt.Println("People older than 30:")
	for _, p := range olderThan30 {
		fmt.Printf("%s: %d\n", p.Name, p.Age)
	}

	// ================================================================
	// PART 7: PRACTICAL STRUCT PATTERNS
	// ================================================================
	fmt.Println("\n\n🎯 PART 7: PRACTICAL STRUCT PATTERNS")

	// 7.1 BUILDER PATTERN
	fmt.Println("\n7.1 Builder Pattern:")

	type Computer struct {
		CPU     string
		RAM     int // GB
		Storage int // GB
		GPU     string
	}

	type ComputerBuilder struct {
		computer Computer
	}

	setCPU := func(b *ComputerBuilder, cpu string) *ComputerBuilder {
		b.computer.CPU = cpu
		return b
	}

	setRAM := func(b *ComputerBuilder, ram int) *ComputerBuilder {
		b.computer.RAM = ram
		return b
	}

	setStorage := func(b *ComputerBuilder, storage int) *ComputerBuilder {
		b.computer.Storage = storage
		return b
	}

	setGPU := func(b *ComputerBuilder, gpu string) *ComputerBuilder {
		b.computer.GPU = gpu
		return b
	}

	buildComputer := func(b *ComputerBuilder) Computer {
		return b.computer
	}

	// Use builder
	builder := ComputerBuilder{}
	myPC := buildComputer(
		setGPU(
			setStorage(
				setRAM(
					setCPU(&builder, "Intel i7"),
					32,
				),
				1000,
			),
			"NVIDIA RTX 3080",
		),
	)

	fmt.Printf("Custom built PC: %+v\n", myPC)

	// 7.2 OPTIONAL FIELDS PATTERN
	fmt.Println("\n7.2 Optional Fields Pattern:")

	type Config struct {
		Host     string
		Port     int
		Timeout  *int    // Optional - nil means not set
		LogLevel *string // Optional
	}

	// Create config with required fields only
	cfg1 := Config{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("Config with required fields: %+v\n", cfg1)

	// Create config with optional fields
	timeout := 30
	logLevel := "debug"
	cfg2 := Config{
		Host:     "example.com",
		Port:     443,
		Timeout:  &timeout,
		LogLevel: &logLevel,
	}
	fmt.Printf("Config with optional fields: %+v\n", cfg2)
	if cfg2.LogLevel != nil {
		fmt.Printf("LogLevel: %s\n", *cfg2.LogLevel)
	}

	// ================================================================
	// PART 8: BEST PRACTICES AND GOTCHAS
	// ================================================================
	fmt.Println("\n\n🏆 PART 8: BEST PRACTICES AND GOTCHAS")

	fmt.Println("\n8.1 Best Practices:")
	fmt.Println("✅ Use descriptive field names")
	fmt.Println("✅ Keep structs focused (Single Responsibility Principle)")
	fmt.Println("✅ Prefer composition over inheritance")
	fmt.Println("✅ Consider making structs immutable where possible")
	fmt.Println("✅ Use constructor functions for complex initialization")
	fmt.Println("✅ Add methods that logically operate on the struct's data")

	fmt.Println("\n8.2 Common Gotchas:")
	fmt.Println("⚠️  Remember that structs are value types (copied on assignment)")
	fmt.Println("⚠️  Zero-value structs have all fields set to their zero values")
	fmt.Println("⚠️  Comparing structs only works if all fields are comparable")
	fmt.Println("⚠️  Be careful with pointer vs value receiver methods")
	fmt.Println("⚠️  Large structs passed by value can impact performance")

	fmt.Println("\n8.3 When to Use Pointers to Structs:")
	fmt.Println("• When you need to modify the struct in a function")
	fmt.Println("• When the struct is large and you want to avoid copying")
	fmt.Println("• When working with interfaces that require pointer receivers")
	fmt.Println("• When the struct needs to be nil-able")

	fmt.Println("\n=== End of Go Structs Tutorial ===")
}