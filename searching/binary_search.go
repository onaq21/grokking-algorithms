package main

import "fmt"

func binarySearch(list []int, el int) int {
	left, right := 0, len(list) - 1

	for left <= right {
		midle := (left + right) / 2

		if list[midle] == el {
			return midle
		} else if list[midle] < el {
			left = midle + 1
		} else {
			right = midle - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{-2, 1, 4, 6, 8, 9, 11, 123}, 9)) // 5
}
