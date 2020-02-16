package main

import (
	"fmt"
	"time"
)

var deposits = make(chan int)
var balances = make(chan int)

func deposit(amount int) { deposits <- amount }
func balance() int {
	i, _ := <-balances
	//fmt.Println(ok)
	return i
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			fmt.Println("hello")
		}
	}
}

func main() {
	go teller()
	deposits <- 100
	for i := 0; i < 10; i++ {
		if i == 5 {
			close(balances)
		}
		if i >= 5 {
			continue
		}
		fmt.Println(balance())
	}

	time.Sleep(5 * time.Second)
}
