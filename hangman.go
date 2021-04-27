/******************************************
Names: Apurva Gandhi and Margaret Haley
Date: 03.15.2021
Course: CSCI324
Professor King
Go Program for Hangman Game
Sample execution: go run hangman.go
*****************************************/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var guessed = make(map[int]bool)
var wrong int = 0
var numGuessed int = 0
var dict string = "dictionary.txt"
var mapping = make(map[string]bool)

/* The following structs and functions create and modify our linked list structure,
   which we use to keep track of what letters the user has guessed in alphabetical
   order. */

type node struct {
	value string
	next  *node
}

type linkedList struct {
	head *node
}

// printLinkedList prints out the list containing every letter the user has guessed
func (list linkedList) printLinkedList() {
	for list.head != nil {
		fmt.Print(list.head.value)
		fmt.Print(" ")
		list.head = list.head.next
	}
	fmt.Println()
}

// insertLetter inserts the user's newest guessed into our linked list in sorted order
func (list *linkedList) insertLetter(letter string) {
	new := &node{value: letter}
	current := list.head
	if list.head == nil {
		list.head = new
	} else if letter < list.head.value {
		new.next = list.head
		list.head = new
	} else {
		for current.next != nil && letter > current.next.value {
			current = current.next
		}
		new.next = current.next
		current.next = new
	}
}

/* checkLetter first determines if this letter has been guessed already. If not,
   it calls insertLetter to place it in the lists and then checks to see if this
   new guess is a letter (or letters) in our word. Then, it updates the necessary
   variables accordingly and prints how many letters the user has correctly
   guessed (which are revealed) and how many incorrect guesses he/she has made. */
func (list *linkedList) checkLetter(letter string, word string) {
	if !(list.duplicateLetter(letter)) {
		fmt.Println("You have already guessed this letter, try a new one.")
		fmt.Println()
	} else {
		list.insertLetter(letter)
		inWord := false
		for i := 0; i < len(word); i++ {
			if strings.ToLower(letter) == strings.ToLower(string(word[i])) {
				guessed[i] = true
				numGuessed++
				inWord = true
			}
		}
		if !inWord {
			wrong++
		}
	}
	fmt.Print("Number of letters guessed: ")
	fmt.Println(numGuessed)
	fmt.Print("Number of incorrect guesses: ")
	fmt.Println(wrong)
	fmt.Println()
}

// duplicateLetter checks if the user's latest guess has already been made.
func (list *linkedList) duplicateLetter(letter string) bool {
	if list.head == nil {
		return true
	}
	for current := list.head; current != nil; current = current.next {
		if letter == current.value {
			return false
		}
	}
	return true
}

// displayHyphen prints out any letters which have been correctly guessed
// by the user, and the hyphens for each letter in the word
func displayHyphen(word string) {
	for j := 0; j < len(word); j++ {
		currLetter := string(word[j])
		runeLetter := []rune(currLetter)
		if guessed[j] {
			// print the letter if they've guessed it
			fmt.Print(currLetter)
			fmt.Print(" ")
		} else if !unicode.IsLetter(runeLetter[0]) {
			// if this letter is a special char, print it
			guessed[j] = true
			fmt.Print(currLetter)
			fmt.Print(" ")
			numGuessed++
		} else {
			// print blank space otherwise
			fmt.Print("  ")
		}
	}
	fmt.Println()
	// print out dashes for the number of letters in the word
	for i := 0; i < len(word); i++ {
		fmt.Print("- ")
	}
	fmt.Println()
}

// adds the words from dictionary to the map
func addMap() {
	file, _ := os.Open(dict)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mapping[scanner.Text()] = true
	}
}

/* pickRandom() picks a random word from dictionary map
    by iterating through the map and breaking out right
	away. Iteration of maps are designed to return the
	values in a different order every time */
func pickRandom() string {
	for key, _ := range mapping {
		return key
		break
	}
	return ""
}

