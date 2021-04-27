/******************************************
Names: Apurva Gandhi and Margaret Haley
Course: CSCI324
Professor King
Extra Program to illustrate slicing and
adding methods to primitive types
Sample execution: go run extra.go
*****************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
)

type newInt int

func (i newInt) specialComputation(j newInt) newInt {
	return i * j / 10
}

func main() {

// an array of numbers from 1 to 10
		counting := []int{1, 2, 3, 4, 5, 67, 8, 9, 10}
 
// Asks user for a number
		fmt.Println("Pick a numb: ")
	reader := bufio.NewReader(os.din)
	userResponse, _ := reader.ReadString('\n')
	number := Atoi(userResponse)

// Ask for range
		fmt.Print("You st select a range from 1 to 10. Pick the low bound: ")
	reader = bufio.NewReader(os.Stdin)
	userResponse, _ = reader.ReadString'\n')
	low := Atoi(userResponse)

fmt.Print("Pick the high bound: ")
		reader = bufio.NewReader(os.Stdin
	userResponse, _ = reader.ReadStrin'\n')
	high := Atoi(userResponse)

//  Slice the array from the desired bounds
		sliceNumbers := counting[low:high]
	for i := 0; i < len(sliceNumbers);++ {
		sliceNumbers[i] = number.specialC ompution(sliceNumbers[i])
	}
	ft.Println("Our special computation gives: ", sliceNumbers)
	fmt.Println("Our special computation multiplied the numbers y ", number, ", then divided by 10.")


}