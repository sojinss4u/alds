package main

import (
	"testing"
	"math"
)

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

func TestFindGreatest(t *testing.T) {
	ar1 := []int{1, 3, 7, 10}
	t.Run("TestFindGreatestExistingElementSuccess", func(t *testing.T) {
		expect := 10
		if got := FindGreatest(10, ar1); got != expect {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})
	t.Run("TestFindGreatestNonExistingElementSuccess", func(t *testing.T) {
		expect := 7
		if got := FindGreatest(9, ar1); got != expect {
			t.Errorf("Expect %d, Got %d", got, expect)
		}
	})
	t.Run("TestFindGreatestReturnDefaultAnswer", func(t *testing.T) {
		expect := math.MinInt
		if got := FindGreatest(0, ar1); expect != got {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})
}


func TestFindFrequency(t *testing.T) {
	ar1 := []int{1, 3, 7, 7, 7, 10}
	t.Run("TestFindFrequencyExistingElement", func(t *testing.T) {
		expect := 3
		if got := FindFrequency(7, ar1); got != expect {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})
	t.Run("TestFindFrequencyNonExistingElement", func(t *testing.T) {
		expect := 0
		if got := FindFrequency(11, ar1); expect != got {
			t.Errorf("Expect %d,Got %d", expect, got)
		}
	})
}

func TestFindPeak(t *testing.T) {
	t.Run("TestFindPeakEdgeElementRightSuccess", func(t *testing.T) {
		ar := []int{4, 7, 1, 2}
		expect := 2
		if got := FindPeak(ar); got != expect {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})
	t.Run("TestFindPeakEdgeElementsLeftSuccess", func(t *testing.T) {
		ar := []int{7, 4, 5, 2, 1}
		expect := 7
		if got := FindPeak(ar); got != expect {
			t.Errorf("Exepct %d, Got %d", expect, got)
		}
	})
	t.Run("TestFindPeakMiddlePeak", func(t *testing.T) {
		ar := []int{3, 4, 9, 5, 1}
		expect := 9
		if got := FindPeak(ar); got != expect {
			t.Errorf("Expect %d, Got %d", expect, got)
		}
	})
}
