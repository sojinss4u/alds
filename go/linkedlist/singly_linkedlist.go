package main

import (
	"fmt"
)

// Node struct

type node struct {
	data string
	next *node
}

// LinkedList struct

type linkedList struct {
	head *node
}

// Function for creating node

func createNode(data string) *node {
	n := node{
		data: data,
	}
	return &n
}

// Methods for linked list are implemented here

// appendNode() : Append a node at the end of the linkedlist

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

// insertHead() method

func (l *linkedList) insertHead(data string) {
	n := createNode(data)
	if l.head == nil {
		l.head = n
	} else {
		currentHead := l.head
		l.head = n
		n.next = currentHead
	}
}

// listLength() method

func (l *linkedList) listLength() int {
	if l.head == nil {
		return 0
	} else {
		currentNode := l.head
		count := 1
		for currentNode.next != nil {
			count += 1
			currentNode = currentNode.next
		}
		return count
	}
}

// insertAt() method

func (l *linkedList) insertAt(position int, data string) {
	if position > l.listLength() || position < 1 {
		fmt.Println("Invalid Positiion")
	} else if position == 1 {
		l.insertHead(data)
	} else {
		n := createNode(data)
		var previousNode *node
		currentNode := l.head
		count := 1
		for count < position {
			count += 1
			previousNode = currentNode
			currentNode = currentNode.next
		}
		previousNode.next = n
		n.next = currentNode
	}
}

// deleteTail() method

func (l *linkedList) deleteTail() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		currentNode := l.head
		var previousNode *node
		for currentNode.next != nil {
			previousNode = currentNode
			currentNode = currentNode.next
		}
		previousNode.next = nil
	}
}

// deleteHead() method

func (l *linkedList) deleteHead() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		currentHead := l.head
		newHead := l.head.next
		l.head = newHead
		currentHead.next = nil
	}
}

// deleteAt() : Delete node at position

func (l *linkedList) deleteAt(position int) {
	if position > l.listLength() || position < 1 {
		fmt.Println("Invalid position")
	} else if position == 1 {
		l.deleteHead()
	} else {
		currentNode := l.head
		var previousNode *node
		count := 1
		for count < position {
			count += 1
			previousNode = currentNode
			currentNode = currentNode.next
		}
		previousNode.next = currentNode.next
		currentNode.next = nil
	}
}

// printList() : Print a linkedlist

func (l *linkedList) printList() {
	fmt.Println("====================")
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
	l.appendNode("Soji")
	l.appendNode("Soniya")
	l.appendNode("Jesbin")
	l.printList()
	l.insertHead("Treesa")
	l.printList()
	l.insertHead("Antony")
	l.printList()
	l.insertAt(3, "Raeyan")
	l.printList()
	l.insertAt(6, "Seira")
	l.printList()
	l.deleteTail()
	l.printList()
	l.deleteTail()
	l.printList()
	l.deleteHead()
	l.printList()
	l.deleteHead()
	l.printList()
	l.deleteAt(2)
	l.printList()
	l.deleteAt(2)
	l.printList()
}
