package main
import (
	"fmt"
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
}