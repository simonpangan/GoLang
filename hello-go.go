package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

// Basic Print + User Input

// func main() {
// 	pl("What is your name?")

// 	reader := bufio.NewReader(os.Stdin)
// 	name, err := reader.ReadString('\n')

// 	if err == nil {
// 		pl("Hello, " + name)
// 	} else {
// 		log.Fatal(err)
// 	}
// }

// Variables

// func main() {
// 	// Format: var name type
// 	// must be small case for name

// 	//Variables are mutable by default

// 	//type declaration
// 	var vName string = "John"

// 	//dynamic declaration
// 	var v1, v2 = 1.2, 2.4
// 	var v3 = "Hello"

// 	v4 := "World"
// }

//Data types

// func main() {
// 	// int, float64, bool, string, rune
// 	// Default type: 0, 0.0, false, "",

// 	pl(reflect.TypeOf(25))
// 	pl(reflect.TypeOf(3.14))
// 	pl(reflect.TypeOf(true))
// 	pl(reflect.TypeOf("Hello"))
// 	pl(reflect.TypeOf("😌"))
// }

// Casting

func main() {
	cV1 := 3.14
	cV2 := int(cV1)

	pl(cV2)

	cv3 := "500000000"
	cv4, err := strconv.Atoi(cv3)

	pl(cv4, err, reflect.TypeOf(cv4))

	cv5 := 500000000
	cv6 := strconv.Itoa(cv5)
	pl(cv6, reflect.TypeOf(cv6))

	cv7 := "3.14"

	if cv8, err := strconv.ParseFloat(cv7, 64); err == nil {
		pl(cv8, reflect.TypeOf(cv8))
	}

	cv9 := fmt.Sprintf("%f", 3.14)
}
