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
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ltime|log.Ltime|log.Lshortfile)
}

type Node struct {
	value, color              string
	neighbors                 map[string]int
	discoveryTime, finishTime int
}

func (n *Node) AddEdge(u string, w int) {
	if n.neighbors == nil {
		n.neighbors = make(map[string]int)
	}
	n.neighbors[u] = w
}

type Graph struct {
	vertices       map[string]*Node
	discoveryTimes map[int]string
}

func (g *Graph) AddVertex(v string) {
	if _, ok := g.vertices[v]; !ok {
		if g.vertices == nil {
			g.vertices = make(map[string]*Node)
		}
		n := &Node{
			value: v,
			color: "black",
		}
		g.vertices[v] = n
	}
}

func (g *Graph) AddEdge(u, v string, w int) {
	_, ok1 := g.vertices[u]
	_, ok2 := g.vertices[v]
	if ok1 && ok2 {
		g.vertices[u].AddEdge(v, w)
		g.vertices[v].AddEdge(u, w)
	}
}

func (g *Graph) Print(w io.Writer) {
	var s1 []string
	for k, _ := range g.vertices {
		s1 = append(s1, k)
	}
	sort.Strings(s1)
	for _, k := range s1 {
		var s2 []string
		for v, _ := range g.vertices[k].neighbors {
			s2 = append(s2, v)
		}
		sort.Strings(s2)
		fmt.Fprintf(w, "%s : %v\n", k, s2)
	}
}

func (g *Graph) dfs(v *Node, t int) int {
	v.color = "red"
	v.discoveryTime = t
	if g.discoveryTimes == nil {
		g.discoveryTimes = make(map[int]string)
	}
	g.discoveryTimes[t] = v.value
	t += 1
	var s []string
	for item, _ := range v.neighbors {
		s = append(s, item)
	}
	sort.Strings(s)
	for _, item := range s {
		if g.vertices[item].color == "black" {
			t = g.dfs(g.vertices[item], t)
		}
	}
	v.color = "blue"
	v.finishTime = t
	t += 1
	return t
}

func (g *Graph) PrintDfs(w io.Writer) {
	var s []int
	for index, _ := range g.discoveryTimes {
		s = append(s, index)
	}
	sort.Ints(s)
	for _, k := range s {
		fmt.Fprintf(w, "%s", g.discoveryTimes[k])
	}
}

func main() {
	g := Graph{}
	v_list := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for _, item := range v_list {
		g.AddVertex(item)
	}
	e_list := map[string]int{"AB": 1, "AE": 1, "BF": 1, "FI": 1, "FG": 1, "GC": 1, "GF": 1, "GJ": 1, "IH": 1, "IF": 1, "HE": 1, "HD": 1, "DE": 1, "DH": 1, "ED": 1, "EH": 1, "EA": 1}
	for k, v := range e_list {
		g.AddEdge(string(k[0]), string(k[1]), v)
	}
	g.Print(os.Stdout)
	g.dfs(g.vertices["A"], 1)
	g.PrintDfs(os.Stdout)
}

