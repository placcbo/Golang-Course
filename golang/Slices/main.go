package main

import (
	"fmt"
	"math"
)

func main() {
	// ==============================================
	// 1. BASIC SLICE DECLARATION
	// ==============================================

	// Declare a nil slice (zero-valued, no underlying array)
	var nilSlice []int
	fmt.Println("\nNil slice:", nilSlice, "Length:", len(nilSlice), "Capacity:", cap(nilSlice))

	// Initialize a slice with values
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Prime slice:", primes, "Length:", len(primes), "Capacity:", cap(primes))

	// Create a slice with make()
	numbers := make([]int, 3, 5) // length 3, capacity 5
	fmt.Println("Slice with make():", numbers, "Length:", len(numbers), "Capacity:", cap(numbers))

	// Slice with mixed types (using interface{})
	mixed := []interface{}{"apple", 5, true, 3.14}
	fmt.Println("Mixed slice:", mixed)

	// ==============================================
	// 2. SLICE OPERATIONS
	// ==============================================

	// Accessing and modifying elements
	colors := []string{"red", "orange", "blue"}
	fmt.Println("\nFirst color:", colors[0])
	fmt.Println("Last color:", colors[len(colors)-1])

	// Modifying an element
	colors[1] = "yellow"
	fmt.Println("Modified colors:", colors)

	// Appending elements
	colors = append(colors, "green", "purple")
	fmt.Println("After append:", colors, "Length:", len(colors), "Capacity:", cap(colors))

	// Slicing a slice (creates new slice that shares underlying array)
	slice1 := colors[1:3]
	fmt.Println("\nSlice of colors[1:3]:", slice1)

	// Full slice expression [low:high:max] controls capacity
	slice2 := colors[1:3:3]
	fmt.Println("Slice with capacity control:", slice2, "Capacity:", cap(slice2))

	// ==============================================
	// 3. SLICE PROPERTIES AND BEHAVIOR
	// ==============================================

	// Slices are reference types
	refSlice := colors
	refSlice[0] = "pink"
	fmt.Println("\nOriginal colors:", colors)
	fmt.Println("Reference slice:", refSlice)

	// Copying slices (creates independent copy)
	copySlice := make([]string, len(colors))
	copy(copySlice, colors)
	copySlice[0] = "black"
	fmt.Println("\nOriginal colors:", colors)
	fmt.Println("Copied slice:", copySlice)

	// Comparing slices (must be done manually)
	fmt.Println("\nSlice comparison (manual):", equalSlices(colors, refSlice))
	fmt.Println("Slice comparison (manual):", equalSlices(colors, copySlice))

	// ==============================================
	// 4. MULTI-DIMENSIONAL SLICES
	// ==============================================

	// 2D slice (jagged array possible)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}
	fmt.Println("\n2D slice:")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	// Dynamic matrix creation
	rows, cols := 3, 3
	dynamicMatrix := make([][]int, rows)
	for i := range dynamicMatrix {
		dynamicMatrix[i] = make([]int, cols)
	}
	dynamicMatrix[1][1] = 5
	fmt.Println("\nDynamic matrix:")
	for _, row := range dynamicMatrix {
		fmt.Println(row)
	}

	// ==============================================
	// 5. SLICE TRICKS AND USEFUL OPERATIONS
	// ==============================================

	// Removing elements
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Println("\nOriginal letters:", letters)

	// Remove element at index 2
	letters = append(letters[:2], letters[3:]...)
	fmt.Println("After removing index 2:", letters)

	// Inserting elements
	letters = insert(letters, 2, "c")
	fmt.Println("After inserting 'c' at index 2:", letters)

	// Filtering elements
	numbers = []int{1, 2, 3, 4, 5, 6}
	evenNumbers := filterEven(numbers)
	fmt.Println("\nOriginal numbers:", numbers)
	fmt.Println("Even numbers:", evenNumbers)

	// ==============================================
	// 6. PRACTICAL SLICE USAGE
	// ==============================================

	// Days of week (can be extended)
	weekdays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	weekdays = append(weekdays, "Saturday", "Sunday")
	fmt.Println("\nAll weekdays:")
	for _, day := range weekdays {
		fmt.Println(day)
	}

	// Game board (dynamic size)
	size := 4
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
	}
	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"
	board[3][3] = "O"
	fmt.Println("\n4x4 Tic-Tac-Toe board:")
	for _, row := range board {
		fmt.Println(row)
	}

	// Mathematical operations
	vectors := [][]float64{
		{1.2, 3.4},
		{5.6, 7.8},
		{9.0, 1.2},
		{3.14, 2.71}, // Can easily add more vectors
	}
	
	fmt.Println("\nVector magnitudes:")
	for _, v := range vectors {
		mag := math.Sqrt(v[0]*v[0] + v[1]*v[1])
		fmt.Printf("Vector %v: magnitude = %.2f\n", v, mag)
	}
}

// Helper function to compare slices
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Helper function to insert into slice
func insert(slice []string, index int, value string) []string {
	return append(slice[:index], append([]string{value}, slice[index:]...)...)
}

// Helper function to filter even numbers
func filterEven(numbers []int) []int {
	var result []int
	for _, n := range numbers {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}