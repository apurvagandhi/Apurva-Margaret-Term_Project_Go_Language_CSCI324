// linked list

package main

import (
	"fmt"
    "os"
)
else if string(word[j]) == "%" {
			guessed[j] = true
			fmt.Print(string(word[j]))
			fmt.Print(" ")
		}

/*
type node struct {
	value string
	next  *node
}

type linkedList struct {
	head *node
}

func (list linkedList) printList() {
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
	} else if letter < list.head.value { // A < B
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

func (list *linkedList) duplicateLetter(letter string) bool {
	for current := list.head; current.next != nil; current = current.next {
		if letter == current.value {
			return true
		}
	}
	return false
}
func main() {
	list := linkedList{}
	list.insertLetter("c")
	list.insertLetter("b")

	list.insertLetter("z")

	list.insertLetter("a")
	fmt.Println(list.duplicateLetter("b"))

	list.printList()
}

/*Apruva
func (l *LinkedList) insert(insertData string) {
	n := Node{}
	n.data = insertData

	if l.head == nil {
		l.head = n
	}
	else if insertData < l.head.value{
		n.next = l.head
		l.head = new
	}
	else {
		current := l.head
		for

	}


}
*/

/*
func duplicateLetter(newLetter string) bool {
	for element := myLinkedList.Front(); element != nil; element = element.Next() {
		if newLetter == element.Value {
			return false
		}
	}
	insertLetter(newLetter)
	return true
}*/

/*
func insertLetter(letter string) {
	//myLinkedlist.PushFront(newLetter)
	current := myLinkedList.Front()
	if myLinkedList.Front == nil {
		current.Value = letter
	} else if letter < current.Value { // A < B
		myLinkedList.InsertBefore(letter, current)
	} else {
		for current.Next != nil && strings.Compare(letter, current.Next.Value) == -1 {
			current = current.Next
		}
		myLinkedList.InsertBefore(letter, current)
	}
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
		fmt.Printf("R.I.P |\n")
		fmt.Printf("========\n")
	}
}
*/

/*
Traverse through the whole tree and save the total number of nodes for all the subtrees in the tree. 
The total number of nodes for all the subtrees can be saved in any format and that takes the space O(n).
*/

func addWordToFile (newWord string) {
	file, err := os.Create(dict)
	if err =! nil {
		fmt.Println(err)
	}
	defer file.close
	_, err2 := file.WriteString(newWord)
	if err2 != nil {
        fmt.Println(err)
    }
	 fmt.Println("Word Added Successfully")

}


func addWordToFile(newWord string) {
	
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
	file, err := os.Create(dict)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	for scanner.Scan() {
		_, err2 := file.WriteString(newWord)
		if err2 != nil {
			fmt.Println(err)
		}
	}
	// check for error in scanning
	scanErr := scanner.Err()
	if err != nil {
		fmt.Println(scanErr)
	}
	fmt.Println("Word " + newWord + "Added Successfully")
}

func duplicateWordInFile(newWord string) bool {
	fmt.Println("You wanna check word: " + newWord)
	
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
		if scanner.Text() == newWord {
			return true
		}
	}
	// check for error in scanning
	scanErr := scanner.Err()
	return false
}