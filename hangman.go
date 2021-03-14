/* Apurva Gandhi and Margaret Haley
Go Program for Hangman

TaskList:
1. Create a function to reverse a string
2. Insert words into the binary tree by comparing reversed strings
3. Using Algorithm, randomly select a node from binary search tree
4. Make game graphical
5. numGuessed Needs to be fixed if letter guessed more than once
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var guessed = make(map[int]bool)
var wrong int = 0
var numGuessed int = 0
var ourList *List = nil

type LLNode struct {
	letterGuessed string
	next          *LLNode
}

type List struct {
	head *LLNode
	tail *LLNode
}

func (L *List) isDuplicate(newLetter string) bool {
	if L.head == nil {
		return false
	} else {
		for L != nil {
			if newLetter == L.head.letterGuessed {
				fmt.Println("You have already guessed this letter, please try another!")
				numGuessed--
				return true
			} else {
				L.head = L.head.next
			}
		}
		// subtract num guessed if this letter is in the word
	}
	return false
}

func (L *List) addLetter(newLetter string) {
	newNode := &LLNode{letterGuessed: newLetter}
	if !(L.isDuplicate(newLetter)) {
		if L.head == nil {
			L.head = newNode
		} else if newNode.letterGuessed < L.head.letterGuessed {
			newNode.next = L.head
			L.head = newNode
		} else {
			current := L.head
			for current.next != nil && current.next.letterGuessed < newNode.letterGuessed {
				current = current.next
			}
			newNode.next = current.next
			current.next = newNode
		}
	}
}

func displayHyphen(word string) {
	for j := 0; j < len(word); j++ {
		if guessed[j] == true {
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
	fmt.Print("Here are the letters you have guessed: ")

}

func checkLetter(letter string, word string, L *List) {
	L.addLetter(letter)
	inWord := false
	for i := 0; i < len(word); i++ {
		if strings.ToLower(letter) == strings.ToLower(string(word[i])+"\n") {
			guessed[i] = true
			numGuessed++
			inWord = true
		}
	}
	if !inWord {
		wrong++
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
	if data <= n.key {
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
	for wrong < 6 && numGuessed < len(word) {
		displayHyphen(word)
		fmt.Print("Please enter a letter: ")
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')

		checkLetter(response, word, ourList)
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
