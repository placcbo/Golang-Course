package main

import (
	"fmt"
	"math"
)

func main() {
	// ==============================================
	// 1. BASIC ARRAY DECLARATION
	// ==============================================

	// Declare an array of 5 integers (zero-valued)

	var numbers [5]int
	fmt.Println("Empty array", numbers)

	// Initialize an array with values

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Prime array", primes)

	// Array with mixed types (using interface{})

	mixed := [4]interface{}{"apple", 5, true, 3.14}

	fmt.Println("mixed array", mixed)

	// ==============================================
	// 2. ACCESSING AND MODIFYING ARRAY ELEMENTS
	// ==============================================
	colors := [3]string{"red", "orange", "blue"}
	fmt.Println("\nfirst color", colors[0])
	fmt.Println("Last color", colors[len(colors)-1])

	// Modify an element

	colors[1] = "yellow"
	fmt.Println("Modified colors:", colors)

	// Iterate through an array
	fmt.Println("\nIterating through colors:")

	for i := 0; i < len(colors); i++ {
		fmt.Printf("Index %d: %s\n", i, colors[i])
	}

	// Range-based iteration
	fmt.Println("\nRange iteration:")
	for index, value := range colors {
		fmt.Printf("Index %d: %s\n", index, value)
	}

	// ==============================================
	// 3. MULTI-DIMENSIONAL ARRAYS
	// ==============================================

	// 2D array (matrix)
	var matrix [2][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	fmt.Println("\n2D array:")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}

	// Initialize 2D array directly
	identityMatrix := [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println("\nIdentity matrix:")
	for _, row := range identityMatrix {
		fmt.Println(row)
	}

	// ==============================================
	// 4. ARRAY PROPERTIES AND OPERATIONS
	// ==============================================

	// Array length is part of its type
	fmt.Println("\nArray lengths:")
	fmt.Printf("colors type: %T, length: %d\n", colors, len(colors))
	fmt.Printf("primes type: %T, length: %d\n", primes, len(primes))

	// Arrays are value types (copied when assigned)
	copyColors := colors
	copyColors[0] = "purple"
	fmt.Println("\nOriginal colors:", colors)
	fmt.Println("Copied colors:", copyColors)

	// Comparing arrays
	arr1 := [2]int{1, 2}
	arr2 := [2]int{1, 2}
	arr3 := [2]int{2, 1}
	fmt.Println("\nArray comparisons:")
	fmt.Println("arr1 == arr2:", arr1 == arr2) // true
	fmt.Println("arr1 == arr3:", arr1 == arr3) // false

	// ==============================================
	// 5. ARRAY LIMITATIONS AND WORKAROUNDS
	// ==============================================

	// Fixed size - can't be changed
	// colors = append(colors, "orange") // Compile error!

	// Workaround: Create a new array and copy elements
	biggerColors := [5]string{}
	for i, c := range colors {
		biggerColors[i] = c
	}
	biggerColors[3] = "orange"
	biggerColors[4] = "purple"
	fmt.Println("\nBigger colors array:", biggerColors)

	// ==============================================
	// 6. PRACTICAL ARRAY USAGE
	// ==============================================
	// Days of week
	weekdays := [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println("\nWeekdays:")
	for _, day := range weekdays {
		fmt.Println(day)
	}

	// Game board
	var board [3][3]string
	board[0][0] = "X"
	board[1][1] = "O"
	board[2][2] = "X"
	fmt.Println("\nTic-Tac-Toe board:")
	for _, row := range board {
		fmt.Println(row)
	}

	// Mathematical operations
	vectors := [3][2]float64{
		{1.2, 3.4},
		{5.6, 7.8},
		{9.0, 1.2},
	}
	
	fmt.Println("\nVector magnitudes:")
	for _, v := range vectors {
		mag := math.Sqrt(v[0]*v[0] + v[1]*v[1])
		fmt.Printf("Vector %v: magnitude = %.2f\n", v, mag)
	}

}
