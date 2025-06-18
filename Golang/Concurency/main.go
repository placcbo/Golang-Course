package main

import (
	"fmt"
	"time"
	"sync"
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

	// Another function for demonstration

	printLetters := func (){
		for c := 'a'; c <= 'e'; c++{
			fmt.Printf("%c", c)
			time.Sleep(150 * time.Millisecond)
		}
	}

	// Launch multiple goroutines

	go printLetters()
	go printLetters()


	fmt.Println("Main continues... (multiple goroutines running)")
	time.Sleep(2 * time.Second)

	fmt.Println("\nObserve how numbers and letters are interleaved")


	// ================================================================
	// PART 3: SYNCHRONIZATION WITH WAITGROUPS
	// ================================================================

	fmt.Println("\n\n🟡 PART 3: SYNCHRONIZATION WITH WAITGROUPS")

	// 3.1 USING SYNC.WAITGROUP
	fmt.Println("\n3.1 Using sync.WaitGroup for Proper Synchronization")

	// Better alternative to time.Sleep
	var wg sync.WaitGroup

	// Enhanced printNumbers function

	printNumbersWithWG := func(){
		defer wg.Done() // Decrement counter when done
		for i := 1; i <= 5; i++ {
			fmt.Printf("%d ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}

	// Enhanced printLetters function
	printLettersWithWG := func() {
		defer wg.Done()
		for c := 'a'; c <= 'e'; c++ {
			fmt.Printf("%c ", c)
			time.Sleep(150 * time.Millisecond)
		}
	}


		// Launch goroutines with proper synchronization

		wg.Add(2) // Set counter to 2 (for 2 goroutines)
	go printNumbersWithWG()
	go printLettersWithWG()

	fmt.Println("Main waiting for goroutines to finish...")
	wg.Wait() // Block until counter is 0
	fmt.Println("\nAll goroutines completed!")


	// 3.2 WAITGROUP WITH ANONYMOUS FUNCTIONS




































	
}