package main

// The path to your project is declared in your
// go.mod file followed by the directory
import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}
