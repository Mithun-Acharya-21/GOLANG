package main

import "fmt"
type Account struct {
    ID      string
    Balance float64
}

func updateBalance(acc *Account, amount float64) {
    acc.Balance += amount
}

func demonstratePointers() {
    user := Account{ID: "123", Balance: 1000}
    updateBalance(&user, -200) // pass pointer
    fmt.Println(user.Balance)  // prints 800
}
