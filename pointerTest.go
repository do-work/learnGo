package main

import "fmt"

func main() {
	// declare myInt as an integer variable
	var myInt int
	// myIntPointer as a var that hold an integer pointer
	var myIntPointer *int
	// assign a value to myInt
	myInt = 43
	// assign a pointer to myInt as the value of myIntPointer
	myIntPointer = &myInt
	// print the value at myIntPointer
	fmt.Println(*myIntPointer)
}
