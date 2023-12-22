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
