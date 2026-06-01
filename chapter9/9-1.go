package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan int)
var withdrawResult = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	withdraws <- amount
	res := <-withdrawResult
	return res
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			continue
		case amount := <-withdraws:
			if balance >= amount {
				balance -= amount
				withdrawResult <- true
			} else {
				withdrawResult <- false
			}
		}
	}
}

func main() {
	go teller() // start the monitor goroutine

	var wg sync.WaitGroup

	// Alice:
	wg.Go(func() {
		Deposit(200) // A1
		Withdraw(50)
	})

	// Bob:
	wg.Go(func() {
		Deposit(100) // B
		Deposit(15)
		Withdraw(10)
	})

	wg.Wait()
	fmt.Println("=", Balance())
}
