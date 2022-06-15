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

// Empty struct is defined here
type void struct{}

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Node struct for representing vertices of the Graph
// Neighbours make use of a map to implement set()
type Node struct {
	value     string
	neighbors map[string]void
}

func (n *Node) AddEdge(u string) {
	if n.neighbors == nil {
		n.neighbors = make(map[string]void)
	}
	n.neighbors[u] = void{}
}

// Graph struct
// Adjacency List Implementation
type Graph struct {
	vertices map[string]*Node
}

// Method to add vertex to Graph

func (g *Graph) AddVertex(n *Node) {
	if _, ok := g.vertices[n.value]; !ok {
		if g.vertices == nil {
			g.vertices = make(map[string]*Node)
		}
		g.vertices[n.value] = n
	}
}

// Method to add edges for a given node
// Graph is assumed to be undirected Graph. So we need to add the edges for both the given edges

func (g *Graph) AddEdge(u, v *Node) {
	_, ok1 := g.vertices[u.value]
	_, ok2 := g.vertices[v.value]
	if ok1 && ok2 {
		g.vertices[u.value].AddEdge(v.value)
		g.vertices[v.value].AddEdge(u.value)
	}
}

func (g *Graph) Print(w io.Writer) {
	var s1 []string
	for _, val := range g.vertices {
		s1 = append(s1, val.value)
	}
	sort.Strings(s1)

	for _, val := range s1 {
		var s2 []string
		for k, _ := range g.vertices[val].neighbors {
			s2 = append(s2, k)
		}
		sort.Strings(s2)
		fmt.Fprintf(w, "%s : %v\n", g.vertices[val].value, s2)
	}
}

func main() {
	n1 := &Node{
		value: "A",
	}
	n2 := &Node{
		value: "B",
	}
	n3 := &Node{
		value: "C",
	}
	n4 := &Node{
		value: "D",
	}
	n5 := &Node{
		value: "E",
	}
	g := Graph{}
	g.AddVertex(n1)
	g.AddVertex(n2)
	g.AddVertex(n3)
	g.AddVertex(n4)
	g.AddVertex(n5)
	g.AddEdge(n1, n2)
	g.AddEdge(n1, n5)
	g.AddEdge(n2, n3)
	g.AddEdge(n4, n3)
	g.AddEdge(n5, n3)
	g.Print(os.Stdout)
}
