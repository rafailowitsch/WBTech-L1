package main

import "fmt"

func median(arr []int, left, mid, right int) {
	switch {
	case mid < left && left < right:
		swap(arr, left, right)
	case left < mid && mid < right:
		swap(arr, mid, right)
	default:
		return
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, mid, right := 0, len(arr)/2, len(arr)-1
	median(arr, left, mid, right)

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	swap(arr, right, left)

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func taskN15() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)
	quickSort(arr)
	fmt.Println("Отсортированный массив:", arr)

	// arr = []int{6, 35, 87, 67, 9, 2, 5, 56, 2, 57, 35, 83, 27, 13, 8, 3, 6, 9, 17, 10, 56}
	// quicksort(arr)
	// fmt.Println("Отсортированный массив:", arr)
}
