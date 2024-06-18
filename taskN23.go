package main

import "fmt"

func taskN23() {
	// способ используя алгоритм
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Before removing the i-th element: ", sl)
	sl = remove1(sl, 4)
	fmt.Println("After removing the i-th element: ", sl)

	// способ используя возможности стандартных функций
	sl = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Before removing the i-th element: ", sl)
	sl = remove2(sl, 4)
	fmt.Println("After removing the i-th element: ", sl)
}

func remove1(arr []int, i int) []int {
	if 0 > i || i >= len(arr) {
		return nil
	}

	// Сдвигаем элементы влево, начиная с i
	for ; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	newArr := make([]int, len(arr)-1) // Создаем новый слайс с уменьшенной длиной и вместимостью
	copy(newArr, arr[:len(arr)-1])

	return newArr
}

func remove2(arr []int, i int) []int {
	if 0 > i || i >= len(arr) {
		return nil
	}

	newArr := append(arr[:i], arr[i+1:]...)
	return newArr
}
