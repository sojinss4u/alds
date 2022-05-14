package main

import (
	"bytes"
	"testing"
)

func TestCreateNode(t *testing.T) {
	tr := Tree{}
	t.Run("Success_Create_Node", func(t *testing.T) {
		expect := &Node{data: 1}
		if got := tr.CreateNode(1); *got != *expect {
			t.Errorf("Expected %v, Got %v", expect, got)
		}
	})
}

func TestInOrderTraversal(t *testing.T) {
	tr := Tree{}
	n := []int{5, 3, 7, 1}
	for _, val := range n {
		tr.Insert(val)
	}
	t.Run("Success_InOrder_Traversal", func(t *testing.T) {
		//tr.Print("io", os.Stdout)
		expect := "1357"
		var b bytes.Buffer
		if tr.Print("io", &b); b.String() != expect {
			t.Errorf("Expected %s, Got %s", expect, b.String())
		}

	})
}

func TestPreOrderTraversal(t *testing.T) {
	tr := Tree{}
	n := []int{5, 3, 7, 1}
	for _, val := range n {
		tr.Insert(val)
	}
	t.Run("TestPreOrderTraversalSuccess", func(t *testing.T) {
		expect := "5317"
		var b bytes.Buffer
		if tr.Print("pro", &b); b.String() != expect {
			t.Errorf("Expected %s, Got %s", expect, b.String())
		}
	})
}

func TestPostOrderTraversal(t *testing.T) {
	tr := Tree{}
	n := []int{5, 3, 7, 1}
	for _, val := range n {
		tr.Insert(val)
	}
	t.Run("TestPostOrderTraversalSuccess", func(t *testing.T) {
		expect := "1375"
		var b bytes.Buffer
		if tr.Print("poo", &b); b.String() != expect {
			t.Errorf("Expected %s, Got %s", expect, b.String())
		}
	})
}

func TestCount(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1}
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestCountSuccess", func(t *testing.T) {
		expect := 4
		if got := tr.Count(tr.root); got != expect {
			t.Errorf("Expected %d, Got %d", expect, got)
		}
	})
}
