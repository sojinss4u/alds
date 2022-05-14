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

