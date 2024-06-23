package main

import (
	"fmt"
	"sync"
	"time"
)

// структура counter с полями для мьютекса и значения счетчика
type counter struct {
	mu  sync.RWMutex
	num int
}

// функция NewCounter создает и возвращает новый экземпляр counter
func NewCounter() *counter {
	return &counter{
		num: 0,
	}
}

// метод increment увеличивает значение счетчика на 1
func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num++
}

// метод get возвращает текущее значение счетчика
func (c *counter) get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

// функция taskN18 демонстрирует конкурентное инкрементирование счетчика
func taskN18() {
	var wg sync.WaitGroup

	counter := NewCounter()

	// первая горутина
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			counter.increment()
			fmt.Println("Counter increment in first routine:", counter.get())
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(500 * time.Millisecond)

	// вторая горутина
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			counter.increment()
			fmt.Println("Counter increment in second routine:", counter.get())
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

	// вывод итогового значения счетчика
	fmt.Println("Final counter value:", counter.get())
}
