package main

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	/*A : [B C]
	B : [A C]
	C : [A B]*/
	n1 := &Node{value: "A"}
	n2 := &Node{value: "B"}
	n3 := &Node{value: "C"}
	g := Graph{}
	g.AddVertex(n1)
	g.AddVertex(n2)
	g.AddVertex(n3)
	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n3, 1)
	g.AddEdge(n2, n1, 1)
	g.AddEdge(n2, n3, 1)
	g.AddEdge(n3, n1, 1)
	g.AddEdge(n3, n2, 1)
	expect := "A : [0 1 1]\nB : [1 0 1]\nC : [1 1 0]\n"
	b := bytes.Buffer{}
	if g.Print(&b); b.String() != expect {
		t.Errorf("Expect %s, Got %s", expect, b.String())
	}
}

