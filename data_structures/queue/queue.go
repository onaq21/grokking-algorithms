package main

import (
	"fmt"
	"errors"
)

type Queue struct {
	head *Node
	tail *Node
	length int
}

type Node struct {
	data int
	next *Node
}

func (queue *Queue) enqueue(data int) {
	newNode := &Node{data, nil}
	if queue.isEmpty() {
		queue.head = newNode
		queue.tail = newNode
	} else {
		queue.tail.next = newNode
	}
	queue.tail = newNode
	queue.length++
}

func (queue *Queue) dequeue() (int, error) {
	if queue.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	if queue.length == 1 { queue.tail = nil }
	data := queue.head.data
	queue.head = queue.head.next
	queue.length--
	return data, nil
}

func (queue *Queue) isEmpty() bool {
	return queue.length == 0
}

func (queue *Queue) peek() (int, error) {
	if queue.isEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return queue.head.data, nil
}

func main() {
	q := &Queue{}

	for _, v := range []int{10, 20, 30} {
		q.enqueue(v)
	}

	for i := 0; i < 4; i++ {
		if el, err := q.dequeue(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(el)
		}
	}

	fmt.Println("Is empty?", q.isEmpty())
}

