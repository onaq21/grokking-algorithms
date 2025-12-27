package main

func quickSort(list []int) {	// average case: O(nlog(n)), worst case: O(n^2)
	if len(list) < 2 {
		return
	}

	size := len(list)
	pivot := list[size / 2]
	left, right := 0, size - 1

	for left <= right {
		for list[left] < pivot { left++ }
		for list[right] > pivot { right-- }
		if left <= right {
			list[left], list[right] = list[right], list[left]
			left++
			right--
		}
	}

	quickSort(list[:right+1])
	quickSort(list[left:])
}