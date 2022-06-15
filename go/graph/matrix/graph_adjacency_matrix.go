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

// Node type is implemented here
type Node struct {
	value string
}

type Graph struct {
	vertices    map[string]*Node
	edges       [][]int
	edgeIndexes map[string]int
}

func (g *Graph) AddVertex(n *Node) {
	if _, ok := g.vertices[n.value]; !ok {
		if g.vertices == nil {
			g.vertices = make(map[string]*Node)
		}
		g.vertices[n.value] = n
		// Append 0 to all existing elements as we have one more vertex getting added
		for i, _ := range g.edges {
			g.edges[i] = append(g.edges[i], 0)
		}
		// Add one more element of the same length as above to make the matrix n * n
		g.edges = append(g.edges, make([]int, len(g.edges)+1))
		// Initialize map to avoid "assignment to entry in nil map" exception
		if g.edgeIndexes == nil {
			g.edgeIndexes = map[string]int{}
		}
		// Store the index of the slice which represent the added node
		g.edgeIndexes[n.value] = len(g.edgeIndexes)
	}
}

func (g *Graph) AddEdge(u, v *Node, w int) {
	_, ok1 := g.vertices[u.value]
	_, ok2 := g.vertices[v.value]
	if ok1 && ok2 {
		g.edges[g.edgeIndexes[u.value]][g.edgeIndexes[v.value]] = w
		g.edges[g.edgeIndexes[v.value]][g.edgeIndexes[u.value]] = w
	}
}

func (g *Graph) Print(w io.Writer) {
	var s []string
	for k1, _ := range g.edgeIndexes {
		s = append(s, k1)
	}
	sort.Strings(s)
	for _, k2 := range s {
		fmt.Fprintf(w, "%s : %v\n", k2, g.edges[g.edgeIndexes[k2]])
	}
}

func main() {
	n1 := &Node{value: "A"}
	n2 := &Node{value: "B"}
	n3 := &Node{value: "C"}
	n4 := &Node{value: "D"}
	n5 := &Node{value: "E"}
	g := Graph{}
	g.AddVertex(n1)
	g.AddVertex(n2)
	g.AddVertex(n3)
	g.AddVertex(n4)
	g.AddVertex(n5)
	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n5, 1)
	g.AddEdge(n2, n3, 1)
	g.AddEdge(n3, n4, 1)
	g.AddEdge(n5, n3, 1)
	g.Print(os.Stdout)
}

