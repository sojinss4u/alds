// Detect loop in a doubly linked list

package main

import "fmt"

// Node struct is given below

type node struct {
	data           string
	next, previous *node
	isVisited      bool
}

// LinkedList struct is given below

type linkedList struct {
	head *node
}

// createNode() method

func createNode(data string) *node {
	n := node{
		data: data,
	}
	return &n
}

// appendNode() method

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
		n.previous = currentNode
	}
}

// detectLoop() method

func (l *linkedList) detectLoop() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			if currentNode.isVisited == true {
				fmt.Printf("Loop detected")
				break
			} else {
				currentNode.isVisited = true
			}
			currentNode = currentNode.next
		}
	}
}

// printList() method

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

func main() {
	l := linkedList{}
	l.appendNode("Antony")
	l.appendNode("Soji")
	l.appendNode("Soniya")
	l.printList()
	l.detectLoop()
	l.head.next.next = l.head
	l.detectLoop()
}

