//Implement a doubly linked linked list
package main

import (
	"fmt"
)

// Node struct
type node struct {
	data           string
	next, previous *node
}

// LinkedList struct
type linkedList struct {
	head *node
}

func createNode(data string) *node {
	n := node{
		data: data,
	}
	return &n
}

// Methods for linkedList

// appedNode() method
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

// printList() method
func (l *linkedList) printList() {
	fmt.Println("=============")
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

// insertHead() method

func (l *linkedList) insertHead(data string) {
	n := createNode(data)
	if l.head == nil {
		l.head = n
	} else {
		currentNode := l.head
		l.head = n
		l.head.next = currentNode
		currentNode.previous = n
	}
}

// listLenght() method

func (l *linkedList) listLength() int {
	if l.head == nil {
		return 0
	} else {
		count := 1
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
			count += 1
		}
		return count
	}
}

// insertAt() method

func (l *linkedList) insertAt(position int, data string) {
	if position > l.listLength() || position < 1 {
		fmt.Println("Invalid position")
	} else if position == 1 {
		l.insertHead(data)
	} else {
		n := createNode(data)
		currentNode := l.head
		count := 1
		for count < position {
			currentNode = currentNode.next
			count += 1
		}
		currentNode.previous.next = n
		n.previous = currentNode.previous
		n.next = currentNode
		currentNode.previous = n
	}
}

// deleteHead()

func (l *linkedList) deleteHead() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else if l.head.next == nil {
		l.head = nil
	} else {
		currentHead := l.head
		l.head = currentHead.next
		l.head.previous = nil
	}
}

// deleteTail()

func (l *linkedList) deleteTail() {
	if l.head == nil {
		fmt.Println("List is empty")
	} else if l.head.next == nil {
		l.deleteHead()
	} else {
		currentHead := l.head
		for currentHead.next != nil {
			currentHead = currentHead.next
		}
		currentHead.previous.next = nil
		currentHead.previous = nil
	}
}

// deleteAt() method

func (l *linkedList) deleteAt(position int) {
	if l.head == nil {
		fmt.Println("List is empty")
	} else if position == 1 {
		l.deleteHead()
	} else {
		currentNode := l.head
		count := 1
		for count < position {
			currentNode = currentNode.next
		}
		currentNode.previous.next = currentNode.next
		currentNode.next.previous = currentNode.previous
	}
}

func main() {
	l := linkedList{}
	l.appendNode("Soji")
	l.appendNode("Soniya")
	l.printList()
	l.insertHead("Antony")
	l.printList()
	l.insertHead("Treesa")
	l.printList()
	fmt.Println(l.listLength())
	l.insertAt(2, "Raeyan")
	l.printList()
	l.deleteHead()
	l.printList()
	l.deleteHead()
	l.printList()
	l.deleteTail()
	l.printList()
	l.deleteTail()
	l.printList()
	l.deleteAt(1)
	l.printList()
}

