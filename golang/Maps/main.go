package main

import (
	"fmt"
	"sort"
)



func main(){

// ==============================================
	// 1. BASIC MAP DECLARATION
	// ==============================================

	// Declare a nil map (zero-valued, cannot be used yet)
	var nilMap map[string]int
	fmt.Println("\nNil map:", nilMap, "Length:", len(nilMap))

	// Initialize a map with make()

	ages := make(map[string] int)

	ages["kevin"] = 20
	ages["kamau"] = 30
	fmt.Println("Map with make():", ages, "Length:", len(ages))

	// Initialize a map with values

	salaries := map[string]float64{
		"Alice": 75000.50,
		"Bob":   82000.75,
		"Carol": 68000.00,
	}
	fmt.Println("Initialized map:", salaries, "Length:", len(salaries))

	
	// Mixed value types using interface{}
	person := map[string]interface{}{
		"name":    "John",
		"age":     30,
		"married": true,
		"height":  1.85,
	}
	fmt.Println("Mixed-type map:", person)


	// ==============================================
	// 2. MAP OPERATIONS
	// ==============================================

	// Accessing elements
	fmt.Println("\nAlice's age:", ages["Alice"])
	fmt.Println("Bob's salary:", salaries["Bob"])

	// Checking for existence
	if age, exists := ages["Charlie"]; exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found in ages map")
	}

	// Modifying elements
	salaries["Alice"] = 80000.00
	fmt.Println("Updated Alice's salary:", salaries["Alice"])

	// Adding new elements
	ages["Charlie"] = 28
	fmt.Println("Added Charlie:", ages)

	// Deleting elements
	delete(ages, "Bob")
	fmt.Println("After deleting Bob:", ages)

	// ==============================================
	// 3. MAP PROPERTIES AND BEHAVIOR
	// ==============================================

	// Maps are reference types
	refSalaries := salaries
	refSalaries["Alice"] = 85000.00
	fmt.Println("\nOriginal salaries:", salaries)
	fmt.Println("Reference salaries:", refSalaries)

	// Zero value for missing keys
	fmt.Println("Zero value for missing key:", ages["David"])

	// Maps are not comparable (can't use ==)
	// fmt.Println(salaries == refSalaries) // Compile error!

	// ==============================================
	// 4. ITERATING OVER MAPS
	// ==============================================

	// Random iteration order
	fmt.Println("\nIterating over salaries (random order):")
	for name, salary := range salaries {
		fmt.Printf("%s: $%.2f\n", name, salary)
	}

	// Sorted iteration
	fmt.Println("\nIterating over salaries (sorted order):")
	keys := make([]string, 0, len(salaries))
	for k := range salaries {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: $%.2f\n", k, salaries[k])
	}

	// ==============================================
	// 5. MAPS WITH COMPLEX VALUES
	// ==============================================

	// Map of slices
	departments := map[string][]string{
		"Engineering": {"Alice", "Bob", "Charlie"},
		"Marketing":   {"David", "Eve"},
		"Sales":       {"Frank", "Grace"},
	}
	fmt.Println("\nDepartment employees:")
	for dept, employees := range departments {
		fmt.Printf("%s: %v\n", dept, employees)
	}


	// Map of maps
	graph := map[string]map[string]float64{
		"A": {"B": 5.0, "C": 2.0},
		"B": {"A": 5.0, "D": 1.0},
		"C": {"A": 2.0, "D": 6.0},
		"D": {"B": 1.0, "C": 6.0},
	}
	fmt.Println("\nGraph connections:")
	for node, edges := range graph {
		fmt.Printf("%s -> %v\n", node, edges)
	}


}
