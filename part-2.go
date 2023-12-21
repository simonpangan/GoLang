package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

// ----- FILE IO -----

func fileIO() {
	// We can create, write and read from files

	// Create a file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Says to close the file after program ends or when
	// there is a closing curly bracket
	defer f.Close()

	// Create list of primes
	iPrimeArr := []int{2, 3, 5, 7, 11}

	// Create string array from int array
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	// Cycle through strings and write to file
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	// Open the created file
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read from file and print once per line
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime :", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	// Append to file
	/*
		 Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified

		O_RDONLY : open the file read-only
		O_WRONLY : open the file write-only
		O_RDWR   : open the file read-write

		These can be or'ed

		O_APPEND : append data to the file when writing
		O_CREATE : create a new file if none exists
		O_EXCL   : used with O_CREATE, file must not exist
		O_SYNC   : open for synchronous I/O
		O_TRUNC  : truncate regular writable file when opened
	*/

	// Check if file exists
	_, err = os.Stat("text.txt")

	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exist")
	} else {
		f, err = os.OpenFile("text.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}

}

// ----- PACKAGES -----

func packages() {
	// Packages allow you to keep related code together
	// Go looks for package code in a directory

	// If you are using VSC and have multiple
	// modules you get this error
	// gopls requires a module at the root of
	// your workspace
	// 1. Settings
	// 2. In search type gopls
	// 3. Paste "gopls": { "experimentalWorkspaceModule": true, }
	// 4. Restart VSC

	// cd /D D:\Tutorials\GoTutorial  //File path

	// Create a go directory : mkdir app
	// cd app
	// Choose a module path and create a go.mod file
	// Type: go mod init example/project

	// Go modules allow you to manage libraries
	// They contain one project or library and a
	// collection of Go packages
	// go.mod : contains the name of the module and versions
	// of other modules your module depends on

	// Create a main.go file at the same level as go.mod

	// You can have many packages and sub packages
	// create a directory called mypackage in the project
	// directory mkdir mypackage
	// cd mypackage

	// Create file mypackage.go in it

	// Package names should be all lowercase
}
