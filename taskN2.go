package main

import (
	"fmt"
	"sync"
)

func routineCalc(slice []int) {
	var wg sync.WaitGroup

	for _, num := range slice {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, ": ", num*num)
		}(num)
	}
	wg.Wait()
}

func taskN2() {
	routineCalc([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
