package task

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// функция taskN5 запускает два горутины: одна для отправки значений в канал, другая для их чтения и обработки
func TaskN5() {
	duration := time.Duration(30) * time.Second // задаем длительность работы программы
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	values := []int{1, 2, 3, 4, 5}
	ch := make(chan int, 2) // создаем канал с буфером размером 2

	// горутина для отправки значений в канал
	go func(values []int) {
		defer close(ch) // закрываем канал по завершении
		for {
			for _, val := range values {
				ch <- val
				fmt.Printf("Msg sent: %d\n", val)
			}
		}
	}(values)

	// горутина для чтения и обработки значений из канала
	go func(messages <-chan int) {
		rand.New(rand.NewSource(time.Now().Unix()))
		for msg := range messages {
			randomNumber := rand.Intn(180) + 60
			time.Sleep(time.Duration(randomNumber) * time.Millisecond) // симуляция обработки даннных из канала
			fmt.Printf("Msg [%d] processed\n", msg)
		}
	}(ch)

	for {
		select {
		case <-ctx.Done(): // завершение работы по истечению времени
			println("program interrupted")
			return
		default:
			time.Sleep(1 * time.Second) // задержка, чтобы не нагружать CPU
		}
	}
}
