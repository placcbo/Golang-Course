package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	fmt.Println("=== Complete Go Goroutines Tutorial ===")
	fmt.Println("From Basics to Advanced Patterns\n")

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
	printNumbers := func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("%d ", i)
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
	printLetters := func() {
		for c := 'a'; c <= 'e'; c++ {
			fmt.Printf("%c ", c)
			time.Sleep(150 * time.Millisecond)
		}
	}

	// Launch multiple goroutines
	go printNumbers()
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
	printNumbersWithWG := func() {
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
	fmt.Println("\n3.2 WaitGroup with Anonymous Functions")

	wg.Add(3) // We'll launch 3 goroutines

	// Goroutine 1
	go func() {
		defer wg.Done()
		fmt.Println("Anonymous goroutine 1 starting")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Anonymous goroutine 1 done")
	}()

	// Goroutine 2
	go func(id int) {
		defer wg.Done()
		fmt.Printf("Anonymous goroutine %d starting\n", id)
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("Anonymous goroutine %d done\n", id)
	}(2)

	// Goroutine 3
	go func() {
		defer wg.Done()
		fmt.Println("Anonymous goroutine 3 starting")
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Anonymous goroutine 3 done")
	}()

	fmt.Println("Main waiting for anonymous goroutines...")
	wg.Wait()
	fmt.Println("All anonymous goroutines completed")

	// ================================================================
	// PART 4: GOROUTINES WITH PARAMETERS
	// ================================================================
	fmt.Println("\n\n🟠 PART 4: GOROUTINES WITH PARAMETERS")

	// 4.1 PASSING PARAMETERS TO GOROUTINES
	fmt.Println("\n4.1 Passing Parameters to Goroutines")

	// Function that takes parameters
	worker := func(id int, delay time.Duration) {
		fmt.Printf("Worker %d starting (delay: %v)\n", id, delay)
		time.Sleep(delay)
		fmt.Printf("Worker %d done\n", id)
	}

	// Launch goroutines with different parameters
	wg.Add(3)
	go func() {
		defer wg.Done()
		worker(1, 800*time.Millisecond)
	}()
	go func() {
		defer wg.Done()
		worker(2, 300*time.Millisecond)
	}()
	go func() {
		defer wg.Done()
		worker(3, 500*time.Millisecond)
	}()

	fmt.Println("Main waiting for workers...")
	wg.Wait()
	fmt.Println("All workers completed")

	// 4.2 CLOSURE CAPTURE GOTCHA
	fmt.Println("\n4.2 Closure Capture Gotcha")

	fmt.Println("Incorrect way (loop variable capture):")
	for i := 0; i < 3; i++ {
		go func() {
			// All goroutines will likely see i=3 due to closure capture
			fmt.Printf("Incorrect: %d ", i)
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\n(Notice all goroutines saw the same value)")

	fmt.Println("\nCorrect way (pass as parameter):")
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Correct: %d ", id)
		}(i) // Pass current loop value as parameter
	}
	time.Sleep(1 * time.Second)
	fmt.Println("\n(Now each goroutine gets its own value)")

	// ================================================================
	// PART 5: ADVANCED GOROUTINE PATTERNS
	// ================================================================
	fmt.Println("\n\n🟣 PART 5: ADVANCED GOROUTINE PATTERNS")

	// 5.1 FAN-OUT/FAN-IN PATTERN
	fmt.Println("\n5.1 Fan-out/Fan-in Pattern")

	// Worker function
	workerFunc := func(id int, jobs <-chan int, results chan<- int) {
		for j := range jobs {
			fmt.Printf("Worker %d processing job %d\n", id, j)
			time.Sleep(time.Duration(j) * time.Millisecond * 100)
			results <- j * 2
		}
	}

	// Create channels
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Launch 3 workers (fan-out)
	for w := 1; w <= 3; w++ {
		go workerFunc(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results (fan-in)
	go func() {
		for r := 1; r <= 5; r++ {
			fmt.Printf("Result: %d\n", <-results)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Fan-out/fan-in pattern completed")

	// 5.2 WORKER POOL PATTERN
	fmt.Println("\n5.2 Worker Pool Pattern")

	type Task struct {
		ID    int
		Value int
	}

	// Worker pool function
	workerPool := func(id int, tasks <-chan Task, results chan<- int) {
		for t := range tasks {
			fmt.Printf("Worker %d processing task %d\n", id, t.ID)
			time.Sleep(time.Duration(t.Value) * time.Millisecond * 50)
			results <- t.Value * 3
		}
	}

	// Create task and result channels
	taskChan := make(chan Task, 10)
	resultChan := make(chan int, 10)

	// Create worker pool
	for w := 1; w <= 4; w++ {
		go workerPool(w, taskChan, resultChan)
	}

	// Submit tasks
	for i := 1; i <= 8; i++ {
		taskChan <- Task{ID: i, Value: i}
	}
	close(taskChan)

	// Collect results
	go func() {
		for i := 1; i <= 8; i++ {
			fmt.Printf("Worker pool result: %d\n", <-resultChan)
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Worker pool pattern completed")

	// ================================================================
	// PART 6: ERROR HANDLING IN GOROUTINES
	// ================================================================
	fmt.Println("\n\n🔴 PART 6: ERROR HANDLING IN GOROUTINES")

	// 6.1 ERROR RETURNING GOROUTINES
	fmt.Println("\n6.1 Error Returning Goroutines")

	// Function that might fail
	errorProneFunc := func(id int) error {
		if id%2 == 0 {
			return fmt.Errorf("error in task %d", id)
		}
		fmt.Printf("Task %d completed successfully\n", id)
		return nil
	}

	// Channel for errors
	errChan := make(chan error, 5)

	// Launch goroutines that might fail
	for i := 1; i <= 5; i++ {
		go func(id int) {
			err := errorProneFunc(id)
			if err != nil {
				errChan <- err
			}
		}(i)
	}

	// Collect errors
	go func() {
		for i := 1; i <= 5; i++ {
			select {
			case err := <-errChan:
				fmt.Printf("Received error: %v\n", err)
			case <-time.After(1 * time.Second):
				fmt.Println("Timeout waiting for errors")
				return
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Error handling demonstration completed")

	// ================================================================
	// PART 7: BEST PRACTICES AND GOTCHAS
	// ================================================================
	fmt.Println("\n\n🏆 PART 7: BEST PRACTICES AND GOTCHAS")

	fmt.Println("\n7.1 Best Practices:")
	fmt.Println("✅ Use WaitGroups for synchronization when needed")
	fmt.Println("✅ Always check for errors in goroutines")
	fmt.Println("✅ Use context.Context for cancellation and timeouts")
	fmt.Println("✅ Keep goroutines focused on a single task")
	fmt.Println("✅ Use buffered channels when producer/consumer rates differ")
	fmt.Println("✅ Document goroutine ownership (who starts/stops them)")

	fmt.Println("\n7.2 Common Gotchas:")
	fmt.Println("⚠️  Goroutines leak if they run indefinitely")
	fmt.Println("⚠️  Loop variable capture in closures (as shown earlier)")
	fmt.Println("⚠️  Unbuffered channels can cause deadlocks")
	fmt.Println("⚠️  Goroutines may not run if main exits first")
	fmt.Println("⚠️  Shared memory access without synchronization (use channels)")
	fmt.Println("⚠️  Blocking operations can starve the goroutine scheduler")

	fmt.Println("\n7.3 When to Use Goroutines:")
	fmt.Println("• I/O-bound operations (network, disk)")
	fmt.Println("• CPU-bound operations that can be parallelized")
	fmt.Println("• Background tasks (logging, metrics)")
	fmt.Println("• Event processing pipelines")
	fmt.Println("• Any task that can benefit from concurrency")

	fmt.Println("\n=== End of Go Goroutines Tutorial ===")
}