package main

import (
	"fmt"
	"io"
	"log"
	"os"
        "math"
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

// LevelOrderTraversal method

func (t *Tree) LevelOrderTraversal(w io.Writer) {
	var s []*Node
	s = append(s, t.root)
	for len(s) > 0 {
		root := s[0]
		s = s[1:]
		fmt.Fprintf(w, "%d", root.data)
		if root.left != nil {
			s = append(s, root.left)
		}
		if root.right != nil {
			s = append(s, root.right)
		}
	}
}

// LevelOrderTraversalWithNewLine()

func (t *Tree) LevelOrderTraversalWithNewLine(w io.Writer) {
	// 		5
	//	  3    7
	// 1
	// q = [5,nil,]
	// q = [nil,3,7]
	// q = [3,7,nil]
	// q = [7,nil,1]
	// q = [nil,1]
	// q = [1,nil]
	// q = [nil]
	var q []*Node
	q = append(q, t.root)
	q = append(q, nil)
        // > 1 is used, to avoid infinite loop at the end of the queue due to repated null getting appended
	for len(q) > 1 {
		root := q[0]
		q = q[1:]
		if root == nil {
			fmt.Fprintf(w, "\n")
			q = append(q, nil)
		} else {
			fmt.Fprintf(w, "%d", root.data)
			if root.left != nil {
				q = append(q, root.left)
			}
			if root.right != nil {
				q = append(q, root.right)
			}
		}

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

// Height method returns height of the node, which is the max distance from root to it's leaf nodes

func (t *Tree) Height(node *Node) float64 {
	if node == nil {
		return -1
	}
	return math.Max(t.Height(node.left), t.Height(node.right)) + 1
}

// Search k in Tree

func (t *Tree) Search(k int, node *Node) bool {
	if node != nil {
		if node.data == k {
			return true
		} else if r := t.Search(k, node.left); r {
			return true
		} else if r := t.Search(k, node.right); r {
			return true
		}
	}
	return false
}

// Find path to 'k', same function as before, but print the nodes which return true

func (t *Tree) FindPath(w io.Writer, node *Node, k int) bool {
	if node != nil {
		if node.data == k {
			fmt.Fprintf(w, "%d", node.data)
			return true
		} else if r := t.FindPath(w, node.left, k); r {
			fmt.Fprintf(w, "%d", node.data)
			return true
		} else if r := t.FindPath(w, node.right, k); r {
			fmt.Fprintf(w, "%d", node.data)
			return true
		}
	}
	return false
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
        fmt.Println()
	h := t.Height(t.root)
	fmt.Println(h)
	t.LevelOrderTraversal(os.Stdout)
        fmt.Println()
	t.LevelOrderTraversalWithNewLine(os.Stdout)
        fmt.Println()
        r := t.Search(4, t.root)
	fmt.Println(r)
        fmt.Println()
        t.FindPath(os.Stdout, t.root, 2)
}

