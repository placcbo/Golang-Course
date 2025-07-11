package main

import (
	"fmt"
	"strings"
	"math"
)

	// 1. SIMPLE FUNCTION
// This is the most basic function - no parameters, no return value
// 1. SIMPLE FUNCTION
// This is the most basic function - no parameters, no return value
func greet() {
	fmt.Println("Hello, Go!")
}

// 2. FUNCTION WITH PARAMETERS
	// Functions can take one or more parameters with specified types

	func greetByName(name string) {
	fmt.Printf("Hello, %s!\n", name)
	}

	// 3. FUNCTION WITH RETURN VALUE
	// Functions can return values of specified types
	func add(a int, b int) int {
		return a + b
	}

		// 4. MULTIPLE PARAMETERS OF SAME TYPE
	// When parameters share a type, we can shorten the declaration
	func multiply(a, b int) int {
		return a * b
	}

	// 5. MULTIPLE RETURN VALUES
	// Go functions can return multiple values
	func divide(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	}

	// 6. NAMED RETURN VALUES
	// Return values can be named in the function signature
	func rectangleArea(width, height float64) (area float64) {
		area = width * height
		return // "naked" return - returns the named value
	}

	// 7. VARIADIC FUNCTIONS
	// Functions that accept variable number of arguments
	func average(numbers ...float64) float64 {
		total := 0.0
		for _, num := range numbers {
			total += num
		}
		return total / float64(len(numbers))
	}

	// 10. CLOSURES
	// Functions that remember values from their outer scope
	func makeCounter() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

	// 11. RECURSION
	// Functions that call themselves
	func factorial(n int) int {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	}

	// 12. PRACTICAL EXAMPLE: PALINDROME CHECKER
	func isPalindrome(s string) bool {
		s = strings.ToLower(s)
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}


func main(){
// Calling our simple function
	greet() // Prints: Hello, Go!

	

	greetByName("kevin")
	greetByName("kamau")

	sum := add(5,10)
	fmt.Println("5 + 10 =",sum)





	product := multiply(4, 7)
	fmt.Println("4 * 7 =", product) // Prints: 4 * 7 = 28

	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result) // Prints: 10 / 2 = 5
	}

	area := rectangleArea(4.5, 2.5)
	fmt.Println("Rectangle area:", area) // Prints: Rectangle area: 11.25

	avg := average(1.0, 2.0, 3.0, 4.0, 5.0)
	fmt.Println("Average:", avg) // Prints: Average: 3

	// 8. ANONYMOUS FUNCTIONS
	// Functions without a name that can be assigned to variables
	double := func(x int) int {
		return x * 2
	}
	
	fmt.Println("Double of 5:", double(5)) // Prints: Double of 5: 10


	// 9. IMMEDIATELY INVOKED FUNCTION
	// Anonymous function that executes right after declaration
	func() {
		fmt.Println("I'm running immediately!")
	}() // Prints: I'm running immediately!


	counter := makeCounter()
	fmt.Println(counter()) // Prints: 1
	fmt.Println(counter()) // Prints: 2
	fmt.Println(counter()) // Prints: 3

	
	fmt.Println("Factorial of 5:", factorial(5)) // Prints: Factorial of 5: 120

	fmt.Println("Is 'madam' a palindrome?", isPalindrome("madam"))   // Prints: true
	fmt.Println("Is 'golang' a palindrome?", isPalindrome("golang")) // Prints: false



	// 13. PRACTICAL EXAMPLE: CIRCLE CALCULATIONS
	  circleStats :=   func (radius float64) (area float64, circumference float64) {
		area = math.Pi * radius * radius
		circumference = 2 * math.Pi * radius
		return
	}
	
	r := 5.0
	area, circ := circleStats(r)
	fmt.Printf("Circle with radius %.1f: Area=%.2f, Circumference=%.2f\n",
		r, area, circ) // Prints calculated values


}