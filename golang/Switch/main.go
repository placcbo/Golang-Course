package main

import (
	"fmt"
	"time"
)

func main() {
	// ==============================================
	// 1. BASIC SWITCH STATEMENT
	// ==============================================

	fmt.Println("1. Basic switch statement:")
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Tuesday":
		fmt.Println("Second day of work week")
	case "Wednesday":
		fmt.Println("Midweek")
	case "Thursday":
		fmt.Println("Almost Friday")
	case "Friday":
		fmt.Println("Weekend is coming!")
	default:
		fmt.Println("Weekend day")
	}

	// ==============================================
	// 2. SWITCH WITH MULTIPLE VALUES
	// ==============================================

	fmt.Println("\n2. Switch with multiple values:")
	month := "April"
	switch month {
	case "January", "February", "December":
		fmt.Println("Winter season")
	case "March", "April", "May":
		fmt.Println("Spring season")
	case "June", "July", "August":
		fmt.Println("Summer season")
	case "September", "October", "November":
		fmt.Println("Autumn season")
	default:
		fmt.Println("Unknown month")
	}

	// ==============================================
	// 3. SWITCH WITH EXPRESSIONS
	// ==============================================

	fmt.Println("\n3. Switch with expressions:")
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// ==============================================
	// 4. SWITCH WITH INITIALIZATION
	// ==============================================

	fmt.Println("\n4. Switch with initialization:")
	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday")
	}

	// ==============================================
	// 5. TYPE SWITCH
	// ==============================================

	fmt.Println("\n5. Type switch:")
	var value interface{} = 3.14
	switch v := value.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case float64:
		fmt.Printf("Float: %f\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// ==============================================
	// 6. FALLTHROUGH BEHAVIOR
	// ==============================================

	fmt.Println("\n6. Fallthrough behavior:")
	num := 2
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough // Execution continues to the next case
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other number")
	}

	// ==============================================
	// 7. PRACTICAL EXAMPLES
	// ==============================================

	fmt.Println("\n7. Practical examples:")

	// Grade calculator
	score := 88
	switch {
	case score >= 90:
		fmt.Println("A grade")
	case score >= 80:
		fmt.Println("B grade")
	case score >= 70:
		fmt.Println("C grade")
	case score >= 60:
		fmt.Println("D grade")
	default:
		fmt.Println("F grade")
	}

	// File extension checker
	file := "document.pdf"
	switch ext := file[len(file)-4:]; ext {
	case ".txt":
		fmt.Println("Text file")
	case ".pdf":
		fmt.Println("PDF document")
	case ".jpg", ".png":
		fmt.Println("Image file")
	default:
		fmt.Println("Unknown file type")
	}

	// ==============================================
	// 8. COMPARISON WITH IF-ELSE
	// ==============================================

	fmt.Println("\n8. Comparison with if-else:")
	// This switch:
	color := "blue"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Unknown color")
	}

	// Is equivalent to this if-else:
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Unknown color")
	}

	// ==============================================
	// 9. COMMON PITFALLS
	// ==============================================

	fmt.Println("\n9. Common pitfalls:")

	// No automatic fallthrough (unlike C/C++)
	fmt.Println("No automatic fallthrough:")
	letter := 'a'
	switch letter {
	case 'a':
		fmt.Println("Letter a")
	case 'b':
		fmt.Println("Letter b")
		// No automatic fallthrough to default
	}
	// Outputs only "Letter a"

	// Break is not needed (but can be used to exit early)
	fmt.Println("\nBreak in switch:")
	val := 42
	switch {
	case val > 0:
		fmt.Println("Positive")
		if val == 42 {
			fmt.Println("The answer!")
			break
		}
		fmt.Println("This won't print for 42")
	default:
		fmt.Println("Non-positive")
	}
}