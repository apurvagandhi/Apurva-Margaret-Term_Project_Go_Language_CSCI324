/* Apurva Gandhi and Margaret Haley
Go Program for Hangman

TaskList:
- Make game graphical
- add actual hangman
- sort linked list
- error checking for user entered response
- check to see if tree is forming correctly
- Once game works, implement better algorithm to pick random word from insert tree

*/

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	//"image/draw"
)

var guessed = make(map[int]bool)
var wrong int = 0
var numGuessed int = 0
var myLinkedlist = list.New()
var t Tree

func printLinkedList() {
	for element := myLinkedlist.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value)
	}
}

func CheckAndAddLetter(newLetter string) bool {
	for element := myLinkedlist.Front(); element != nil; element = element.Next() {
		if newLetter == element.Value {
			return false
		}
	}
	// change this to add in sorted order
	myLinkedlist.PushFront(newLetter)
	fmt.Println("Guessed Letter Added Successfully to the Linked List")
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
	dictionaryWords, err := ioutil.ReadFile("smallDictionary.txt")
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

func checkLetter(letter string, word string) {
	if !(CheckAndAddLetter(letter)) {
		fmt.Println("You have already guessed this letter, try a new one.")
	} else {
		inWord := false
		for i := 0; i < len(word); i++ {
			//fmt.Println("Letter you guessed is: " + strings.ToLower(letter))
			//fmt.Println("The current letter of the word is: " + string(word[i]))
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
	fmt.Print("# right is: ")
	fmt.Println(numGuessed)
	fmt.Print("# wrong is: ")
	fmt.Println(wrong)

	//fmt.Print("current map of letters guessed in word: ")
	//fmt.Println(guessed)
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
	word, err := os.Open("smallDictionary.txt")
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

func main() {
	word := "hello"
	randomWord := pickRandomWordFromDict()
	fmt.Println("Random word is " + randomWord)
	dictionaryToBinaryTree("smallDictionary.txt")
	printPostOrder(t.root)
	for wrong < 6 && numGuessed < len(word) {
		displayHyphen(word)

		fmt.Print("Please enter a letter: ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')

		// get just the letter, not the carriage return
		for i := 0; i < 1; i++ {
			response = string(response[i])
		}

		checkLetter(response, word)
		//fmt.Print("Here are the letters you have already guessed: ")
		fmt.Print("LINKED LIST CONTAINS: ")
		printLinkedList()
		fmt.Println()

	}
	if wrong > 6 {
		fmt.Println("Sorry, you have exceeded the maximum guesses. You Lost!")
	}
	if wrong < 6 && numGuessed == len(word) {
		fmt.Println("You have successfully guessed all the letters. You Won!")
		displayHyphen(word)
	}
}

/* Sources
https://www.bogotobogo.com/GoLang/GoLang_Binary_Search_Tree.php

https://flaviocopes.com/golang-data-structure-binary-search-tree/

https://zetcode.com/golang/readfile/

https://hackernoon.com/how-to-select-a-random-node-from-a-tree-with-equal-probability-childhood-moments-with-father-today-0ip32dp

https://www.socketloop.com/tutorials/golang-convert-uint-value-to-string-type#:~:text=You%20have%20an%20uint%20value,the%20uint64%20value%20to%20string.
*/
