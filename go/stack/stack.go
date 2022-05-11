package main

import (
	"log"
	"os"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

// logger is initialized here
func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// A new type for struct is defined here with empty interface, so that it can store any data type as value
type Stack []interface{}

// Method to check length of the stack
func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	} else {
		return false
	}
}

// Method to push data to stack

func (s *Stack) Push(data interface{}) {
	*s = append(*s, data)
}

// Method to remove data from the top of the stack & retun the same back

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() == true {
		errorLogger.Print("Stack is empty. Pop() is not possible")
		return nil
	} else {
		topIndex := len(*s) - 1
		topElement := (*s)[topIndex]
		*s = (*s)[:topIndex]
		return topElement
	}
}

// Method to print stack data

func (s *Stack) Print() {
	infoLogger.Print(*s)
}

// Method to return top of the stack

func (s *Stack) Top() interface{} {
	if s.IsEmpty() == true {
		infoLogger.Print("Stack is empty. Can't return top of the stack")
		return nil
	} else {
		topIndex := len(*s) - 1
		return (*s)[topIndex]
	}
}

// Reset stack to empty stack
func (s *Stack) Reset() {
	*s = nil
}

func (s *Stack) Length() int {
	return len(*s)
}

func main() {
	s := &Stack{}
	r := s.IsEmpty()
	infoLogger.Println(r)
	s.Push("Soji")
	s.Push("Soniya")
	s.Push("Jesbin")
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Print()
	s.Pop()
	s.Top()
	s.Push("Raeyan")
	s.Push("Seira")
	top := s.Top()
	infoLogger.Print(top)
	s.Reset()
	s.Print()
}
