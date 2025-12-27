package main

import "fmt"

func demoSelectionSort() {
	list := []int{-2, 9, 10, 134, -210, 1423, 99, 10, -1, -2}
	selectionSort(list)
	fmt.Println(list)
}

func demoQuickSort() {
	list := []int{57, -23, 8, 104, 0, -76, 42, 19, -5, 88}
	quickSort(list)
	fmt.Println(list)
}

func main() {
	demoSelectionSort()	 // -210 -2 -2 -1 9 10 10 99 134 1423
	demoQuickSort() //	-76 -23 -5 0 8 19 42 57 88 104
}
