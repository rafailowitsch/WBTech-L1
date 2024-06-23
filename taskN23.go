package main

import "fmt"

// функция taskN23 демонстрирует удаление элемента из слайса двумя способами
func taskN23() {
	// способ используя алгоритм
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Before removing the i-th element:", sl)
	sl = remove1(sl, 4)
	fmt.Println("After removing the i-th element:", sl)

	// способ используя возможности стандартных функций
	sl = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Before removing the i-th element:", sl)
	sl = remove2(sl, 4)
	fmt.Println("After removing the i-th element:", sl)
}

// функция remove1 удаляет i-ый элемент из слайса, сдвигая оставшиеся элементы
func remove1(arr []int, i int) []int {
	if 0 > i || i >= len(arr) {
		return nil
	}

	// сдвигаем элементы влево, начиная с i
	for ; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	// создаем новый слайс с уменьшенной длиной и вместимостью
	newArr := make([]int, len(arr)-1)
	copy(newArr, arr[:len(arr)-1])

	return newArr
}

// функция remove2 удаляет i-ый элемент из слайса с использованием append
func remove2(arr []int, i int) []int {
	if 0 > i || i >= len(arr) {
		return nil
	}

	return append(arr[:i], arr[i+1:]...)
}
