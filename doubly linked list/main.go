package main

import (
	"fmt"
)

type Node struct {
	prev *Node
	data int
	next *Node
}

type DoublyLinkedLIst struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedLIst) insertAtBeginning(data int) {
	newNode := Node{nil, data, nil}

	if dll.head == nil {
		dll.head = &newNode
		dll.tail = &newNode
	}

	dll.head.prev = &newNode
	newNode.next = dll.head
	dll.head = &newNode

}

func (dll *DoublyLinkedLIst) insertAtEnd(data int){
	newNode := Node{nil,data,nil}

	if dll.head == nil {
		dll.head = &newNode
		dll.tail = &newNode
	}

	newNode.prev = dll.tail
	dll.tail.next = &newNode
	dll.tail = &newNode
}

func (dll *DoublyLinkedLIst) display(){
	current := dll.head

	if current == nil {
		fmt.Print("List is emptyyy")
		return
	}

	for current != nil {
		fmt.Printf("%d->",current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := DoublyLinkedLIst{}

	list.insertAtEnd(69)
	list.insertAtEnd(690)
	list.insertAtEnd(420)

	list.display()
}