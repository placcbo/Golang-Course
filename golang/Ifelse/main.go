package main

import "fmt"

func main() {
    // ==============================================
    // 1. BASIC IF STATEMENT
    // ==============================================
    
    fmt.Println("1. Basic if statement:")
    x := 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
    
    // ==============================================
    // 2. IF-ELSE STATEMENT
    // ==============================================
    
    fmt.Println("\n2. If-else statement:")
    y := 3
    if y > 5 {
        fmt.Println("y is greater than 5")
    } else {
        fmt.Println("y is 5 or less")
    }
    
    // ==============================================
    // 3. IF-ELSE IF-ELSE STATEMENT
    // ==============================================
    
    fmt.Println("\n3. If-else if-else statement:")
    score := 85
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: D or F")
    }
    
    // ==============================================
    // 4. VARIABLE DECLARATION IN IF STATEMENT
    // ==============================================
    
    fmt.Println("\n4. Variable declaration in if statement:")
    if z := 15; z > 10 {
        fmt.Println("z is greater than 10")
    } else {
        fmt.Println("z is 10 or less")
    }
    // Note: z is only accessible within the if-else blocks
    
    // ==============================================
    // 5. NESTED IF STATEMENTS
    // ==============================================
    
    fmt.Println("\n5. Nested if statements:")
    age := 25
    hasLicense := true
    if age >= 16 {
        fmt.Println("Old enough to drive")
        if hasLicense {
            fmt.Println("Can drive legally")
        } else {
            fmt.Println("Needs to get a license first")
        }
    } else {
        fmt.Println("Too young to drive")
    }
    
    // ==============================================
    // 6. PRACTICAL EXAMPLES
    // ==============================================
    
    fmt.Println("\n6. Practical examples:")
    
    // Checking for even/odd numbers
    num := 7
    if num%2 == 0 {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
    
    // Checking multiple conditions
    username := "admin"
    password := "secret"
    if username == "admin" && password == "secret" {
        fmt.Println("\nLogin successful")
    } else {
        fmt.Println("\nInvalid credentials")
    }
    
    // Checking for zero values
    var emptyString string
    if emptyString == "" {
        fmt.Println("\nemptyString is zero value")
    }
    
    // ==============================================
    // 7. COMMON PITFALLS
    // ==============================================
    
    fmt.Println("\n7. Common pitfalls:")
    
    // Using assignment (=) instead of comparison (==)
    /*
    if x = 5 {  // This would cause a compilation error
        fmt.Println("This won't work")
    }
    */
    
    // Proper way:
    if x == 5 {
        fmt.Println("This is the correct comparison")
    }
    
    // Curly brace placement
    /*
    if x > 5 
    {  // This would cause a compilation error
        fmt.Println("Wrong brace placement")
    }
    */
    
    // Proper way:
    if x > 5 {
        fmt.Println("Correct brace placement")
    }
    
    // ==============================================
    // 8. TERNARY OPERATOR ALTERNATIVE
    // ==============================================
    
    fmt.Println("\n8. Ternary operator alternative:")
    // Go doesn't have ternary operator, but you can use if-else
    status := "active"
    result := "can login"
    if status != "active" {
        result = "cannot login"
    }
    fmt.Println("User", result)
    
    // Alternative for simple assignments
    max := x
    if y > x {
        max = y
    }
    fmt.Println("Maximum value:", max)
}