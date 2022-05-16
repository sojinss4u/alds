package main

import (
	"bytes"
	"testing"
)

func TestTopView(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 2, 6, 8}
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestTopViewSuccess", func(t *testing.T) {
		expect := "13578"
		var b bytes.Buffer
		if tr.TopView(&b); b.String() != expect {
			t.Errorf("Expected %s, Got %s", expect, b.String())
		}
	})
}

func TestLeftView(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		5       	 level = 0
	//  	 3  	 7		 level = 1
	//	 1       		8    level = 2
	// 0   2         6     9 level = 3
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestLeftViewSuccess", func(t *testing.T) {
		expect := "5310"
		var b bytes.Buffer
		if tr.LeftView(&b); b.String() != expect {
			t.Errorf("Expecrted %s, Got %s", expect, b.String())
		}
	})
}

func TestRightView(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     	     5       	 level = 0
	//  	 3  	 7		 level = 1
	//    1      6	    8    level = 2
	// 0     2              9 level = 3
	
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestRightViewSuccess", func(t *testing.T) {
		var b bytes.Buffer
		expect := "5789"
		if tr.RightView(&b); b.String() != expect {
			t.Errorf("Expected %s, Got %s", expect, b.String())
		}
	})
}

func TestVerticalLevelOrderTraversalLeftToRight(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		5       	   level = 0
	//  	3  	    7		   level = 1
	//	 1      6	   8       level = 2
	// 0     2             9   level = 3
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestVerticalLevelOrderTraversalLeftToRightSuccess", func(t *testing.T) {
		var b bytes.Buffer
		expect := "013256789"
		if tr.VerticalLevelOrderTraversalLeftToRight(&b); b.String() != expect {
			t.Errorf("Expected %s, Got %s", expect, b.String())
		}
	})
}

func TestBottomView(t *testing.T) {
	tr := Tree{}
	i := []int{5, 3, 7, 1, 8, 0, 2, 6, 9}
	//     		5       	   level = 0
	//  	3  	    7		   level = 1
	//	 1      6	   8       level = 2
	// 0     2             9   level = 3
	for _, val := range i {
		tr.Insert(val)
	}
	t.Run("TestBottomViewSuccess", func(t *testing.T) {
		var b bytes.Buffer
		expect := "0126789"
		if tr.BottomView(&b); b.String() != expect {
			t.Errorf("Expect %s, Got %s", expect, b.String())
		}
	})
}
