package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func taskN9() {
	var wg sync.WaitGroup

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	wg.Add(1)
	go func(arr [20]int) {
		defer wg.Done()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, num := range arr {
			ch1 <- num
			// println("processor 1 received:", num)
			randNum := r.Intn(1500) + 500
			time.Sleep(time.Duration(randNum) * time.Millisecond)
		}
		close(ch1)
	}(arr)

	wg.Add(1)
	go func(ch1, ch2 chan int) {
		defer wg.Done()
		rand.New(rand.NewSource(time.Now().Unix()))
		for num := range ch1 {
			ch2 <- (num * 2)
			// println("processor 2 received:", num)
			randNum := rand.Intn(180) + 60
			time.Sleep(time.Duration(randNum) * time.Millisecond)
		}
		close(ch2)
	}(ch1, ch2)

	wg.Add(1)
	go func(ch chan int) {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("PROCESSOR 3: %d received\n", num)
		}
	}(ch2)

	wg.Wait()
}
