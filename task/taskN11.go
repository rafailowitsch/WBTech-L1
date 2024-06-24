package task

import "fmt"

// функция intersection находит пересечение двух множеств
func intersection(set1, set2 []int) []int {
	var inter []int
	elems := make(map[int]bool)

	// добавляем все элементы первого множества в мапу
	for _, v := range set1 {
		elems[v] = true
	}

	// проверяем элементы второго множества на наличие в мапе
	for _, v := range set2 {
		if elems[v] {
			inter = append(inter, v)
		}
	}

	return inter
}

// функция taskN11 демонстрирует пересечение двух множеств
func TaskN11() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersection(set1, set2)
	fmt.Println("Intersection:", result)
}
