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

func TestHeight(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1}
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestHeightSuccess", func(t *testing.T) {
		expect := 2.0
		if got := tr.Height(tr.root); got != expect {
			t.Errorf("Expected %f, Got %f", expect, got)
		}
	})
}

func TestLevelOrderTraversal(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1}
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestLevelOrderTraversalSuccess", func(t *testing.T) {
		expect := "5371"
		var b bytes.Buffer
		if tr.LevelOrderTraversal(&b); b.String() != expect {

		}
	})
}

func TestSearch(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 2, 6, 8}
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestSearchTrue", func(t *testing.T) {
		expect := true
		if got := tr.Search(2, tr.root); got != expect {
			t.Errorf("Expected %t, Got %t", expect, got)
		}
	})
	t.Run("TestSearchFalse", func(t *testing.T) {
		expect := false
		if got := tr.Search(9, tr.root); got != expect {
			t.Errorf("Expected %t, Got %t", expect, got)
		}
	})
}

func TestFindPath(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 2, 6, 8}
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestFindPathSuccess", func(t *testing.T) {
		expect := "2135"
		var b bytes.Buffer
		if tr.FindPath(&b, tr.root, 2); b.String() != expect {
			t.Errorf("Expect %s, Got %s", expect, b.String())
		}
	})
}

func TestLevelOrderTraversalWithNewLine(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8}
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestLevelOrderTraversalWithNewLineSuccess", func(t *testing.T) {
		var b bytes.Buffer
		expect := fmt.Sprintf("5\n37\n18")
		if tr.LevelOrderTraversalWithNewLine(&b); b.String() != expect {
			t.Errorf("Expected %s, Got %s", expect, b.String())
		}
	})

}

