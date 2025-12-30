package main

import (
	"fmt"
	"errors"
)

type List struct {
	head *Node
	length int
}

type Node struct {
	data int
	next *Node
}

func (list *List) reverse() {
	if list.length < 2 { return }

	var prev *Node = nil
	current := list.head
	var next *Node = nil

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	list.head = prev
}

func (list *List) insertAtHead(data int) {
	list.head = &Node{data, list.head}
	list.length++
}

func (list *List) insertAtTail(data int) {
	newNode := &Node{data, nil}
	if list.isEmpty() {
		list.head = newNode
	} else {
		for node := list.head; node != nil; node = node.next {
			if node.next == nil {
				node.next = newNode
				break
			}
		}	
	}
	list.length++
}

func (list *List) insertAtIndex(data, index int) error {
	if index < 0 || index > list.length { return errors.New("Index out of bounds") }
	if index == 0 {
		list.insertAtHead(data)
		return nil
	} else {
		node := list.head
		for i := 0; i < index - 1; i++ { node = node.next }
		newNode := &Node{data, node.next}
		node.next = newNode
		list.length++
		return nil
	}
}

func (list *List) deleteAtHead() (int , error) {
	if list.isEmpty() { return 0, errors.New("List is empty")}
	data := list.head.data
	list.head = list.head.next
	list.length--
	return data, nil
}

func (list *List) deleteAtTail() (int, error) {
	if list.isEmpty() { return 0, errors.New("List is empty")}
	data:= 0
	if list.length == 1 {
		data = list.head.data
		list.head = nil
	} else {
		for node := list.head; node.next != nil; node = node.next {
			if node.next.next == nil {
				data = node.next.data
				node.next = nil
				break
			}
		}
	}
	list.length--
	return data, nil
}

func (list *List) deleteAtIndex(index int) (int, error) {
	if index < 0 || index >= list.length { return 0, errors.New("Index out of bounds")}
	if index == 0 {
		return list.deleteAtHead()
	} else {
		node := list.head
		for i := 0; i < index - 1; i++ { node = node.next }
		data := node.next.data
		node.next = node.next.next
		list.length--
		return data, nil
	}
}

func (list *List) peek() (int, error) {
	if list.isEmpty() {
		return 0, errors.New("List is empty")
	}
	return list.head.data, nil
}

func (list *List) Size() int {
	return list.length
}

func (list *List) isEmpty() bool {
	return list.length == 0
}

func main() {
	list := &List{}
	list.insertAtHead(10)
	list.insertAtTail(20)
	list.insertAtTail(30)

	for node := list.head; node != nil; node = node.next {
		fmt.Print(node.data, " ")
	}
	fmt.Println()

	list.reverse()
	for node := list.head; node != nil; node = node.next {
		fmt.Print(node.data, " ")
	}
	fmt.Println()	//10 20 30
								//30 20 10
}