package main

import "fmt"

func selectionSort(list []int) {	// o(n^2)
	for i := 0; i < len(list); i++ {
		min, minIndex := list[i], i

		for j := i; j < len(list); j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}

		list[i], list[minIndex] = list[minIndex], list[i]
	}
}

func main() {
	list := []int{-2, 9, 10, 134, -210, 1423, 99, 10, -1, -2}
	selectionSort(list)
	fmt.Println(list)	// -210 -2 -2 -1 9 10 10 99 134 1423
}