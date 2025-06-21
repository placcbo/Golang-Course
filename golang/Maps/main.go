package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	// ==============================================
	// 1. BASIC MAP DECLARATION
	// ==============================================

	var nilMap map[string]int
	fmt.Println("\nNil map:", nilMap, "Length:", len(nilMap))

	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	fmt.Println("Map with make():", ages, "Length:", len(ages))

	salaries := map[string]float64{
		"Alice": 75000.50,
		"Bob":   82000.75,
		"Carol": 68000.00,
	}
	fmt.Println("Initialized map:", salaries, "Length:", len(salaries))

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

	fmt.Println("\nAlice's age:", ages["Alice"])
	fmt.Println("Bob's salary:", salaries["Bob"])

	if age, exists := ages["Charlie"]; exists {
		fmt.Println("Charlie's age:", age)
	} else {
		fmt.Println("Charlie not found in ages map")
	}

	salaries["Alice"] = 80000.00
	fmt.Println("Updated Alice's salary:", salaries["Alice"])

	ages["Charlie"] = 28
	fmt.Println("Added Charlie:", ages)

	delete(ages, "Bob")
	fmt.Println("After deleting Bob:", ages)

	// ==============================================
	// 3. MAP PROPERTIES AND BEHAVIOR
	// ==============================================

	refSalaries := salaries
	refSalaries["Alice"] = 85000.00
	fmt.Println("\nOriginal salaries:", salaries)
	fmt.Println("Reference salaries:", refSalaries)

	fmt.Println("Zero value for missing key:", ages["David"])

	// ==============================================
	// 4. ITERATING OVER MAPS
	// ==============================================

	fmt.Println("\nIterating over salaries (random order):")
	for name, salary := range salaries {
		fmt.Printf("%s: $%.2f\n", name, salary)
	}

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

	departments := map[string][]string{
		"Engineering": {"Alice", "Bob", "Charlie"},
		"Marketing":   {"David", "Eve"},
		"Sales":       {"Frank", "Grace"},
	}
	fmt.Println("\nDepartment employees:")
	for dept, employees := range departments {
		fmt.Printf("%s: %v\n", dept, employees)
	}

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

	// ==============================================
	// 6. PRACTICAL MAP USAGE
	// ==============================================

	text := "hello world hello go world go hello"
	words := make(map[string]int)
	for _, word := range strings.Fields(text) {
		words[word]++
	}
	fmt.Println("\nWord frequencies:")
	for word, count := range words {
		fmt.Printf("%s: %d\n", word, count)
	}

	config := map[string]interface{}{
		"debug":      true,
		"logLevel":   "info",
		"maxRetries": 3,
		"timeout":    30.5,
		"allowedIPs": []string{"192.168.1.1", "10.0.0.1"},
	}
	fmt.Println("\nConfiguration:")
	for key, value := range config {
		fmt.Printf("%s: %v (%T)\n", key, value, value)
	}

	vectors := map[string][]float64{
		"v1": {1.2, 3.4},
		"v2": {5.6, 7.8},
		"v3": {9.0, 1.2},
	}
	fmt.Println("\nVector magnitudes:")
	for name, v := range vectors {
		mag := math.Sqrt(v[0]*v[0] + v[1]*v[1])
		fmt.Printf("%s: %v → magnitude = %.2f\n", name, v, mag)
	}
}
