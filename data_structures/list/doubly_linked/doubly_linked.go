package main

import (
	"fmt"
	"errors"
)

type List struct {
	head *Node
	tail *Node
	length int
}
type Node struct {
	data int
	prev *Node
	next *Node
}

func (list *List) reverse() {
	if list.length < 2 { return }

	var prev *Node = nil
	current := list.head
	var next *Node = nil

	list.tail = current
	for current != nil {
		next = current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	list.head = prev
} 

func (list *List) insertAtHead(data int) {
	if list.isEmpty() {
		newNode := &Node{data, nil, nil}
		list.head = newNode
		list.tail = newNode
	} else {
		newNode := &Node{data, list.head, nil}
		list.head.prev = newNode
		list.head = newNode
	}
	list.length++
}

func (list *List) insertAtTail(data int) {
	if list.isEmpty() {
		newNode := &Node{data, nil, nil}
		list.head = newNode
		list.tail = newNode
	} else {
		newNode := &Node{data, list.tail, nil}
		list.tail.next = newNode
		list.tail = newNode
	}
	list.length++
}

func (list *List) deleteAtHead() (int, error) {
	if list.isEmpty() { return 0, errors.New("List is empty")}
	data := list.head.data
	if list.length == 1{
		list.head = nil
		list.tail = nil
	} else {
		list.head.next.prev = nil
		list.head = list.head.next
	}
	list.length--
	return data, nil
}

func (list *List) deleteAtTail() (int, error) {
	if list.isEmpty() { return 0, errors.New("List is empty")}
	data := list.tail.data
	if list.length == 1 {
		list.head = nil
		list.tail = nil
	} else {
		list.tail = list.tail.prev
		list.tail.next = nil
	}
	list.length--
	return data, nil
}

func (list *List) isEmpty() bool {
	return list.length == 0
}
func (list *List) size() int {
	return list.length
}
func (list *List) peek() (int, error) {
	if list.isEmpty() { return 0, errors.New("List is empty")}
	return list.head.data, nil
}

func main() {
	list := &List{}

	list.insertAtHead(10)
	list.insertAtTail(20)
	list.insertAtTail(30)

	for node := list.head; node != nil; node = node.next {
		fmt.Print(node.data, " ")
	}
	fmt.Println() // 10 20 30

	list.reverse()

	for node := list.head; node != nil; node = node.next {
		fmt.Print(node.data, " ")
	}
	fmt.Println() // 30 20 10
}
