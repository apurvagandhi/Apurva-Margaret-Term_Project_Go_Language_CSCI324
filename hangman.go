/* Apurva Gandhi and Margaret Haley
Go Program for Hangman

TaskList:
- Once game works, implement algorithm to pick random word from insert tree NOT file
- Allow user to add word to the dictonary (by inserting the word into the binary tree???)
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var guessed = make(map[int]bool)
var wrong int = 0
var numGuessed int = 0
var t Tree
var dict string = "test.txt"

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
   guessed (which are revealed) and how many incorrect guessed he/she has made.
*/
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
		if guessed[j] {
			// print the letter
			fmt.Print(string(word[j]))
			fmt.Print(" ")
		} else {
			// print blank space
			fmt.Print("  ")
		}
	}
	fmt.Println()
	for i := 0; i < len(word); i++ {
		fmt.Print("- ")
	}
	fmt.Println()
}

/* The following structs and functions are used to select our random word from the
   dictionary and to create and store our dictionary into a binary search tree.
*/

func pickRandomWordFromDict() string {
	dictionaryWords, err := ioutil.ReadFile(dict)
	dictionary := string(dictionaryWords)
	words := strings.Split(dictionary, "\n")
	totalWords := len(words) - 1
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randNum := r1.Intn(totalWords)
	//fmt.Print("Random number is ")
	//fmt.Println(randNum)

	randomWord := words[randNum]

	if err != nil {
		fmt.Println(err)
	}
	return randomWord

}

type Tree struct {
	root *Node
}

type Node struct {
	key   string
	left  *Node
	right *Node
}

// Tree creation
func (t *Tree) insert(data string) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Node creation
func (n *Node) insert(data string) {
	if data == n.key {
		// we dont want to accept this new word
	} else if reverse(data) < reverse(n.key) {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

// Prints our binary search tree, used for testing purposes
func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Println(n.key)
	}
}

// reverse takes in a string and returns the same string except reversed
func reverse(str string) (result string) {
	for _, l := range str {
		result = string(l) + result
	}
	return
}

// dictionaryToBinaryTree reads in the dictionary file and produces a BST from it
func dictionaryToBinaryTree(fileName string) {
	word, err := os.Open(dict)
	// print if there's an error
	if err != nil {
		fmt.Println(err)
	}
	// close file descriptor at the end of the main function
	defer word.Close()
	// create new scanner, split content by words
	scanner := bufio.NewScanner(word)
	scanner.Split(bufio.ScanWords)
	// read each word and insert it into the BST
	for scanner.Scan() {
		t.insert(scanner.Text())
	}
	// check for error in scanning
	scanErr := scanner.Err()
	if err != nil {
		fmt.Println(scanErr)
	}
}

// searchTree returns whether or not a certain word is in the BST
func searchTree(n *Node, word string) bool {
	if n == nil {
		return false
	} else if strings.ToLower(word) < strings.ToLower(n.key) {
		return searchTree(n.left, word)
	} else if strings.ToLower(word) > strings.ToLower(n.key) {
		return searchTree(n.right, word)
	} else {
		// only alternative is that they must be equal, so return true
		return true
	}
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

func addWordToFile(newWord string) {
	dictionaryWords, err := ioutil.ReadFile(dict)
	dictionary := string(dictionaryWords)

	file, err := os.Create(dict)
	if err != nil {
		fmt.Println(err)
	}
	_, err2 := file.WriteString(dictionary)
	if err2 != nil {
		fmt.Println(err)
	}
	_, err3 := file.WriteString(newWord)
	if err3 != nil {
		fmt.Println(err3)
	}
	defer file.Close()
	fmt.Println("Your word has been added to the dictionary.")
}

func main() {
	keepPlaying := true
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
			randomWord := pickRandomWordFromDict()
			dictionaryToBinaryTree(dict)
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

/* Sources
https://www.bogotobogo.com/GoLang/GoLang_Binary_Search_Tree.php

https://flaviocopes.com/golang-data-structure-binary-search-tree/

https://zetcode.com/golang/readfile/

https://hackernoon.com/how-to-select-a-random-node-from-a-tree-with-equal-probability-childhood-moments-with-father-today-0ip32dp

https://www.socketloop.com/tutorials/golang-convert-uint-value-to-string-type#:~:text=You%20have%20an%20uint%20value,the%20uint64%20value%20to%20string.

https://tutorialedge.net/golang/go-linked-lists-tutorial/

for all the list stuff:
https://dev.to/divshekhar/golang-linked-list-data-structure-h20

ascii hangman: https://inventwithpython.com/invent4thed/chapter8.html

Random Node from a Tree Algorithm: https://hackernoon.com/how-to-select-a-random-node-from-a-tree-with-equal-probability-childhood-moments-with-father-today-0ip32dp

*/
