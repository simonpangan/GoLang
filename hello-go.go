package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var pl = fmt.Println

// Basic Print + User Input

func basics() {
	pl("What is your name?")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl("Hello, " + name)
	} else {
		log.Fatal(err)
	}
}

// Variables

func variables() {
	// Format: var name type
	// must be small case for name

	//Variables are mutable by default

	//type declaration
	var vName string = "John"

	//dynamic declaration
	var v1, v2 = 1.2, 2.4
	var v3 = "Hello"

	v4 := "World"

	pl(vName, v1, v2, v3, v4)
}

//Data types

func dataTypes() {
	// int, float64, bool, string, rune
	// Default type: 0, 0.0, false, "",

	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf("😌"))
}

// Casting

func casting() {
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
	pl(cv9)
}

//Conditional

func conditional() {
	iAge := 51

	if (iAge >= 1) && (iAge <= 18) {
		pl("Important birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important birthday")
	} else if iAge >= 65 {
		pl("Important birthday")
	} else {
		pl("Not an Important birthday")
	}

	pl("!true =", !true)
}

//Strings

func str() {
	//Strings are array of bytes
	// []bytes
	sv1 := "A word"

	replacer := strings.NewReplacer("A", "Another")
	sv2 := replacer.Replace(sv1)

	pl(sv2)
	pl("Length: ", len(sv2))
	pl("Contains \"Another\": ", strings.Contains(sv2, "Another"))

	pl("o index :", strings.Index(sv2, "o"))
	pl("Replace :", strings.Replace(sv2, "o", "0", -1))

	sv3 := "\n Some Words \n"

	pl(sv3)
	sv3 = strings.TrimSpace(sv3)
	pl(sv3)
	pl("Split:", strings.Split("a-b-c-d", "-"))
	pl("Lower :", strings.ToLower((sv2)))
	pl("Upper :", strings.ToUpper((sv2)))

	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))

}

//Runes

func run() {
	//In go characters are called runes
	rStr := "abcdefg"

	pl("Rune Count:", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}

//Time

func oras() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day(), now.Weekday())
	pl(now.Hour(), now.Minute(), now.Second())
}

// Math

func mathematics() {
	pl("5 + 4 =", 5+4)
	pl("5 - 4 =", 5-4)
	pl("5 * 4 =", 5*4)
	pl("5 / 4 =", 5/4)
	pl("5 % 4 =", 5%4)

	// Shorthand increment
	// Instead of mInt = mInt + 1 (mInt += 1)
	// -= *= /=
	mInt := 1
	mInt += 1
	// Also increment or decrement with ++ and --
	mInt++

	// Float precision increases with the size of your values
	pl("Float Precision =", 0.11111111111111111111+
		0.11111111111111111111)

	// Create a random value between 0 and 50
	// Get a seed value for our random number generator based on
	// seconds since 1/1/70 to make our random number more random
	seedSecs := time.Now().Unix() // Returns seconds as int
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)

	// There are many math functions
	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert 1.5708 radians to degrees
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))

	// There are also functions for Cos, Tan, Acos, Asin
	// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos
	// Htpot
}

// ----- FORMATTED PRINT -----

func formatted() {
	// Go has its own version of C's printf
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value

	fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A',
		3.14, true, 1, 1)

	// Float formatting
	fmt.Printf("%9f\n", 3.14)      // Width 9
	fmt.Printf("%.2f\n", 3.141592) // Decimal precision 2
	fmt.Printf("%9.f\n", 3.141592) // Width 9 no precision

	// Sprintf returns a formatted string instead of printing
	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
}

//Loop

func loop() {
	// for initialization; condition; postStatement {BODY}
	// Print numbers 1 through 5
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	// Do the opposite
	for x := 5; x >= 1; x-- {
		pl(x)
	}

	// x is out of the scope of the for loop so it doesn't exist
	// pl("x :", x)

	// For is used to create while loops as well
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

}

//While true loop

func while() {
	seedSecs := time.Now().Unix() // Returns seconds as int
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 0 and 50 : ")
		pl("Random Number is :", randNum)

		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)

		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randNum {
			pl("Lower")
		} else if iGuess < randNum {
			pl("Higher")
		} else {
			pl("You Guessed it")
			break
		}
	}
}

//array loop

func arrayLoop() {
	// Cycle through an array with range
	// We don't need the index so we ignore it
	// with the blank identifier

	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}

}

// ----- ARRAYS -----

func arrays() {
	// Collection of values with the same data type
	// and the size can't be changed
	// Default values are 0, 0.0, false or ""

	// Declare integer array with 5 elements
	var arr1 [5]int

	// Assign value to index
	arr1[0] = 1

	// Declare and initialize
	arr2 := [5]int{1, 2, 3, 4, 5}

	// Get by index
	pl("Index 0 :", arr2[0])

	// Length
	pl("Arr Length :", len(arr2))

	// Iterate with index
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}

	// Iterate with range
	for i, v := range arr2 {
		fmt.Printf("%d : %d", i, v)
	}

	// Multidimensional Array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	// Print multidimensional array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	// String into slice of runes
	aStr1 := "abcde"
	rArr := []rune(aStr1)

	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

	// Byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)
}

// ----- SLICES -----

func slicess() {
	// Slices are like arrays but they can grow
	// Format: var name []dataType

	// Create a slice with make
	sl1 := make([]string, 6)

	// Assign values by index
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	// Size of slice
	pl("Slice Size :", len(sl1))

	// Cycle with for
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	// Cycle with range
	for _, x := range sl1 {
		pl(x)
	}

	// Create a slice literal
	sl2 := []int{12, 21, 1974}
	pl(sl2)

	// A slice points at an array and you can create a slice
	// of an array (A slice is a view of an underlying array)
	// You can have multiple slices point to the same array
	sArr := [5]int{1, 2, 3, 4, 5}
	
	// Start at 0 index up to but not including the 2nd index
	sl3 := sArr[0:2]
	pl(sl3)

	// Get slice from beginning
	pl("1st 3 :", sArr[:3])

	// Get slice to the end
	pl("Last 3 :", sArr[2:])

	// If you change the array the slice also changes
	sArr[0] = 10
	pl("sl3 :", sl3)

	// Changing the slice also changes the array
	sl3[0] = 1
	pl("sArr :", sArr)

	// Append a value to a slice (Also overwrites array)
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	// Printing empty slices will return nils which show
	// as empty slices
	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])
}
