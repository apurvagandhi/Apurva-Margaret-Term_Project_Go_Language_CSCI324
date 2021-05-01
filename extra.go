/******************************************
Names: Apurva Gandhi and Margaret Haley
Course: CSCI324
Professor King
Extra Program to illustrate Go's ability to
add methods to primitive types
Sample execution: go run extra.go
*****************************************/

package main

import (
	"fmt"
)

// Declares new Int type
type Int int

// Adds method to that type
func (i Int) specialComputation(j Int) Int {
	return i*j + 10
}

func main() {
	var firstNumber int
	var secondNumber int

	// Ask user for numbers
	fmt.Println("What numbers would you like us to perform our special computation on?")
	fmt.Println("Please enter the numbers separated by a space (Ex. 2 5): ")
	fmt.Scanf("%d %d", &firstNumber, &secondNumber)

	// Convert to Int type
	i := Int(firstNumber)
	j := Int(secondNumber)

	// Print result of the special computation on the user's numbers
	fmt.Print("The result of our computation is: ")
	fmt.Println(i.specialComputation(j))
	fmt.Println("We multiplied your two numbers together, then added 10.")
}
