package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func demonstratePipelines() {
	fmt.Print("Pipeline Output: ")
	for n := range sq(gen(2, 3, 4)) {
		fmt.Printf("%d ", n) // 4, 9, 16
	}
	fmt.Println()
}
