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

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ltime|log.Ldate|log.Lshortfile)
}

// Queue struct

type Queue []interface{}

// IsEmpty() method for Queue

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}

// Enqueue() method for Queue

func (q *Queue) Enqueue(data interface{}) {
	*q = append(*q, data)
}

// Print() Method for Queue

func (q *Queue) Print() {
	infoLogger.Print(*q)
}

// Dequeue() Method for Queue

func (q *Queue) Dequeue() interface{} {
	if e := q.IsEmpty(); e == true {
		infoLogger.Print("Queue is empty")
		return nil
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

// EmptyQueue() Method for Queue

func (q *Queue) EmptyQueue() {
	*q = nil
}

// QueueLength() Mewthod

func (q *Queue) QueueLength() int {
	return len(*q)
}

func main() {
	q := Queue{}
	infoLogger.Print(q.IsEmpty())
	q.Enqueue("Soji")
	q.Enqueue("Soniya")
	q.Print()
	q.Dequeue()
	q.Print()
	q.Dequeue()
	q.Print()
	q.Dequeue()
}

