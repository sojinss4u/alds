// Check if the linked list has a loop

package main

import (
	"fmt"
)

// Node struct
type node struct {
	data      string
	next      *node
	isVisited bool
}

// LinkedList struct
type linkedList struct {
	head *node
}

// function creates a node & return node pointer with given data
func createNode(data string) *node {
	n := node{
		data: data,
	}
	return &n
}

//Methods for linkedlist
func (l *linkedList) appendNode(data string) {
	n := createNode(data)
	if l.head == nil {
		l.head = n
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = n
	}
}

//Print linkedlist
func (l *linkedList) printList() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		currentNode := l.head
		fmt.Println(currentNode.data)
		for currentNode.next != nil {
			currentNode = currentNode.next
			fmt.Println(currentNode.data)
		}
	}
}

// detectCycle method

func (l *linkedList) detectCycle() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			if currentNode.isVisited == true {
				fmt.Println("Loop Detected")
				break
			} else {
				currentNode.isVisited = true
				currentNode = currentNode.next
			}
		}
	}
}

func main() {
	l := linkedList{}
	l.appendNode("Soji")
	l.appendNode("Soniya")
	l.printList()
	l.detectCycle()
	l.head.next.next = l.head
	l.detectCycle()
}

