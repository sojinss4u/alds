package main

import (
	"bytes"
	"testing"
	"fmt"
	"math"
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

func TestLevelOrderTraversalLeftToRight(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1}
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestLevelOrderTraversalLeftToRightSuccess", func(t *testing.T) {
		expect := "5371"
		var b bytes.Buffer
		if tr.LevelOrderTraversalLeftToRight(&b); b.String() != expect {

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

func TestLevelOrderTraversalRightToLeft(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		5       	 level = 0
	//  	 3  	 7		 level = 1
	//	 1         6	8    level = 2
	// 0   2              9 level = 3
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestLevelOrderTraversalRightToLeftSuccess", func(t *testing.T) {
		var b bytes.Buffer
		expect := "573861920"
		if tr.LevelOrderTraversalRightToLeft(&b); b.String() != expect {
			t.Errorf("Expect %s, Got %s", expect, b.String())
		}
	})
}

func TestLCA(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		 5
	//  	 3  	 7
	//	  1      6	     8
	// 0     2                 9
	for _, val := range i {
		tr.Insert(val)
	}
	var b bytes.Buffer
	expect := "1"
	if tr.LCA(&b, 0, 2); b.String() != expect {
		t.Errorf("Expected %s, Got %s", expect, b.String())
	}
}


func TestPathBetween(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		 5
	//  	 3  	 7
	//	  1      6	     8
	// 0     2                 9
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestPathBetweenSuccess", func(t *testing.T) {
		var b bytes.Buffer
		expect := "6789"
		if tr.PathBetween(&b, 6, 9); b.String() != expect {
			t.Errorf("Expect %s, Got %s", expect, b.String())
		}
	})
}

func TestBst(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		 5
	//  	 3  	 7
	//	  1      6	     8
	// 0     2                 9
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestBSTSuccess", func(t *testing.T) {
		expect := true
		if got := tr.CheckBST(); got != expect {
			t.Errorf("Expected %t, Got %t", expect, got)
		}
	})
	// Generate Invalid BST Using Following Data
	//     		 5
	//  	 3  	 7
	//	  1      2
	tr = Tree{}
	tr.root = tr.CreateNode(5)
	tr.root.left = tr.CreateNode(3)
	tr.root.right = tr.CreateNode(7)
	tr.root.left.left = tr.CreateNode(1)
	tr.root.left.right = tr.CreateNode(2)
	t.Run("TestBSTFailure", func(t *testing.T) {
		expect := false
		if got := tr.CheckBST(); got != expect {
			t.Errorf("Expect %t, Got %t", expect, got)
		}
	})
}


func TestGreatestElementLessThanK(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		 5
	//  	 3  	 7
	//	  1      6	     8
	// 0     2                 9
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestGreatestElementLessThanKSuccess", func(t *testing.T) {
		expect := 2
		if got := tr.GreatestElementLessThanK(3); got != expect {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})

	tr = Tree{}
	//    5
	// Check an i/p which doesn't exist in tree
	tr.root = tr.CreateNode(5)
	t.Run("TestGreatestElementLessThanKSingleNodeTree", func(t *testing.T) {
		expect := math.MinInt
		if got := tr.GreatestElementLessThanK(1); expect != got {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})

}

func TestLeastElementGreaterThanK(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		 5
	//  	 3  	 7
	//	  1      6	     8
	// 0     2                 9
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestLeastElementGreaterThanKSuccess", func(t *testing.T) {
		expect := 7
		if got := tr.LeastElementGreaterThanK(6); got != expect {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})
	tr = Tree{}
	tr.root = tr.CreateNode(5)
	t.Run("TestLeastElementGreaterThanKSuccessSingleNodeTree", func(t *testing.T) {
		expect := math.MinInt
		if got := tr.LeastElementGreaterThanK(6); got != expect {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})
}

func TestDelete(t *testing.T) {
	tr := Tree{}
	//           5
	//        3     8
	//      2   4 6   9
	//              7
	//
	n := []int{5, 3, 8, 2, 4, 6, 9, 7}
	for _, val := range n {
		tr.Insert(val)
	}
	t.Run("TestDeleteNoChildren", func(t *testing.T) {
		expect := "3456789"
		var b bytes.Buffer
		tr.Delete(tr.root, 2)
		if tr.Print("io", &b); b.String() != expect {
			t.Errorf("Expect %s, Got %s", expect, b.String())
		}
	})
	tr = Tree{}
	//           5
	//        3     8
	//      2   4 6   9
	//              7
	//
	n = []int{5, 3, 8, 2, 4, 6, 9, 7}
	for _, val := range n {
		tr.Insert(val)
	}
	t.Run("TestDeleteSingleChild", func(t *testing.T) {
		expect := "2345789"
		var b bytes.Buffer
		tr.Delete(tr.root, 6)
		if tr.Print("io", &b); b.String() != expect {
			t.Errorf("Expect %s, Got %s", expect, b.String())
		}
	})
	tr = Tree{}
	//           5
	//        3     8
	//      2   4 6   9
	//              7
	//
	n = []int{5, 3, 8, 2, 4, 6, 9, 7}
	for _, val := range n {
		tr.Insert(val)
	}
	t.Run("TestDeleteDoubleChild", func(t *testing.T) {
		expect := "2345679"
		var b bytes.Buffer
		tr.Delete(tr.root, 8)
		if tr.Print("io", &b); b.String() != expect {
			t.Errorf("Expect %s, Got %s", expect, b.String())
		}
	})
}

