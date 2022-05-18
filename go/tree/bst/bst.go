package main

import (
	"fmt"
	"io"
	"log"
	"os"
        "math"
        "bytes"
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

func (t *Tree) LevelOrderTraversalLeftToRight(w io.Writer) {
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

// LevelOrderTraversalRightToLeft to print elements in each level right to left

func (t *Tree) LevelOrderTraversalRightToLeft(w io.Writer) {
	var q []*Node
	q = append(q, t.root)
	for len(q) > 0 {
		root := q[0]
		q = q[1:]
		fmt.Fprintf(w, "%d", root.data)
		if root.right != nil {
			q = append(q, root.right)
		}
		if root.left != nil {
			q = append(q, root.left)
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

// LCA() Method

func (t *Tree) LCA(w io.Writer, a, b int) {
	// LCA => Least Common Ancestor
	// To find LCA, Find the path between root to each node & find the node which is commeon for both of them & close to the nodes.
	//     5
	//   3   7
	// 1   2   9
	// Path(1) = [5,3,1]
	// Path(2) = [5,3,2]
	// LCA(1,2) = 3

	var s1 []string
	var s2 []string
	var b1 bytes.Buffer
	var b2 bytes.Buffer
	t.FindPath(&b1, t.root, a)
	t.FindPath(&b2, t.root, b)
	for i := len(b1.String()) - 1; i >= 0; i -= 1 {
		s1 = append(s1, string(b1.String()[i]))
	}
	for i := len(b2.String()) - 1; i >= 0; i -= 1 {
		s2 = append(s2, string(b2.String()[i]))
	}
	var lca string
	var minLength int
	s1Length := len(s1)
	s2Length := len(s2)
	if s1Length < s2Length {
		minLength = s1Length
	} else {
		minLength = s2Length
	}
	for i := 0; i < minLength; i += 1 {
		if s1[i] == s2[i] {
			lca = s1[i]
		}
	}
	fmt.Fprintf(w, "%s", lca)
}

// PathBetween() method

func (t *Tree) PathBetween(w io.Writer, a, b int) {
	// In order to find the path between 'a' & 'b', we will find the path between root & each nodes.
	// Then we will, find the path from each node to LCA. Then we will combine, the paths to get the
	// path between a & b.
	//     		5       	   level = 0
	//  	3  	    7		   level = 1
	//	 1      6	   8       level = 2
	// 0     2             9   level = 3
	// PathFromRootNode(0) = [5,3,1,0]
	// PathFromRootNode(2) = [5,3,1,2]
	// PathBetween(0,2) = 0 1[LCA] 2
	var s1 []string
	var s2 []string
	var bf1 bytes.Buffer
	var bf2 bytes.Buffer
	var bl bytes.Buffer

	t.FindPath(&bf1, t.root, a)
	t.FindPath(&bf2, t.root, b)
	for i := 0; i < len(bf1.String()); i += 1 {
		s1 = append(s1, string(bf1.String()[i]))
	}
	for i := len(bf2.String()) - 1; i >= 0; i -= 1 {
		s2 = append(s2, string(bf2.String()[i]))
	}
	// Find LCA(a) & LCA(b)
	t.LCA(&bl, a, b)
	lca := bl.String()
	for _, val := range s1 {
		if val == lca {
			break
		}
		fmt.Fprintf(w, string(val))
	}
	fmt.Fprintf(w, lca)
	var p bool
	for _, val := range s2 {
		if p {
			fmt.Fprintf(w, string(val))
		}
		if val == lca {
			p = true
		}
	}
}

// CheckBSTNode() Method

func (t *Tree) CheckBSTNode(node *Node, rangeStart, rangeEnd int, side string) bool {
	//     		 5
	//  	 3  	 7
	//	  1      6	     8
	// 0     2                 9
	// Option 1: Do in order traversal & see if the o/p is sorted. For a BST InOrderTraversl o/p will be always sorted.
	// Eg: InOrderTraversal OP: 012356789
	// Option 2: Second option is to identify the range in which each node should fall. This range is created by using the
	// property of the BST, which says root.left<root.data<root.data.
	//       a[-IntMax,+IntMax]
	//    b[-IntMax,a]     c [a,+IntMax]
	// d[-InMax,b]     e      f[c,+IntMax]
	// If you move from a node to left, the child node range will be [left-range-value-of-parent-node, parent-node-data] &
	// If you move from a node to right, the child node range will be [parent-node-data, right-range-value-of-parent-node].
	// For the second method, time complexity will be O(n), as we visit each node in the worst case & space complexity will be O(1)

	if node == nil {
		return true
	}

	if side == "left" {
		if node.data > rangeStart && node.data < rangeEnd {
			resultLR := t.CheckBSTNode(node.left, rangeStart, node.data, "left") && t.CheckBSTNode(node.right, node.data, rangeEnd, "right")
			if resultLR {
				return true
			}
		}
	} else {
		if node.data >= rangeStart && node.data < rangeEnd {
			resultLR := t.CheckBSTNode(node.left, rangeStart, node.data, "left") && t.CheckBSTNode(node.right, node.data, rangeEnd, "right")
			if resultLR {
				return true
			}
		}
	}
	return false
}

// Wrapper function to call CheckBST method with relevant values 

func (t *Tree) CheckBST() bool {

	return t.CheckBSTNode(t.root.left, math.MinInt, t.root.data, "left") && t.CheckBSTNode(t.root.right, t.root.data, math.MaxInt, "right")
}

func main() {
	t := Tree{}
	/*t.Insert(5)
	t.Insert(3)
	t.Insert(7)
	t.Insert(1)*/
	n := []int{5, 3, 7, 1}
        //     5
        //   3   7
        // 1
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
        t.LevelOrderTraversalLeftToRight(os.Stdout)
        fmt.Println()
	t.LevelOrderTraversalWithNewLine(os.Stdout)
        fmt.Println()
        r := t.Search(4, t.root)
	fmt.Println(r)
        fmt.Println()
        t.FindPath(os.Stdout, t.root, 2)
        fmt.Println()
	t.LevelOrderTraversalRightToLeft(os.Stdout)
        fmt.Println()
        t.LCA(os.Stdout, 3, 7)
        fmt.Println()
	t.PathBetween(os.Stdout, 1, 7)
        fmt.Println()
        r1 := t.CheckBST()
	fmt.Println(r1)
}

