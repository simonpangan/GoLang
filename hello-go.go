package main

import (
	"bufio"
	"fmt"
	"log"
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

func rune() {
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
