package main

import (
	"fmt"
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

func (t *Tree) InOrderTraversal(root *Node) {
	if root != nil {
		t.InOrderTraversal(root.left)
		fmt.Print(root.data)
		t.InOrderTraversal(root.right)
	}
}

// PreOrderTraversal() Method

func (t *Tree) PreOrderTraversal(root *Node) {
	if root != nil {
		fmt.Print(root.data)
		t.PreOrderTraversal(root.left)
		t.PreOrderTraversal(root.right)
	}
}

// PostOrderTraveral() Method

func (t *Tree) PostOrderTraversal(root *Node) {
	if root != nil {
		t.PostOrderTraversal(root.left)
		t.PostOrderTraversal(root.right)
		fmt.Print(root.data)
	}
}

// PrintInOrder() Method

func (t *Tree) Print(method string) {
	switch method {
	case "io":
		t.InOrderTraversal(t.root)
	case "pro":
		t.PreOrderTraversal(t.root)
	case "poo":
		t.PostOrderTraversal(t.root)
	default:
		fmt.Print("Unknown Print Method")
	}
}

func main() {
	t := Tree{}
	t.Insert(5)
	t.Insert(3)
	t.Insert(7)
	t.Insert(1)
	t.Print("io")
	fmt.Print("\n")
	t.Print("pro")
	fmt.Print("\n")
	t.Print("poo")
}

