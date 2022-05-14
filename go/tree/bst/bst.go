package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Node Struct

type Node struct {
	left, right *Node
	data        int
}

// Tree Struct

type Tree struct {
	root *Node
}

// CreateNode() Method

func (t *Tree) CreateNode(data int) *Node {
	return &Node{
		data: data,
	}
}

// InsertNode() Method

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = t.CreateNode(data)
	} else {
		t.InsertNode(t.root, data)
	}
}

// Insert() Method with recursion

func (t *Tree) InsertNode(node *Node, data int) *Node {
	// Base condition for recursion
	if node == nil {
		return t.CreateNode(data)
	}
	if data < node.data {
		// Recursion
		node.left = t.InsertNode(node.left, data)
	} else {
		// Recursion
		node.right = t.InsertNode(node.right, data)
	}
	return node
}

// InOrderTraversal() Method

func (t *Tree) InOrderTraversal(w io.Writer, root *Node) {
	if root != nil {
		t.InOrderTraversal(w, root.left)
		fmt.Fprint(w, root.data)
		t.InOrderTraversal(w, root.right)
	}
}

// PreOrderTraversal() Method

func (t *Tree) PreOrderTraversal(w io.Writer, root *Node) {
	if root != nil {
		fmt.Fprint(w, root.data)
		t.PreOrderTraversal(w, root.left)
		t.PreOrderTraversal(w, root.right)
	}
}

// PostOrderTraveral() Method

func (t *Tree) PostOrderTraversal(w io.Writer, root *Node) {
	if root != nil {
		t.PostOrderTraversal(w, root.left)
		t.PostOrderTraversal(w, root.right)
		fmt.Fprint(w, root.data)
	}
}

// PrintInOrder() Method

func (t *Tree) Print(method string, w io.Writer) {
	switch method {
	case "io":
		t.InOrderTraversal(w, t.root)
	case "pro":
		t.PreOrderTraversal(w, t.root)
	case "poo":
		t.PostOrderTraversal(w, t.root)
	default:
		fmt.Print("Unknown Print Method")
	}
}

// Count Method return count of nodes

func (t *Tree) Count(node *Node) int {
	// Count Of Nodes Of A Tree = Count(root.left) + 1 + Count(root.right)
	if node != nil {
		return t.Count(node.left) + 1 + t.Count(node.right)
	}
	return 0
}

func main() {
	t := Tree{}
	/*t.Insert(5)
	t.Insert(3)
	t.Insert(7)
	t.Insert(1)*/
	n := []int{5, 3, 7, 1}
	for _, val := range n {
		t.Insert(val)
	}
	t.Print("io", os.Stdout)
	fmt.Print("\n")
	t.Print("pro", os.Stdout)
	fmt.Print("\n")
	t.Print("poo", os.Stdout)
        fmt.Println()
	c := t.Count(t.root)
	fmt.Println(c)
}

