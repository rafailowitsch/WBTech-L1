package main

import "fmt"

func intersection(set1, set2 []int) []int {
	var inter []int
	elems := make(map[int]bool)

	for _, v := range set1 {
		elems[v] = true
	}

	for _, v := range set2 {
		if elems[v] {
			inter = append(inter, v)
		}
	}

	return inter
}

func taskN11() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := intersection(set1, set2)
	fmt.Println("Intersection:", result)
}
