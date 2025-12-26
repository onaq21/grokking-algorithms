package main

import (
	"fmt"
	"errors"
)

type Node struct {
	data int
	prev *Node
}

type Stack struct {
	top *Node
	length int
}

func push(stack *Stack, element int) {
	newNode := &Node{element, stack.top}
	stack.top = newNode
	stack.length++
}

func pop(stack *Stack) (int, error) {
	if isEmpty(stack) {
		return 0, errors.New("Stack is empty")
	}
	el := stack.top.data
	stack.top = stack.top.prev
	stack.length--
	return el, nil
}

func isEmpty(stack *Stack) bool {
	return stack.length == 0
}

func top(stack *Stack) (int, error) {
	if isEmpty(stack) {
		return 0, errors.New("Stack is empty")
	}
	return stack.top.data, nil
}

func size(stack *Stack) int {
	return stack.length
}

func clear(stack *Stack) {
	stack.top = nil
	stack.length = 0
}

func main() {
	stack := &Stack{}

	fmt.Println("Is stack empty:", isEmpty(stack))

	push(stack, 10)
	push(stack, 20)
	push(stack, 30)

	if value, err := top(stack); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top element:", value)
	}

	fmt.Println("Stack size:", size(stack))

	if value, err := pop(stack); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped:", value)
	}

	if value, err := pop(stack); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped:", value)
	}

	if value, err := top(stack); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top element:", value)
	}

	fmt.Println("Stack size:", size(stack))
	fmt.Println("Is stack empty:", isEmpty(stack))

	clear(stack)
	fmt.Println("After clearing, is stack empty:", isEmpty(stack))
	fmt.Println("Stack size:", size(stack))
}
