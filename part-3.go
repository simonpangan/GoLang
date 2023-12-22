package main

import (
	"fmt"
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
