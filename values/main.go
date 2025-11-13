package main

import "fmt"

func main() {
    // --- Integers ---
    fmt.Println("Integers:")
    fmt.Println("1 + 2 =", 1+2)
    fmt.Println("7 * 3 =", 7*3)
    fmt.Println("10 / 3 =", 10/3) // integer division (truncates)

    // --- Floats ---
    fmt.Println("\nFloats:")
    fmt.Println("10.0 / 3.0 =", 10.0/3.0) // float division

    // --- Booleans ---
    fmt.Println("\nBooleans:")
    fmt.Println("true && false =", true && false)
    fmt.Println("true || false =", true || false)
    fmt.Println("!true =", !true)

    // --- Strings ---
    fmt.Println("\nStrings:")
    fmt.Println("Hello" + " " + "Go!") // concatenation

    // --- Variables ---
    fmt.Println("\nVariables:")
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    e := 42 // short declaration
    fmt.Println(e)

    // --- Constants ---
    fmt.Println("\nConstants:")
    const Pi = 3.14159
    fmt.Println("Pi =", Pi)

    // --- Type inference ---
    fmt.Println("\nType inference:")
    x := 10        // int
    y := 3.5       // float64
    z := "GoLang!" // string
    fmt.Println(x, y, z)
	fmt.Println(true && true || false)
}
