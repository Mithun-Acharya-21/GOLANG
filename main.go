package main

import (
	"fmt"
)

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║             GO LEARNING PROJECT - ALL EXAMPLES               ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")

	// 1. Variables Demo
	fmt.Println("\n1. VARIABLES DEMONSTRATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	variable()

	// 2. String Conversion Demo
	fmt.Println("\n2. STRING CONVERSION DEMONSTRATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	convertString()

	// 3. Pointers Demo
	fmt.Println("\n3. POINTERS DEMONSTRATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	demonstratePointers()

	// 4. Pipelines Demo
	fmt.Println("\n4. PIPELINES DEMONSTRATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	demonstratePipelines()

	// 5. Select Demo
	fmt.Println("\n5. SELECT STATEMENT DEMONSTRATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	demonstrateSelect()

	// 6. Context Demo
	fmt.Println("\n6. CONTEXT DEMONSTRATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	demonstrateContext()

	// 7. Payment Processing Demo
	fmt.Println("\n7. PAYMENT PROCESSING (RECOVERY FROM PANIC) DEMONSTRATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	processPayment("user123", 100.50)

	// 8. Channels Comprehensive Demo
	demonstrateChannels()

	fmt.Println("\n╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                  ALL EXAMPLES COMPLETED!                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝\n")
}


