package main

import (
	"testing"
)

var q = Queue{}

func TestIsEmpty(t *testing.T) {
	q.EmptyQueue()
	t.Run("TestIsEmptyTrue", func(t *testing.T) {
		if got := q.IsEmpty(); got != true {
			t.Errorf("Expected true, Got %t", got)
		}
	})
	q.Enqueue("d1")
	t.Run("TestIsEmptyFalse", func(t *testing.T) {
		if got := q.IsEmpty(); got != false {
			t.Errorf("Expected false, got %t", got)
		}
	})
}

func TestEnqueue(t *testing.T) {
	q.EmptyQueue()
	q.Enqueue("d1")
	q.Enqueue("d2")
	t.Run("TestEnqueueSuccessLength", func(t *testing.T) {
		if got := q.QueueLength(); got != 2 {
			t.Errorf("Expected 2, got %d", got)
		}
	})
	q.Enqueue("d3")
	t.Run("TestEnqueueSuccessItem", func(t *testing.T) {
		if got := q[0]; got != "d1" {
			t.Errorf("Expected d1, got %v", got)
		}
	})
}

func TestDequeue(t *testing.T) {
	q.EmptyQueue()
	q.Enqueue("d1")
	q.Enqueue("d2")
	q.Enqueue("d3")
	q.Dequeue()
	t.Run("TestDequeueLength", func(t *testing.T) {
		if got := q.QueueLength(); got != 2 {
			t.Errorf("Expected 2, Got %d", got)
		}
	})
	q.Dequeue()
	t.Run("TestDequeueItem", func(t *testing.T) {
		if got := q[0]; got != "d3" {
			t.Errorf("Expected d3, got %v", got)
		}
	})
	q.Dequeue()
	t.Run("TestDequeueEmptyQueue", func(t *testing.T) {
		if got := q.Dequeue(); got != nil {
			t.Errorf("Expected nil, got %s", got)
		}
	})
}

