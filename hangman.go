/* Apurva Gandhi and Margaret Haley
Go Program for Hangman

TaskList:
1. Create a function to reverse a string
2. Insert words into the binary tree by comparing reversed strings
3. Using Algorithm, randomly select a node from binary search tree
4. Make game graphical
5. numGuessed Needs to be fixed if letter guessed more than once
6. Delete \n from entered response.
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
		if element.Value != "\n" {
			fmt.Print(element.Value)
		}
	}
}

func CheckAndAddLetter(newLetter string) bool {
	for element := myLinkedlist.Front(); element != nil; element = element.Next() {
		if newLetter == element.Value {
			return true
		}
	}
	myLinkedlist.PushFront(newLetter)
	fmt.Println("Guessed Letter Added Successfully to the Linked List")
	return false
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
	//fmt.Print("Here are the letters you have guessed: ")

}
func pickRandomWordFromDict() string {
	letters, err := ioutil.ReadFile("smallDictionary.txt")
	lettersString := string(letters)
	words := strings.Split(lettersString, "\n")
	totalWords := len(words) - 1
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randnum := r1.Intn(totalWords)
	fmt.Print("Random number is ")
	fmt.Println(randnum)

	randomWord := words[randnum]

	if err != nil {
		fmt.Println(err)
	}
	return randomWord

}

func checkLetter(letter string, word string) {
	checkRepeat := CheckAndAddLetter(letter)
	if checkRepeat {
		fmt.Println("You have already entered this guess, try new one")
	} else {

		inWord := false
		for i := 0; i < len(word); i++ {
			fmt.Println("Letter is " + strings.ToLower(letter))
			fmt.Println("letter of the word is " + string(word[i]) + "\n")
			//fmt.Println(strings.Compare(strings.ToLower(letter), strings.ToLower(string(word[i])+"\n")))
			fmt.Print("Letter equals letter of the word? ")
			fmt.Println(strings.ToLower(letter) == strings.ToLower(string(word[i])+"\n"))
			
			if strings.ToLower(letter) == strings.ToLower(string(word[i])+"\n") {
				//if (strings.Compare(strings.ToLower(letter), strings.ToLower(string(word[i])+"\n"))) == 0 {

				fmt.Println("Check")
				guessed[i] = true
				numGuessed++
				inWord = true
			}
		}

		if !inWord {
			wrong++
		}
	}
	fmt.Println("num guessed right, num guessed wrong:")
	fmt.Println(numGuessed, wrong)
	fmt.Println(guessed)
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
		checkLetter(response, word)
		fmt.Print("So far you entered following guesses: ")
		printLinkedList()

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
