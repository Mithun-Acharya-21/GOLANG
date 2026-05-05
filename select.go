select {
case msg := <-ch:
    fmt.Println("received", msg)
case <-ctx.Done():
    fmt.Println("cancelled")
}
