package main

import "fmt"

func main() {
    // ==============================================
    // 1. VARIABLE DECLARATION
    // ==============================================
    
    // Explicit type declaration (verbose)
    var name string = "Alice"
    var age int = 30
    var isProgrammer bool = true
    
    fmt.Println("1. Explicit declarations:")
    fmt.Println(name, age, isProgrammer)
    
    // Type inference (Go determines the type)
    var city = "New York"    // string
    var zipCode = 10001      // int
    var temperature = 72.5   // float64
    
    fmt.Println("\n2. Type inference:")
    fmt.Println(city, zipCode, temperature)
    
    // Multiple declarations
    var (
        firstName, lastName string
        height float64
        employed bool
    )
    
    firstName = "Bob"
    lastName = "Smith"
    height = 5.9
    employed = true
    
    fmt.Println("\n3. Multiple declarations:")
    fmt.Println(firstName, lastName, height, employed)
    
    // ==============================================
    // 2. SHORT VARIABLE DECLARATION
    // ==============================================
    // Inside functions, we can use := for concise declaration
    
    country := "Canada"
    population := 38000000
    cold := true
    
    fmt.Println("\n4. Short declarations:")
    fmt.Println(country, population, cold)
    
    // Multiple short declarations
    x, y, z := 10, 20.5, "Hello"
    
    fmt.Println("\n5. Multiple short declarations:")
    fmt.Println(x, y, z)
    
    // ==============================================
    // 3. VARIABLE SCOPE
    // ==============================================
    
    // Package-level variables (visible throughout package)
    var packageVar = "I'm available to all functions in this package"
    
    fmt.Println("\n6. Package-level variable:")
    fmt.Println(packageVar)
    
    // Block scope demonstration
    {
        innerVar := "I'm only visible in this block"
        fmt.Println("\n7. Inner block variable:")
        fmt.Println(innerVar)
    }
    // fmt.Println(innerVar) // This would cause a compile error
    
    // ==============================================
    // 4. REASSIGNMENT AND SHADOWING
    // ==============================================
    
    score := 100
    fmt.Println("\n8. Original score:", score)
    
    // Reassignment
    score = 200
    fmt.Println("Reassigned score:", score)
    
    // Shadowing (creating a new variable in inner scope)
    if true {
        score := 300  // This is a new variable that shadows the outer one
        fmt.Println("Shadowed score:", score)
    }
    fmt.Println("Outer score after shadowing:", score) // Still 200
    
    // ==============================================
    // 5. SPECIAL VARIABLE CASES
    // ==============================================
    
    // The blank identifier _
    // Used when you need to ignore a value
    _, exists := someFunction()
    fmt.Println("\n9. Blank identifier usage:")
    fmt.Println("Exists:", exists)
    
    // Redeclaration (only allowed in specific cases)
    // When at least one new variable is being declared
    a, b := 1, 2
    fmt.Println("\n10. Before redeclaration:", a, b)
    a, c := 3, 4  // a is being reassigned, c is new
    fmt.Println("After redeclaration:", a, b, c)
    
    // ==============================================
    // 6. CONSTANTS VS VARIABLES
    // ==============================================
    
    const pi = 3.14159
    // pi = 3.14  // Error: cannot assign to constant
    
    fmt.Println("\n11. Constant value:")
    fmt.Println("Pi:", pi)
}

func someFunction() (int, bool) {
    return 42, true
}