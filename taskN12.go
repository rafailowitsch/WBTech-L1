package main

import "fmt"

func set(series []string) []string {
	setM := make(map[string]struct{})   // сначала работаем с мапой, так как мы имеем константное время O(1) в случае чтения элемента
	set := make([]string, 0, len(setM)) // итоговый слайс (заранее устанавливаем вместимость)

	for _, str := range series {
		setM[str] = struct{}{} // используем пустую структуру типа struct{}, так как она не занимает память и это эффективно для пространственной сложности
	}

	for key := range setM {
		set = append(set, key)
	}

	return set
}

func taskN12() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	result := set(s)
	fmt.Println("Set:", result)
}
