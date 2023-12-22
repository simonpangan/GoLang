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

// ----- MAPS -----

func maps() {
	// Maps are collections of key/value pairs
	// Keys can be any data type that can be compared
	// using == (They can be a different type than
	// the value)
	// Format: var myMap map[keyType]valueType

	// Declare a map variable
	var heroes map[string]string

	// Create the map
	heroes = make(map[string]string)

	// You can do it in one step
	villians := make(map[string]string)

	// Add keys and values
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"

	// Define with map literal
	superPets := map[int]string{1: "Krypto",
		2: "Bat Hound"}

	// Get value with key (Use %v with Printf)
	fmt.Printf("Batman is %v\n", heroes["Batman"])

	// If you access a key that doesn't exist
	// you get nil
	pl("Chip :", superPets[3])

	// You can check if there is a value or nil
	_, ok := superPets[3]
	pl("Is there a 3rd pet :", ok)

	// Cycle through map
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// Delete a key value
	delete(heroes, "The Flash")
}

// ----- GENERICS -----

func generics() {
	// We can specify the data type to be used at a
	// later time with generics
	// It is mainly used when we want to create
	// functions that can work with
	// multiple data types
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))
}

// ----- FUNCTION THAT EXCEPTS GENERICS -----
// This generic type parameter is capital, between
// square brackets and has a rule for what data
// it will except called a constraint
// any : anything
// comparable : Anything that supports ==
// More Constraints : pkg.go.dev/golang.org/x/exp/constraints

// You can also define what is excepted like this
// Define that my generic must be an int or float64
type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

// ----- STRUCTS -----

func structss() {
	// Structs allow you to store values with many
	// different data types structure

	// Add values
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 Main St"
	tS.bal = 234.56

	// Pass to function as values
	getCustInfo(tS)
	// or as reference
	newCustAdd(&tS, "123 South st")
	pl("Address :", tS.address)

	// Create a struct literal
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name :", sS.name)

	// Structs with functions
	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area :", rect1.Area())

	// Go doesn't support inheritance, but it does
	// support composition by embedding a struct
	// in another
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}

	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}

	bus1.info()
}

type customer struct {
	name    string
	address string
	bal     float64
}

// Customer passed as values
func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

// This struct has a function associated
type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

// Struct composition : Putting a struct in another
type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}

// ----- DEFINED TYPES -----

func definedTypes() {
	// We used a defined type previously with structs
	// You can use them also to enhance the quality
	// of other data types
	// We'll create them for different measurements

	// Convert from tsp to mL
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f mL\n", ml1)

	// Convert from TBs to mL
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f mL\n", ml2)

	// You can use arithmetic and comparison
	// operators
	pl("2 tsp + 4 tsp =", Tsp(2), Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))

	// We can convert with functions
	// Bad Way
	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))
	fmt.Printf("3 TBs = %.2f mL\n", TBToML(3))

	// We can solve this by using methods which
	// are functions associated with a type
	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())
}

// I'll define different cooking measurement types
// so we can do conversions
type Tsp float64
type TBs float64
type ML float64

// Convert with functions (Bad Way)
func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

// Associate method with types
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

// ----- PROTECTING DATA -----

func main() {
	// We want to protect our data from receiving
	// bad values by moving our date struct
	// to another package using encapsulation
	// We'll use mypackage like before

	//EX: go to mypackage/main.go line 22
}
