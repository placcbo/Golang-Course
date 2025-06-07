package main

import (
    "fmt"
    "sort"
    "strings"
)

func main() {
    fmt.Println("=== Complete Go Maps Tutorial ===")
    fmt.Println("From Basics to Advanced CRUD Operations\n")

    // ================================================================
    // PART 1: INTRODUCTION TO MAPS
    // ================================================================
    fmt.Println("🟢 PART 1: INTRODUCTION TO MAPS")
 

    fmt.Println("\n1.1 What are Maps?")
    fmt.Println("• Maps are key-value pairs (like dictionaries/hash tables)")
    fmt.Println("• Keys must be comparable types (strings, numbers, booleans)")
    fmt.Println("• Values can be any type")
    fmt.Println("• Unordered collection (iteration order not guaranteed)")
    fmt.Println("• Reference type (like slices)")

    // 1.2 DECLARING MAPS
    fmt.Println("\n1.2 Different Ways to Declare Maps:")
    
    // Method 1: Using make()
    ages := make(map[string]int)
    fmt.Printf("Empty map with make(): %v\n", ages)
    
    // Method 2: Map literal with values
    scores := map[string]int{
        "Alice":   95,
        "Bob":     87,
        "Charlie": 92,
    }
    fmt.Printf("Map literal: %v\n", scores)
    
    // Method 3: Empty map literal
    inventory := map[string]int{}
    fmt.Printf("Empty map literal: %v\n", inventory)
    
    // Method 4: Var declaration (creates nil map)
    var nilMap map[string]int
    fmt.Printf("Nil map: %v, nil=%t\n", nilMap, nilMap == nil)
    fmt.Println("⚠️  Cannot add to nil map - will panic!")

    // 1.3 MAP PROPERTIES
    fmt.Println("\n1.3 Map Properties:")
    fmt.Printf("Length of scores map: %d\n", len(scores))
    fmt.Printf("Length of empty map: %d\n", len(ages))
    fmt.Printf("Length of nil map: %d\n", len(nilMap))

    // ================================================================
    // PART 2: CRUD OPERATIONS - CREATE
    // ================================================================
    fmt.Println("\n\n🔵 PART 2: CRUD OPERATIONS - CREATE")


    // 2.1 ADDING ELEMENTS (CREATE)
    fmt.Println("\n2.1 Adding Elements to Maps:")
    
    // Create a student database
    students := make(map[int]string) // ID -> Name
    
    // Add individual students
    students[101] = "John Doe"
    students[102] = "Jane Smith"
    students[103] = "Mike Johnson"
    
    fmt.Printf("Students after adding: %v\n", students)
    
    // Add multiple elements
    newStudents := map[int]string{
        104: "Sarah Wilson",
        105: "Tom Brown",
        106: "Lisa Davis",
    }
    
    // Merge maps (manual way)
    for id, name := range newStudents {
        students[id] = name
    }
    fmt.Printf("After merging: %v\n", students)

    // 2.2 DIFFERENT KEY-VALUE TYPES
    fmt.Println("\n2.2 Maps with Different Types:")
    
    // String to slice
    courseStudents := map[string][]string{
        "Math":    {"Alice", "Bob", "Charlie"},
        "Science": {"Alice", "Diana", "Eve"},
        "History": {"Bob", "Charlie", "Frank"},
    }
    fmt.Printf("Course students: %v\n", courseStudents)
    
    // String to map (nested maps)
    studentGrades := map[string]map[string]int{
        "Alice": {
            "Math":    95,
            "Science": 87,
            "History": 92,
        },
        "Bob": {
            "Math":    78,
            "Science": 85,
            "History": 88,
        },
    }
    fmt.Printf("Student grades: %v\n", studentGrades)
    
    // Complex struct as value
    type Person struct {
        Name string
        Age  int
        City string
    }
    
    people := map[string]Person{
        "p001": {"John Doe", 30, "New York"},
        "p002": {"Jane Smith", 25, "Los Angeles"},
    }
    fmt.Printf("People map: %v\n", people)

    // ================================================================
    // PART 3: CRUD OPERATIONS - READ
    // ================================================================
    fmt.Println("\n\n🟡 PART 3: CRUD OPERATIONS - READ")
  

    // 3.1 ACCESSING VALUES
    fmt.Println("\n3.1 Reading Values from Maps:")
    
    products := map[string]float64{
        "laptop":     999.99,
        "mouse":      29.99,
        "keyboard":   79.99,
        "monitor":    299.99,
        "headphones": 149.99,
    }
    
    // Simple access
    fmt.Printf("Laptop price: $%.2f\n", products["laptop"])
    
    // Safe access with "comma ok" idiom
    price, exists := products["tablet"]
    if exists {
        fmt.Printf("Tablet price: $%.2f\n", price)
    } else {
        fmt.Println("Tablet not found in inventory")
    }
    
    // Multiple safe accesses
    items := []string{"laptop", "tablet", "mouse", "phone"}
    fmt.Println("\nChecking multiple items:")
    for _, item := range items {
        if price, found := products[item]; found {
            fmt.Printf("✓ %s: $%.2f\n", item, price)
        } else {
            fmt.Printf("✗ %s: Not available\n", item)
        }
    }

    // 3.2 READING FROM NESTED MAPS
    fmt.Println("\n3.2 Reading from Nested Maps:")
    
    // Safe access to nested maps
    if aliceGrades, exists := studentGrades["Alice"]; exists {
        if mathGrade, exists := aliceGrades["Math"]; exists {
            fmt.Printf("Alice's Math grade: %d\n", mathGrade)
        }
    }
    
    // Function to safely get nested value
    getGrade := func(student, subject string) (int, bool) {
        if grades, exists := studentGrades[student]; exists {
            if grade, exists := grades[subject]; exists {
                return grade, true
            }
        }
        return 0, false
    }
    
    if grade, found := getGrade("Bob", "Science"); found {
        fmt.Printf("Bob's Science grade: %d\n", grade)
    }

    // 3.3 CHECKING FOR KEY EXISTENCE
    fmt.Println("\n3.3 Different Ways to Check Key Existence:")
    
    inventory2 := map[string]int{
        "apples":  50,
        "bananas": 30,
        "oranges": 0, // Zero value exists!
    }
    
    // Method 1: Comma ok idiom (recommended)
    if count, exists := inventory2["apples"]; exists {
        fmt.Printf("Apples in stock: %d\n", count)
    }
    
    // Method 2: Check zero value carefully
    orangeCount := inventory2["oranges"] // Returns 0 even if key exists
    fmt.Printf("Orange count (could be zero or missing): %d\n", orangeCount)
    
    // Better way for zero values
    if count, exists := inventory2["oranges"]; exists {
        fmt.Printf("Oranges in stock (confirmed): %d\n", count)
    } else {
        fmt.Println("Oranges not in inventory")
    }

    // ================================================================
    // PART 4: CRUD OPERATIONS - UPDATE
    // ================================================================
    fmt.Println("\n\n🟠 PART 4: CRUD OPERATIONS - UPDATE")
 

    // 4.1 UPDATING VALUES
    fmt.Println("\n4.1 Updating Map Values:")
    
    bankAccounts := map[string]float64{
        "acc001": 1000.50,
        "acc002": 2500.75,
        "acc003": 750.25,
    }
    
    fmt.Printf("Before update: %v\n", bankAccounts)
    
    // Simple update
    bankAccounts["acc001"] = 1200.50
    fmt.Printf("After updating acc001: %v\n", bankAccounts)
    
    // Conditional update
    accountID := "acc002"
    if balance, exists := bankAccounts[accountID]; exists {
        bankAccounts[accountID] = balance + 500.00 // Add $500
        fmt.Printf("Added $500 to %s: $%.2f\n", accountID, bankAccounts[accountID])
    }
    
    // Update or create pattern
    updateAccount := func(id string, amount float64) {
        if currentBalance, exists := bankAccounts[id]; exists {
            bankAccounts[id] = currentBalance + amount
            fmt.Printf("Updated %s: $%.2f\n", id, bankAccounts[id])
        } else {
            bankAccounts[id] = amount
            fmt.Printf("Created %s: $%.2f\n", id, bankAccounts[id])
        }
    }
    
    updateAccount("acc003", -100.00) // Withdraw $100
    updateAccount("acc004", 3000.00) // Create new account

    // 4.2 UPDATING COMPLEX VALUES
    fmt.Println("\n4.2 Updating Complex Values:")
    
    // Update slice values in map
    fmt.Printf("Before updating course students: %v\n", courseStudents["Math"])
    courseStudents["Math"] = append(courseStudents["Math"], "Grace")
    fmt.Printf("After adding Grace to Math: %v\n", courseStudents["Math"])
    
    // Update nested map
    if bobGrades, exists := studentGrades["Bob"]; exists {
        bobGrades["Math"] = 85 // Update Bob's Math grade
        fmt.Printf("Updated Bob's grades: %v\n", studentGrades["Bob"])
    }
    
    // Add new subject to existing student
    if aliceGrades, exists := studentGrades["Alice"]; exists {
        aliceGrades["Physics"] = 91
        fmt.Printf("Added Physics to Alice: %v\n", studentGrades["Alice"])
    }

    // 4.3 BULK UPDATES
    fmt.Println("\n4.3 Bulk Update Operations:")
    
    priceUpdates := map[string]float64{
        "laptop":   899.99, // Price drop
        "mouse":    24.99,  // Price drop
        "speaker":  199.99, // New product
    }
    
    fmt.Println("Applying bulk price updates:")
    for product, newPrice := range priceUpdates {
        oldPrice, existed := products[product]
        products[product] = newPrice
        
        if existed {
            fmt.Printf("Updated %s: $%.2f → $%.2f\n", product, oldPrice, newPrice)
        } else {
            fmt.Printf("Added %s: $%.2f\n", product, newPrice)
        }
    }

    // ================================================================
    // PART 5: CRUD OPERATIONS - DELETE
    // ================================================================
    fmt.Println("\n\n🔴 PART 5: CRUD OPERATIONS - DELETE")
  

    // 5.1 DELETING ELEMENTS
    fmt.Println("\n5.1 Deleting Map Elements:")
    
    colors := map[string]string{
        "red":    "#FF0000",
        "green":  "#00FF00",
        "blue":   "#0000FF",
        "yellow": "#FFFF00",
        "purple": "#800080",
    }
    
    fmt.Printf("Before deletion: %v\n", colors)
    
    // Delete single element
    delete(colors, "yellow")
    fmt.Printf("After deleting yellow: %v\n", colors)
    
    // Safe deletion (won't panic if key doesn't exist)
    delete(colors, "orange") // Key doesn't exist - no error
    fmt.Printf("After attempting to delete non-existent key: %v\n", colors)

    // 5.2 CONDITIONAL DELETION
    fmt.Println("\n5.2 Conditional Deletion:")
    
    userSessions := map[string]int{
        "user1": 30,  // minutes
        "user2": 120,
        "user3": 5,
        "user4": 90,
        "user5": 2,
    }
    
    fmt.Printf("Before cleanup: %v\n", userSessions)
    
    // Delete sessions shorter than 10 minutes
    minDuration := 10
    var toDelete []string
    
    for user, duration := range userSessions {
        if duration < minDuration {
            toDelete = append(toDelete, user)
        }
    }
    
    for _, user := range toDelete {
        fmt.Printf("Removing short session: %s (%d min)\n", user, userSessions[user])
        delete(userSessions, user)
    }
    
    fmt.Printf("After cleanup: %v\n", userSessions)

    // 5.3 CLEARING ENTIRE MAP
    fmt.Println("\n5.3 Clearing Maps:")
    
    tempMap := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }
    
    fmt.Printf("Before clearing: %v\n", tempMap)
    
    // Method 1: Delete all keys individually
    for key := range tempMap {
        delete(tempMap, key)
    }
    fmt.Printf("After manual clear: %v\n", tempMap)
    
    // Method 2: Reassign to new empty map (more efficient)
    tempMap2 := map[string]int{"x": 1, "y": 2, "z": 3}
    fmt.Printf("Before reassignment: %v\n", tempMap2)
    tempMap2 = make(map[string]int) // New empty map
    fmt.Printf("After reassignment: %v\n", tempMap2)

    // ================================================================
    // PART 6: ADVANCED MAP OPERATIONS
    // ================================================================
    fmt.Println("\n\n🟣 PART 6: ADVANCED MAP OPERATIONS")
    

    // 6.1 ITERATING THROUGH MAPS
    fmt.Println("\n6.1 Different Ways to Iterate:")
    
    cityPopulation := map[string]int{
        "New York":    8400000,
        "Los Angeles": 3900000,
        "Chicago":     2700000,
        "Houston":     2300000,
        "Phoenix":     1700000,
    }
    
    // Method 1: Key and value
    fmt.Println("Cities and populations:")
    for city, population := range cityPopulation {
        fmt.Printf("%s: %,d\n", city, population)
    }
    
    // Method 2: Keys only
    fmt.Println("\nJust city names:")
    for city := range cityPopulation {
        fmt.Printf("- %s\n", city)
    }
    
    // Method 3: Values only (less common)
    fmt.Println("\nJust populations:")
    for _, population := range cityPopulation {
        fmt.Printf("%,d ", population)
    }
    fmt.Println()

    // 6.2 SORTING MAP DATA
    fmt.Println("\n6.2 Sorting Map Data:")
    
    // Maps are unordered, but we can sort for display
    
    // Sort by keys
    var cities []string
    for city := range cityPopulation {
        cities = append(cities, city)
    }
    sort.Strings(cities)
    
    fmt.Println("Cities sorted alphabetically:")
    for _, city := range cities {
        fmt.Printf("%s: %,d\n", city, cityPopulation[city])
    }
    
    // Sort by values (more complex)
    type CityPop struct {
        Name       string
        Population int
    }
    
    var cityList []CityPop
    for city, pop := range cityPopulation {
        cityList = append(cityList, CityPop{city, pop})
    }
    
    // Sort by population (descending)
    sort.Slice(cityList, func(i, j int) bool {
        return cityList[i].Population > cityList[j].Population
    })
    
    fmt.Println("\nCities sorted by population (descending):")
    for _, city := range cityList {
        fmt.Printf("%s: %,d\n", city.Name, city.Population)
    }

    // 6.3 MAP FILTERING AND TRANSFORMATION
    fmt.Println("\n6.3 Filtering and Transforming Maps:")
    
    allScores := map[string]int{
        "Alice":   95,
        "Bob":     67,
        "Charlie": 89,
        "Diana":   72,
        "Eve":     91,
        "Frank":   58,
    }
    
    // Filter: Students with scores >= 70
    passingStudents := make(map[string]int)
    for name, score := range allScores {
        if score >= 70 {
            passingStudents[name] = score
        }
    }
    fmt.Printf("Passing students (≥70): %v\n", passingStudents)
    
    // Transform: Convert to letter grades
    letterGrades := make(map[string]string)
    for name, score := range allScores {
        var grade string
        switch {
        case score >= 90:
            grade = "A"
        case score >= 80:
            grade = "B"
        case score >= 70:
            grade = "C"
        case score >= 60:
            grade = "D"
        default:
            grade = "F"
        }
        letterGrades[name] = grade
    }
    fmt.Printf("Letter grades: %v\n", letterGrades)

    // 6.4 MAP AGGREGATIONS
    fmt.Println("\n6.4 Map Aggregations and Statistics:")
    
    salesData := map[string][]float64{
        "Q1": {1200.50, 1500.75, 980.25, 1750.00},
        "Q2": {1400.00, 1650.25, 1100.50, 1850.75},
        "Q3": {1300.25, 1580.00, 1050.75, 1900.50},
        "Q4": {1600.75, 1720.25, 1250.00, 2100.00},
    }
    
    fmt.Println("Quarterly sales analysis:")
    yearTotal := 0.0
    for quarter, sales := range salesData {
        quarterTotal := 0.0
        for _, sale := range sales {
            quarterTotal += sale
        }
        average := quarterTotal / float64(len(sales))
        yearTotal += quarterTotal
        
        fmt.Printf("%s: Total=$%.2f, Average=$%.2f\n", quarter, quarterTotal, average)
    }
    fmt.Printf("Year total: $%.2f\n", yearTotal)

    // ================================================================
    // PART 7: PRACTICAL EXAMPLES AND PATTERNS
    // ================================================================
    fmt.Println("\n\n🎯 PART 7: PRACTICAL EXAMPLES")
    

    // 7.1 CACHING PATTERN
    fmt.Println("\n7.1 Caching Pattern:")
    
    cache := make(map[string]string)
    
    expensiveOperation := func(input string) string {
        fmt.Printf("Computing result for: %s\n", input)
        return strings.ToUpper(input) + "_PROCESSED"
    }
    
    cachedOperation := func(input string) string {
        if result, exists := cache[input]; exists {
            fmt.Printf("Cache hit for: %s\n", input)
            return result
        }
        
        result := expensiveOperation(input)
        cache[input] = result
        return result
    }
    
    // Test caching
    fmt.Println(cachedOperation("hello"))  // Cache miss
    fmt.Println(cachedOperation("world"))  // Cache miss
    fmt.Println(cachedOperation("hello"))  // Cache hit
    fmt.Printf("Cache contents: %v\n", cache)

    // 7.2 COUNTING PATTERN
    fmt.Println("\n7.2 Counting Pattern:")
    
    text := "hello world hello universe hello go"
    words := strings.Fields(text)
    
    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }
    
    fmt.Printf("Word frequencies: %v\n", wordCount)

    // 7.3 GROUPING PATTERN
    fmt.Println("\n7.3 Grouping Pattern:")
    
    employees := []struct {
        Name       string
        Department string
        Salary     int
    }{
        {"Alice", "Engineering", 90000},
        {"Bob", "Marketing", 65000},
        {"Charlie", "Engineering", 85000},
        {"Diana", "Marketing", 70000},
        {"Eve", "Engineering", 95000},
    }
    
    byDepartment := make(map[string][]string)
    salaryByDept := make(map[string][]int)
    
    for _, emp := range employees {
        byDepartment[emp.Department] = append(byDepartment[emp.Department], emp.Name)
        salaryByDept[emp.Department] = append(salaryByDept[emp.Department], emp.Salary)
    }
    
    fmt.Println("Employees by department:")
    for dept, names := range byDepartment {
        fmt.Printf("%s: %v\n", dept, names)
    }

    // 7.4 CONFIGURATION PATTERN
    fmt.Println("\n7.4 Configuration Pattern:")
    
    config := map[string]interface{}{
        "database_url": "localhost:5432",
        "max_connections": 100,
        "enable_logging": true,
        "timeout_seconds": 30.5,
    }
    
    // Helper function to get config values
    getConfig := func(key string, defaultValue interface{}) interface{} {
        if value, exists := config[key]; exists {
            return value
        }
        return defaultValue
    }
    
    dbURL := getConfig("database_url", "localhost:3306").(string)
    maxConn := getConfig("max_connections", 50).(int)
    timeout := getConfig("timeout_seconds", 60.0).(float64)
    
    fmt.Printf("DB URL: %s\n", dbURL)
    fmt.Printf("Max connections: %d\n", maxConn)
    fmt.Printf("Timeout: %.1f seconds\n", timeout)

    // ================================================================
    // PART 8: MAP FUNCTIONS AND BEST PRACTICES
    // ================================================================
    fmt.Println("\n\n🏆 PART 8: FUNCTIONS AND BEST PRACTICES")
    

    // 8.1 PASSING MAPS TO FUNCTIONS
    fmt.Println("\n8.1 Passing Maps to Functions:")
    
    testMap := map[string]int{"a": 1, "b": 2}
    fmt.Printf("Before function: %v\n", testMap)
    
  
    fmt.Printf("After modifyMap: %v\n", testMap) // Modified!
    
 

    // 8.2 BEST PRACTICES
    fmt.Println("\n8.2 Best Practices:")
    fmt.Println("✅ Always check if key exists with comma ok idiom")
}