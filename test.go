package main

import (
	"fmt"
	"sync"
)

func test() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i // Конкурентная запись в map
		}(i)
	}

	wg.Wait()
	fmt.Println("Map:", m)
}
