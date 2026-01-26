package main

import (
    "fmt"
    "log"
)


func processPayment(userID string, amount float64) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Recovered from panic: %v", r)
        }
        log.Println("Finished processing payment for user:", userID)
    }()
    if amount <= 0 {
        panic(fmt.Sprintf("Invalid payment amount: %.2f", amount))
    }
    log.Println("Debiting account for user:", userID, "Amount:", amount)
}
