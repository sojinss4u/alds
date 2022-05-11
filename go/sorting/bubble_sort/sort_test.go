package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	ar := []int{2, 3, 1, 4, 5}
	expect := []int{1, 2, 3, 4, 5}
	t.Run("Sorting_Success", func(t *testing.T) {
		if got := BubbleSort(ar); !reflect.DeepEqual(got, expect) {
			t.Errorf("Expected %v, Got %v", expect, got)
		}
	})
}

