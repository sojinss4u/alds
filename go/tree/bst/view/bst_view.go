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

// LeftView() method for the tree

func (t *Tree) LeftView(w io.Writer) {
	//     		5       	 level = 0
	//  	 3  	 7		 level = 1
	//	 1       		8    level = 2
	// 0   2         6     9
	// i := [537180269]
	// op := 5310
	var q []*Node
	h := make(map[int]int)
	q = append(q, t.root)
	for len(q) != 0 {
		root := q[0]
		q = q[1:]
		if _, ok := h[root.level]; !ok {
			h[root.level] = root.data
		}
		if root.left != nil {
			root.left.level = root.level + 1
			q = append(q, root.left)
		}
		if root.right != nil {
			root.right.level = root.level + 1
			q = append(q, root.right)
		}
	}
	// Since map is an unordered collection of key value pairs, we can't gurantee the order in which the map elements will be printed.
	// Hence we need to append the hash keys into a slice, then sort it & iterate trough it to print the values of hash in a predicted order
	// https://yourbasic.org/golang/sort-map-keys-values/

	var k []int
	for keys, _ := range h {
		k = append(k, keys)
	}
	// Sort slice values in increasing order
	sort.Ints(k)
	for _, val := range k {
		fmt.Fprintf(w,"%d", h[val])
	}
}

// RightView() method

func (t Tree) RightView(w io.Writer) {
	//		    5
	// 		 3     7           <==== View
	//	   1    6     8
	// 	 0   2          9
	// RightView: 5789
	// Ans: LevelOrderTraversalRightToLeft + Hash[level][node]
	var q []*Node
	h := make(map[int]*Node)
	q = append(q, t.root)
	for len(q) != 0 {
		root := q[0]
		q = q[1:]
		if _, ok := h[root.level]; !ok {
			h[root.level] = root
		}
		if root.right != nil {
			root.right.level = root.level + 1
			q = append(q, root.right)
		}
		if root.left != nil {
			root.left.level = root.level + 1
			q = append(q, root.left)
		}
	}
	// Add the hash keys to a slice to sort them in order [has is an unordered list]
	var s []int
	for k, _ := range h {
		s = append(s, k)
	}
	sort.Ints(s)
	for _, val := range s {
		fmt.Fprintf(w, "%d", h[val].data)
	}
}

func main() {
	t := Tree{}
	i := []int{5, 3, 7, 1, 2, 6, 8}
        //           5
        //        3     7
        //     1     6    8
	//        2
	for _, val := range i {
		t.Insert(val)
	}
	t.InOrderTraversal(os.Stdout, t.root)
	fmt.Println()
	t.TopView(os.Stdout)
        fmt.Println()
	t.LeftView(os.Stdout)
        fmt.Println()
        t.RightView(os.Stdout)
}

