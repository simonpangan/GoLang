package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

// ----- CONCURRENCY -----
// Concurrency allows us to have multiple
// blocks of code share execution time by
// pausing their execution. We can also
// run blocks of codes in parallel at the same
// time. (Concurrent tasks in Go are called
// goroutines)

// To execute multiple functions in new
// goroutines add the word go in front of
// the function calls (Those functions can't
// have return values)

// We can't control when functions execute
// so we may get different results

func goroutines() {
	go printTo10()
	go printTo15()

	// We have to pause the main function because
	// if main ends so will the goroutines
	time.Sleep(2 * time.Second) // Pause 2 seconds
}

func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Func 1 :", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Func 2 :", i)
	}
}

// ----- CHANNELS -----
// https://www.youtube.com/watch?v=B9uR2gLM80E

func channels() {
	// You can have goroutines communicate by
	// using channels. The sending goroutine
	// also makes sure the receiving goroutine
	// receives the value before it attempts
	// to use it

	// Create a channel : Only carries values of
	// 1 type
	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)

	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)
}

// These functions will print in order using
// channels
// Func receives a channel and then sends values
// over channels once each time it is called
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

// Using locks to protect data from being
// accessed by more than one user at a time
// Locks are another option when you don't
// have to pass data

func locks() {
	var acct Account
	acct.balance = 100
	pl("Balance :", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}

// ----- BANK ACCOUNT EXAMPLE -----
// Here I'll simulate customers accessing a
// bank account and lock out customers to
// allow for individual access
type Account struct {
	balance int
	lock    sync.Mutex // Mutual exclusion which means can only be access at one time
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock() // unlock after the function is done
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n",
			v, a.balance)
		a.balance -= v
	}
}

// ----- CLOSURES -----

func closures() {
	// Closures are functions that don't have to be
	// associated with an identifier (Anonymous)

	// Create a closure that sums values
	intSum := func(x, y int) int {
		return x + y
	}

	pl("5 + 4 =", intSum(5, 4))

	// Closures can change values outside the function
	samp1 := 1
	changeVar := func() { samp1 += 1 }
	changeVar()
	pl("samp1 =", samp1)

	// Pass a function to a function
	useFunc(sumVals, 5, 8)
}

// ----- CLOSURES -----
// Pass a function to a function
func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

// ----- RECURSION -----

func recursion() {
// Recursion occurs when a function calls itself
	// There must be a condition that ends this
	// Finding a factorial is commonly used
	pl("Factorial 4 =", factorial(4))
	// 1st : result = 4 * factorial(3) = 4 * 6 = 24
	// 2nd : result = 3 * factorial(2) = 3 * 2 = 6
	// 3rd : result = 2 * factorial(1) = 2 * 1 = 2
}

func factorial(num int) int {
	// This condition ends calling functions
	if num == 0 {
		return 1
	}
	
	return num * factorial(num-1)
}