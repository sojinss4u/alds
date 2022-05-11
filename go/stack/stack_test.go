package main

import (
	"testing"
)

var s = Stack{}

// Tests For IsEmpty() Function
func TestIsEmpty(t *testing.T) {
	// Reset stack data
	s.Reset()
	if got := s.IsEmpty(); got != true {
		t.Errorf("Expected true, Got %t", got)
	}
	s.Push("d1")
	if got := s.Length(); got != 1 {
		t.Errorf("Expected 1, got %d", got)
	}
}

// Tests For Push() Function

func TestPush(t *testing.T) {
	// Reset stack data
	s.Reset()
	s.Push("d1")
	s.Push("d2")
	if got := s.Length(); got != 2 {
		t.Errorf("Expected 2, Got %d", got)
	}
	if got := s.Top(); got != "d2" {
		t.Errorf("Expected d2, got %s", got)
	}
}

// Tests For Pop() Function

func TestPop(t *testing.T) {
	// Reset Stack Data
	s.Reset()
	if got := s.Pop(); got != nil {
		t.Errorf("Expected nil, got %v", got)
	}
	s.Push("d1")
	s.Push("d2")
	if got := s.Pop(); got != "d2" {
		t.Errorf("Expected d2, got %s", got)
	}
	if got := s.Length(); got != 1 {
		t.Errorf("Expected 1, got %d", got)
	}
}

// Tests for Top() function

func TestTop(t *testing.T) {
	s.Reset()
	if got := s.Top(); got != nil {
		t.Errorf("Expected nil, Got %v", got)
	}
	s.Push("d1")
	s.Push("d2")
	if got := s.Top(); got != "d2" {
		t.Errorf("Expected D2, Got %s", got)
	}
	s.Pop()
	if got := s.Top(); got != "d1" {
		t.Errorf("Expected d1, got %s", got)
	}
}
