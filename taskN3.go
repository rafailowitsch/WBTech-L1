package main

import (
	"fmt"
	"sync"
)

func getSumSquares(slice []int) int {
	var wg sync.WaitGroup
	var sum int = 0

	for _, num := range slice {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sum += num * num
		}(num)
	}
	wg.Wait()

	return sum
}

func taskN3() {
	ss := getSumSquares([]int{2, 3, 4, 5})
	fmt.Println(ss)
}
