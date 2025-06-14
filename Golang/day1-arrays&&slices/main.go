package main

import (
	"fmt"
)

func main() {
	fmt.Println("==complete go arrays tutorial==")
	fmt.Println("from basic arrays (1D) to 2D arrays")

	// PART 1: BASIC ARRAYS (1D)
	fmt.Println("1.1: declaring arrays")

	var numbers [5]int
	fmt.Println("Empty int array:", numbers)
	var names [3]string
	fmt.Println("Empty string array:", names)

	fruits := [3]string{"mango", "banana", "orange"}
	fmt.Printf("initialised array %v\n", fruits)
	colors := [...]string{"red", "yellow", "green"}
	fmt.Printf("Auto-sized array: %v\n", colors)

	fmt.Println("Accessing and modifying elements")
	fmt.Printf("first fruit: %v\n", fruits[0])
	fmt.Println("last fruit:", fruits[len(fruits)-1])

	for i := 0; i < len(numbers); i++ {
		numbers[i] = (i + 1) * 10
	}
	fmt.Println("filled number array", numbers)

	fmt.Println("**Array properties**")
	fmt.Println("length of fruits array:", len(fruits))
	fmt.Println("length of numbers array:", len(fruits))

	fmt.Println("looping thru arrays")
	for i := 0; i < len(colors); i++ {
		fmt.Println("index", i, ":", colors[i])
	}

	for index, color := range colors {
		fmt.Println("index ", index, ":", color)
	}

	for _, fruit := range fruits {
		fmt.Println("fruit:", fruit)
	}

	scores := [5]int{10, 22, 80, 75, 90}
	max := scores[0]
	for _, value := range scores {
		if value > max {
			max = value
		}
	}
	fmt.Println("max score:", max)

	sum := 0
	for _, score := range scores {
		sum = sum + score
	}
	fmt.Println("Average is:", float64(sum)/float64(len(scores)))

	searchFor := 92
	found := false
	for index, score := range scores {
		if score == searchFor {
			found = true
			fmt.Println("found", score, "at index", index)
			break
		}
	}
	if !found {
		fmt.Println("could not find", searchFor)
	}

	arr1 := [3]int{3, 6, 8}
	arr2 := [3]int{3, 6, 8}
	arr3 := [3]int{3, 7, 8}
	fmt.Println("arr1 == arr2 ?", arr1 == arr2)
	fmt.Println("arr1 == arr3 ?", arr1 == arr3)

	fmt.Println("\n\n Part2: Arrays vs Slices")
	fixedArray := [3]int{1, 2, 3}
	fmt.Println("\nArray has a fixed size, length:", len(fixedArray))
	dynamicSlice := []int{1, 2, 3}
	fmt.Printf("Slice (dynamic): %v, length: %d\n", dynamicSlice, len(dynamicSlice))
	dynamicSlice = append(dynamicSlice, 4)
	fmt.Println("\n appended dynamicSlice is", dynamicSlice)

	fmt.Println("\n\n PART 3: 2D ARRAYS")
	fmt.Println("\n\n Declaring 2D Arrays:")
	var grid [3][4]int
	fmt.Println("\n Empty 3 * 4 grid", grid)

	matrix := [3][4]int{{1, 2, 3, 5}, {5, 5, 4, 7}, {9, 6, 6, 3}}
	fmt.Println("\n\n initialized 3D matrix:", matrix)

	fmt.Println("\n3.2 Understanding 2D Array Structure:")
	fmt.Printf("Number of rows: %d\n", len(matrix))
	fmt.Printf("Number of columns: %d\n", len(matrix[0]))
	fmt.Println("Visual representation:")
	fmt.Println("  Col0 Col1 Col2")
	for i, row := range matrix {
		fmt.Printf("Row%d: %d   %d   %d\n", i, row[0], row[1], row[2])
	}

	fmt.Println("\n3.3 Accessing 2D Array Elements:")
	fmt.Printf("Element at [0][1] (row 0, col 1): %d\n", matrix[0][1])
	fmt.Printf("Element at [1][2] (row 1, col 2): %d\n", matrix[1][2])
	matrix[0][0] = 10
	fmt.Printf("After changing [0][0] to 10: %v\n", matrix)

	fmt.Println("\n3.4 Different Initialization Methods:")
	var zeros [2][3]int
	fmt.Printf("Default zeros: %v\n", zeros)

	partial := [3][2]int{{1, 2}, {3}, {}}
	fmt.Printf("Partial init: %v\n", partial)
	studentGrades := [3][4]int{{85, 90, 78, 92}, {88, 76, 95, 89}, {91, 87, 83, 94}}
	fmt.Printf("Student grades: %v\n", studentGrades)

	fmt.Println("\n3.5 Looping Through 2D Arrays:")
	fmt.Println("Method 1 - Nested for loops:")
	for i := 0; i < len(studentGrades); i++ {
		fmt.Printf("Student %d grades: ", i+1)
		for j := 0; j < len(studentGrades[i]); j++ {
			fmt.Printf("%d ", studentGrades[i][j])
		}
		fmt.Println()
	}
	fmt.Println("\nMethod 2 - Using range (recommended):")
	for studentIndex, grades := range studentGrades {
		fmt.Printf("Student %d: ", studentIndex+1)
		for _, grade := range grades {
			fmt.Printf("%d ", grade)
		}
		fmt.Println()
	}

	fmt.Println("\n3.6 Practical Example - Game Board:")
	var gameBoard [4][4]string
	for i := range gameBoard {
		for j := range gameBoard[i] {
			gameBoard[i][j] = "."
		}
	}
	gameBoard[0][0] = "X"
	gameBoard[1][1] = "O"
	gameBoard[2][2] = "X"
	gameBoard[3][3] = "O"
	fmt.Println("Game Board:")
	for _, row := range gameBoard {
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}

	fmt.Println("\n3.7 2D Array Operations:")
	studentScores := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	totalSum := 0
	for _, grades := range studentScores {
		for _, grade := range grades {
			totalSum += grade
		}
	}
	fmt.Printf("Sum of all grades: %d\n", totalSum)

	maxGrade := studentScores[0][0]
	var maxStudent, maxSubject int
	for i, grades := range studentScores {
		for j, grade := range grades {
			if grade > maxGrade {
				maxGrade = grade
				maxStudent = i
				maxSubject = j
			}
		}
	}
	fmt.Printf("Highest grade: %d (Student %d, Subject %d)\n", maxGrade, maxStudent+1, maxSubject+1)

	fmt.Println("\nEach student's total:")
	for i, grades := range studentScores {
		studentSum := 0
		for _, grade := range grades {
			studentSum += grade
		}
		fmt.Printf("Student %d total: %d\n", i+1, studentSum)
	}

	fmt.Println("\n\n🟣 PART 4: ADVANCED CONCEPTS")
	fmt.Println("=" + fmt.Sprintf("%40s", "="))

	fmt.Println("\n4.1 Passing Arrays to Functions:")
	testArray := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Original array: %v\n", testArray)
	processedSum := sumArray(testArray)
	fmt.Printf("Sum calculated by function: %d\n", processedSum)
	fmt.Printf("Original array unchanged: %v\n", testArray)

	fmt.Println("\n4.2 Arrays of Different Types:")
	flags := [4]bool{true, false, true, false}
	fmt.Printf("Boolean array: %v\n", flags)
	temperatures := [7]float64{23.5, 25.0, 22.8, 26.3, 24.1, 23.9, 25.5}
	fmt.Printf("Temperature array: %v\n", temperatures)

	fmt.Println("\n\n📋 SUMMARY AND BEST PRACTICES")
	fmt.Println("=" + fmt.Sprintf("%50s", "="))
	fmt.Println("✅ Arrays have FIXED size determined at compile time")
	fmt.Println("✅ Use [size]type for 1D arrays: [5]int")
	fmt.Println("✅ Use [rows][cols]type for 2D arrays: [3][4]int")
	fmt.Println("✅ Access elements with array[index] or array[row][col]")
	fmt.Println("✅ Use range for idiomatic looping")
	fmt.Println("✅ Arrays are passed by VALUE (copied) to functions")
	fmt.Println("✅ Consider slices ([]type) for dynamic sizing")
	fmt.Println("✅ Zero values: 0 for numbers, \"\" for strings, false for bools")
	fmt.Println("\n🎯 When to use Arrays vs Slices:")
	fmt.Println("\u2022 Arrays: Fixed size known at compile time, small collections")
	fmt.Println("\u2022 Slices: Dynamic size, most common choice in Go")
}

func sumArray(arr [5]int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}