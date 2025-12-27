package main

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