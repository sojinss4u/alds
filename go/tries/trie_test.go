package main

import (
	"testing"
)

func TestPrefixCount(t *testing.T) {
	root := Node{}
	w_list := []string{"damp", "dark", "data", "drake", "draw", "drew", "dried", "drunk", "drew"}
	for _, w := range w_list {
		InsertWord(w, &root)
	}
	t.Run("TestPrefixCountShortPrefixSuccess", func(t *testing.T) {
		expect := 9
		if got := PrefixCount("d", &root); got != expect {
			t.Errorf("Expected %d, Got %d", expect, got)
		}
	})
	t.Run("TestPrefixCountLongPrefixSuccess", func(t *testing.T) {
		expect := 1
		if got := PrefixCount("drake", &root); got != expect {
			t.Errorf("Expected %d, got %d", expect, got)
		}
	})
	t.Run("TestPrefixCountNonExistingPrefix", func(t *testing.T) {
		expect := 0
		if got := PrefixCount("doom", &root); got != expect {
			t.Errorf("Expected %d, Got %d", expect, got)
		}
	})
}

