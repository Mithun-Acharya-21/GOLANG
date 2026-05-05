package main

import (
	"fmt"
	"time"
)

// ============== 1. Basic Channel Example ==============
func basicChannel() {
	fmt.Println("\n--- BASIC CHANNEL EXAMPLE ---")
	ch := make(chan string)

	// Launch goroutine
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "Hello from goroutine!"
	}()

	// Receive from channel
	message := <-ch
	fmt.Println("Received:", message)
}

// ============== 2. Buffered Channel Example ==============
func bufferedChannel() {
	fmt.Println("\n--- BUFFERED CHANNEL EXAMPLE ---")
	// Create buffered channel with capacity 2
	ch := make(chan int, 2)

	// Send values without blocking
	ch <- 1
	ch <- 2
	fmt.Println("Sent two values to buffered channel")

	// Receive values
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}

// ============== 3. Channel with Goroutines ==============
func channelWithGoroutines() {
	fmt.Println("\n--- CHANNEL WITH GOROUTINES ---")
	numbers := make(chan int)
	squares := make(chan int)

	// Goroutine 1: Send numbers
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(numbers)
	}()

	// Goroutine 2: Square the numbers
	go func() {
		for num := range numbers {
			squares <- num * num
		}
		close(squares)
	}()

	// Main goroutine: Receive and print
	for square := range squares {
		fmt.Println("Square:", square)
	}
}

// ============== 4. Worker Pool Pattern ==============
func workerPool() {
	fmt.Println("\n--- WORKER POOL PATTERN ---")
	jobs := make(chan int, 5)
	results := make(chan string, 5)

	// Create 3 workers
	for w := 1; w <= 3; w++ {
		go func(workerID int) {
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
				time.Sleep(100 * time.Millisecond)
				results <- fmt.Sprintf("Job %d completed by Worker %d", job, workerID)
			}
		}(w)
	}

	// Send jobs
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Receive results
	for i := 0; i < 5; i++ {
		fmt.Println("Result:", <-results)
	}
}

// ============== 5. Ping-Pong Example ==============
func pingPong() {
	fmt.Println("\n--- PING-PONG EXAMPLE ---")
	ping := make(chan string)
	pong := make(chan string)

	go func() {
		ping <- "ping"
	}()

	go func() {
		msg := <-ping
		fmt.Println("Received:", msg)
		pong <- "pong"
	}()

	msg := <-pong
	fmt.Println("Received:", msg)
}

// ============== 6. Fan-Out, Fan-In Pattern ==============
func fanOutFanIn() {
	fmt.Println("\n--- FAN-OUT, FAN-IN PATTERN ---")

	// Single input channel
	input := make(chan int)

	// Fan-out: Create multiple workers
	ch1 := worker1(input)
	ch2 := worker2(input)

	// Send data
	go func() {
		for i := 1; i <= 3; i++ {
			input <- i
		}
		close(input)
	}()

	// Fan-in: Combine results
	for i := 0; i < 6; i++ {
		select {
		case v := <-ch1:
			fmt.Println("Worker 1 result:", v)
		case v := <-ch2:
			fmt.Println("Worker 2 result:", v)
		}
	}
}

func worker1(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for n := range input {
			output <- n * 2 // Double the number
		}
		close(output)
	}()
	return output
}

func worker2(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for n := range input {
			output <- n * 3 // Triple the number
		}
		close(output)
	}()
	return output
}

// ============== 7. Timeout Pattern ==============
func timeoutPattern() {
	fmt.Println("\n--- TIMEOUT PATTERN ---")
	ch := make(chan string)

	// Goroutine that takes time
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "completed"
	}()

	// Wait with timeout
	select {
	case result := <-ch:
		fmt.Println("Result:", result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: operation took too long")
	}
}

// ============== 8. Multiple Channel Select ==============
func multipleChannelSelect() {
	fmt.Println("\n--- MULTIPLE CHANNEL SELECT ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "Channel 1 message"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch2 <- "Channel 2 message"
	}()

	// Wait for both
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}

// ============== MAIN DEMONSTRATION FUNCTION ==============
func demonstrateChannels() {
	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║         GO CHANNELS - COMPREHENSIVE EXAMPLES            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")

	basicChannel()
	bufferedChannel()
	channelWithGoroutines()
	workerPool()
	pingPong()
	fanOutFanIn()
	timeoutPattern()
	multipleChannelSelect()

	fmt.Println("\n╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║              ALL CHANNEL EXAMPLES COMPLETED             ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝\n")
}
