package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	mu  sync.RWMutex
	num int
}

func NewCounter() *counter {
	return &counter{
		num: 0,
	}
}

func (c *counter) increment() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.num++
}

func (c *counter) get() int {
	defer c.mu.RUnlock()
	c.mu.RLock()
	return c.num
}

func taskN17() {
	var wg sync.WaitGroup

	counter := NewCounter()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			counter.increment()
			fmt.Println("Counter increment in first routine: ", counter.get())

			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(500 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			counter.increment()
			fmt.Println("Counter increment in second routine: ", counter.get())

			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
