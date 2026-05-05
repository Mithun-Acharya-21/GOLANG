package main

import (
	"context"
	"fmt"
	"time"
)

func demonstrateSelect() {
	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "received message"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Channel:", msg)
	case <-ctx.Done():
		fmt.Println("Timeout: operation cancelled")
	}
}
