/* The following structs and functions are used to select our random word from the
   dictionary and to create and store our dictionary into a binary search tree.
*/

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