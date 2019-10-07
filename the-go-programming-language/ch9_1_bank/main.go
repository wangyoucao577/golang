package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int) //send amount to deposit
var balances = make(chan int) //receive balance
var withdrawResults = make(chan bool)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	if amount <= 0 {
		return false // invalid parameter
	}
	deposits <- (-amount)
	return <-withdrawResults
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			if amount >= 0 {
				fmt.Println("+", amount)
				balance += amount
			} else {
				fmt.Println(amount)
				amount = -amount
				if balance >= amount {
					balance -= amount
					fmt.Println("succeed")
					withdrawResults <- true
				} else {
					fmt.Println("failed")
					withdrawResults <- false
				}
			}
			fmt.Println("=", balance)
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		Deposit(200)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Withdraw(300)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Deposit(100)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Withdraw(150)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Withdraw(150)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("total: ", Balance())
}
