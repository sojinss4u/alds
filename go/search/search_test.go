package main

import "testing"

func TestInorderLinerarSearch(t *testing.T) {
	t.Run("TestInorderLinerarSearchIntegersSuccess", func(t *testing.T) {
		ar := []interface{}{3, 1, 7, 2, 6}
		i := UnorderedLinerarSearch(2, ar)
		expect := 3
		if expect != i {
			t.Errorf("Expected %v, Got %v", expect, i)
		}
	})
	t.Run("TestInorderLinerarSearchStringSuccess", func(t *testing.T) {
		ar := []interface{}{"a", "c", "b", "g", "h"}
		i := UnorderedLinerarSearch("b", ar)
		expect := 2
		if expect != i {
			t.Errorf("Expect %v, Got %v", expect, i)
		}
	})
}