// drawHangman prints out the body parts of the man on the gallows based on
// the number of incorrect letter guesses
func drawHangman(numLost int) {
	if numLost == 0 {
		fmt.Printf("  +---+\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	} else if numLost == 1 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	} else if numLost == 2 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	} else if numLost == 3 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	} else if numLost == 4 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	} else if numLost == 5 {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" /    |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")
	} else {
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf(" /|\\  |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" / \\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("R.I.P |     You Lose!\n")
		fmt.Printf("========\n")
	}
}

// addWordToFile() permanently adds a new word to the dictionary file
// and to the map, only if that word does not already exist in the map
func addWordToFile(newWord string) {
	var justWord string
	// grab just the word to add, without the newline
	for i := 0; i < len(newWord)-1; i++ {
		justWord += string(newWord[i])
	}
	// check if this word is already in the dictionary
	if checkDup(justWord) == true {
		fmt.Println("Sorry, this word already exists in the dictionary")
	} else {
		// add new word to the map
		mapping[justWord] = true

		//Creates a new file
		file, err := os.Create(dict)
		if err != nil {
			fmt.Println(err)
		}

		//adds all the words to the file from map
		for key, _ := range mapping {
			_, err := file.WriteString(key + "\n")
			if err != nil {
				fmt.Println(err)
			}
		}
		defer file.Close()
		fmt.Println("Your word has been added to the dictionary.")
	}
}

// checkDup checks to see if the new word to add is already in the map, returns true
// if the word exists in the map, false otherwise
func checkDup(newWord string) bool {
	_, found := mapping[newWord]
	return found
}

func main() {
	keepPlaying := true
	addMap()
	// loop as long as user wants to continue playing
	for keepPlaying {
		// prompt user for what they want to do
		fmt.Print("Do you want to play a game (P), add a word to the dictionary (A), or quit (Q)?: ")
		reader := bufio.NewReader(os.Stdin)
		userChoice, _ := reader.ReadString('\n')
		userChoice = strings.ToLower(string(userChoice[0]))
		if len(userChoice) > 1 || !(userChoice == "p" || userChoice == "a" || userChoice == "q") {
			fmt.Println("Sorry, you've entered an invalid response.")
		}
		switch userChoice {
		case "a":
			// call function to add their word to the dictionary permanently
			fmt.Println("Please enter the word you wish to include in the dictionary")
			reader := bufio.NewReader(os.Stdin)
			userWord, _ := reader.ReadString('\n')
			addWordToFile(userWord)
		case "q":
			// end the game
			keepPlaying = false
		case "p":
			// declare/assign necessary variables for a game and select random word from our dictionary
			list := linkedList{}
			randomWord := pickRandom()
			numGuessed = 0
			wrong = 0
			guessed = make(map[int]bool)

			// loop as long as they have not guessed every letter or exceeded
			// number of allowable incorrect guesses
			for wrong < 6 && numGuessed < len(randomWord) {
				displayHyphen(randomWord)

				fmt.Print("Please enter a letter: ")
				reader := bufio.NewReader(os.Stdin)
				response, _ := reader.ReadString('\n')
				if len(response) > 2 {
					fmt.Println("Sorry, you must enter a single letter.")
				} else {
					// get just the letter, not the carriage return
					response = string(response[0])
					fmt.Println()
					list.checkLetter(response, randomWord)
					fmt.Print("You have guessed the following letters: ")
					list.printLinkedList()
					fmt.Println()
				}
				drawHangman(wrong)
				// print whether they won or lost
				if wrong == 6 {
					fmt.Println("Sorry, you've exceeded the maximum guesses. You Lost!")
					fmt.Println("Your word was " + randomWord + ". Better luck next time!")
				} else if numGuessed == len(randomWord) {
					fmt.Println("You have successfully guessed all the letters. You Won!")
					displayHyphen(randomWord)
				}
			}
		}
	}
	fmt.Println("Goodbye!")
}
