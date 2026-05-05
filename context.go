package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopped")
			return
		case val := <-ch:
			fmt.Println("processing", val)
		}
	}
}

func demonstrateContext() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	go worker(ctx, ch)

	ch <- 42
	time.Sleep(100 * time.Millisecond)
	cancel() // stop worker cleanly
}
