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

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    ch := make(chan int)

    go worker(ctx, ch)

    ch <- 42
    cancel() // stop worker cleanly
}
