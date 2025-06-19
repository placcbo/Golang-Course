package main

import "fmt"

func main() {
    // ==============================================
    // 1. BASIC FOR LOOP (C-STYLE)
    // ==============================================
    
    fmt.Println("1. Basic for loop:")
    for i := 0; i < 5; i++ {
        fmt.Println("Iteration", i)
    }
    
    // ==============================================
    // 2. WHILE-LIKE LOOP
    // ==============================================
    
    fmt.Println("\n2. While-like loop:")
    j := 0
    for j < 3 {
        fmt.Println("While iteration", j)
        j++
    }
    
    // ==============================================
    // 3. INFINITE LOOP
    // ==============================================
    
    /*
    fmt.Println("\n3. Infinite loop:")
    k := 0
    for {
        fmt.Println("Infinite iteration", k)
        k++
        if k == 3 {
            break
        }
    }
    */
    // Commented out to prevent actual infinite loop
    
    // ==============================================
    // 4. FOR-RANGE LOOP (ITERATING COLLECTIONS)
    // ==============================================
    
    fmt.Println("\n4. For-range loop:")
    
    // Array iteration
    nums := [3]int{10, 20, 30}
    for index, value := range nums {
        fmt.Printf("nums[%d] = %d\n", index, value)
    }
    
    // Slice iteration
    colors := []string{"red", "green", "blue"}
    for i, color := range colors {
        fmt.Printf("colors[%d] = %s\n", i, color)
    }
    
    // Map iteration
    ages := map[string]int{
        "Alice": 25,
        "Bob":   30,
    }
    for name, age := range ages {
        fmt.Printf("%s is %d years old\n", name, age)
    }
    
    // String iteration (UTF-8 characters)
    str := "Go语言"
    for pos, char := range str {
        fmt.Printf("Character %#U starts at byte position %d\n", char, pos)
    }
    
    // ==============================================
    // 5. LOOP CONTROL STATEMENTS
    // ==============================================
    
    fmt.Println("\n5. Loop control statements:")
    
    // Break example
    fmt.Println("Break example:")
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }
        fmt.Println(i)
    }
    
    // Continue example
    fmt.Println("\nContinue example:")
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue
        }
        fmt.Println(i)
    }
    
    // Labeled break
    fmt.Println("\nLabeled break example:")
    OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("i=%d, j=%d\n", i, j)
            if i == 1 && j == 1 {
                break OuterLoop
            }
        }
    }
    
    // ==============================================
    // 6. PRACTICAL EXAMPLES
    // ==============================================
    
    fmt.Println("\n6. Practical examples:")
    
    // Summing numbers
    numbers := []int{1, 2, 3, 4, 5}
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    fmt.Println("Sum of numbers:", sum)
    
    // Filtering elements
    fmt.Println("\nEven numbers:")
    for _, num := range numbers {
        if num%2 == 0 {
            fmt.Println(num)
        }
    }
    
    // Early exit
    names := []string{"Alice", "Bob", "Charlie", "Dave"}
    found := false
    for _, name := range names {
        if name == "Charlie" {
            found = true
            break
        }
    }
    fmt.Println("\nFound Charlie:", found)
    
    // ==============================================
    // 7. NESTED LOOPS
    // ==============================================
    
    fmt.Println("\n7. Nested loops (multiplication table):")
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("%d x %d = %d\t", i, j, i*j)
        }
        fmt.Println()
    }
}