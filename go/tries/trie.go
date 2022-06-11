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

// Node struct is defined here
// hm, contains the characters as key & theire corresponding node as value

type Node struct {
	hm    map[string]*Node
	count int
}

func (n *Node) InsertCharacter(c string) *Node {
	if _, ok := n.hm[string(c)]; !ok {
		// To avoid 'panic: assignment to entry in nil map' error
		if n.hm == nil {
			n.hm = make(map[string]*Node)
		}
		n.hm[string(c)] = &Node{}
	}
	n.hm[string(c)].count += 1
	return n.hm[string(c)]
}

func InsertWord(w string, r *Node) {
	for _, c := range w {
		r = r.InsertCharacter(string(c))
	}
}

func PrefixCount(pfx string, r *Node) int {
	var count int
	for _, c := range pfx {
		if _, ok := r.hm[string(c)]; ok {
			r = r.hm[string(c)]
			count = r.count
		} else {
			return 0
		}
	}
	return count
}

func main() {
	root := Node{}
	w_list := []string{"damp", "dark", "data", "drake", "draw", "drew", "dried", "drunk", "drew"}
	for _, w := range w_list {
		InsertWord(w, &root)
	}
	c := PrefixCount("damp", &root)
	fmt.Println(c)
}

