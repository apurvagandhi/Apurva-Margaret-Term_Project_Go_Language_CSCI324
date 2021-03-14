/* Apurva Gandhi and Margaret Haley
Go Program for Hangman

TaskList:
- Once game works, implement algorithm to pick random word from insert tree NOT file
- Allow user to add word to the dictonary (by inserting the word into the binary tree)
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
var dict string = "smallDictionary.txt"

type node struct {
	value string
	next  *node
}

type linkedList struct {
	head *node
}

func (list linkedList) printLinkedList() {
	for list.head != nil {
		fmt.Print(list.head.value)
		fmt.Print(" ")
		list.head = list.head.next
	}
	fmt.Println()
}

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

//Tree
func (t *Tree) insert(data string) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

//Node
func (n *Node) insert(data string) {
	if reverse(data) <= reverse(n.key) {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else { // if greater than bc we dont want to accept the insert if they're equal
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Println(n.key)
	}
}

func reverse(str string) (result string) {
	for _, l := range str {
		result = string(l) + result
	}
	return
}

func dictionaryToBinaryTree(fileName string) {
	word, err := os.Open(dict)
	//if error, prints the error
	if err != nil {
		fmt.Println(err)
	}
	//The file descriptor is closed at the end of the main function.
	defer word.Close()
	// A new scanner is created.
	scanner := bufio.NewScanner(word)
	//split the content by words.
	scanner.Split(bufio.ScanWords)
	//The Scan advances the Scanner to the next token,
	//which will then be available through the Text function.

	for scanner.Scan() {
		t.insert(scanner.Text())
		//fmt.Println("Inserted Successfully")
	}
	//Checks for Error in scanning
	scanErr := scanner.Err()
	if err != nil {
		fmt.Println(scanErr)
	}

}
func searchTree(n *Node, word string) bool {
	if n == nil {
		return false
	}
	if strings.ToLower(word) < strings.ToLower(n.key) {
		return searchTree(n.left, word)
	}
	if strings.ToLower(word) > strings.ToLower(n.key) {
		return searchTree(n.right, word)
	}
	return true
}

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

func main() {
	list := linkedList{}
	randomWord := pickRandomWordFromDict()
	fmt.Println("Random word is " + randomWord)
	dictionaryToBinaryTree(dict)
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
	}

	if wrong > 6 {
		fmt.Println("Sorry, you've exceeded the maximum guesses. You Lost!")
	} else if numGuessed == len(randomWord) {
		fmt.Println("You have successfully guessed all the letters. You Won!")
		displayHyphen(randomWord)
	}
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

*/
