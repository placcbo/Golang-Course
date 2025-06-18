package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("=== Complete Go Goroutines Tutorial ===")
	fmt.Println("From Basics to Advanced Patterns")

	// ================================================================
	// PART 1: INTRODUCTION TO GOROUTINES
	// ================================================================
	fmt.Println("🟢 PART 1: INTRODUCTION TO GOROUTINES")
	
	fmt.Println("\n1.1 What are Goroutines?")
	fmt.Println("• Lightweight threads managed by the Go runtime")
	fmt.Println("• Much cheaper than OS threads (can have millions)")
	fmt.Println("• Communicate via channels (share memory by communicating)")
	fmt.Println("• Created with the 'go' keyword before a function call")

	// ================================================================
	// PART 2: BASIC GOROUTINE USAGE
	// ================================================================

	fmt.Println("\n\n🔵 PART 2: BASIC GOROUTINE USAGE")

	// 2.1 LAUNCHING A SIMPLE GOROUTINE
	fmt.Println("\n2.1 Launching a Simple Goroutine")

	// A simple function we'll run concurrently

	printNumbers := func(){
		for i := 0; i <= 5; i++{
			fmt.Printf("%d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}

	
	// Launch goroutine (will run concurrently with main)
	go printNumbers()

		// Main function continues immediately
	fmt.Println("Main continues execution while goroutine runs...")

	// Wait to see goroutine output (naive approach - better methods later)
	time.Sleep(1 * time.Second)

	fmt.Println("\n\n(Note: We used time.Sleep here just for demonstration)")

	// 2.2 LAUNCHING MULTIPLE GOROUTINES
	fmt.Println("\n2.2 Launching Multiple Goroutines")




































	
}