package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// функция taskN9 демонстрирует создание конвейера для обработки чисел
func taskN9() {
	var wg sync.WaitGroup

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20} // функция будет работать до тех пор, пока не получит квадрат каждого числа

	// горутина для записи чисел в первый канал
	wg.Add(1)
	go func(arr [20]int) {
		defer wg.Done()
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for _, num := range arr {
			ch1 <- num
			fmt.Printf("processor 1 sent: %d\n", num)
			randNum := r.Intn(1500) + 500
			time.Sleep(time.Duration(randNum) * time.Millisecond)
		}
		close(ch1)
	}(arr)

	// горутина для чтения из первого канала, умножения на 2 и записи во второй канал
	wg.Add(1)
	go func(ch1, ch2 chan int) {
		defer wg.Done()
		rand.New(rand.NewSource(time.Now().Unix()))
		for num := range ch1 {
			ch2 <- (num * 2)
			fmt.Printf("processor 2 sent: %d\n", num*2)
			randNum := rand.Intn(180) + 60
			time.Sleep(time.Duration(randNum) * time.Millisecond)
		}
		close(ch2)
	}(ch1, ch2)

	// горутина для чтения из второго канала и вывода в stdout
	wg.Add(1)
	go func(ch chan int) {
		defer wg.Done()
		for num := range ch {
			fmt.Printf("processor 3 received: %d\n", num)
		}
	}(ch2)

	wg.Wait()
}
