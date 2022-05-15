package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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
	data, level int
}

// Tree Struct

type Tree struct {
	root *Node
}

// CreateNode Function

func CreateNode(data int) *Node {
	return &Node{
		data: data,
	}
}

// InsertNode() Method
// Uses recursion
func (t *Tree) InsertNode(node *Node, data int) *Node {
	if node == nil {
		return CreateNode(data)
	}
	if data < node.data {
		node.left = t.InsertNode(node.left, data)
	} else {
		node.right = t.InsertNode(node.right, data)
	}
	return node
}

// Insert() method

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = CreateNode(data)
	} else {
		t.InsertNode(t.root, data)
	}
}

// InOrder Traversal

func (t *Tree) InOrderTraversal(w io.Writer, node *Node) {
	// Print in [LeftRootRight] format
	// Base condition
	if node == nil {
		return
	}
	t.InOrderTraversal(w, node.left)
	fmt.Fprintf(w, "%d", node.data)
	t.InOrderTraversal(w, node.right)

}

// Print Top View Of Tree
func (t *Tree) TopView(w io.Writer) {
	var q []*Node
	h := make(map[int]int)
	q = append(q, t.root)
	for len(q) != 0 {
		root := q[0]
		q = q[1:]
		level := root.level
		if _, ok := h[level]; !ok {
			h[level] = root.data
		}
		if root.left != nil {
			q = append(q, root.left)
			root.left.level = level - 1
		}
		if root.right != nil {
			q = append(q, root.right)
			root.right.level = level + 1
		}
	}
	// Print values in hash
	var sl []int
	for i, _ := range h {
		sl = append(sl, i)
	}
	sort.Ints(sl)
	for _, val := range sl {
		fmt.Fprintf(w, "%d", h[val])
	}
}

func main() {
	t := Tree{}
	i := []int{5, 3, 7, 1, 2, 6, 8}
	for _, val := range i {
		t.Insert(val)
	}
	t.InOrderTraversal(os.Stdout, t.root)
	fmt.Println()
	t.TopView(os.Stdout)
}
