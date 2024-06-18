package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		println(arr[mid])
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func taskN16() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 13
	fmt.Printf("Element %d has index %d in array", target, binarySearch(arr, target))
}
