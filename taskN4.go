/*
 *	Реализовать постоянную запись данных в канал (главный поток).
 *	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
 *	 Необходима возможность выбора количества воркеров при старте.
 *
 * 	Программа должна завершаться по нажатию Ctrl+C.
 *	Выбрать и обосновать способ завершения работы всех воркеров.
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type structWithPresentation interface {
	String() string
	GetType() string
}

type someStruct struct {
	num int
}

func (s someStruct) String() string {
	return fmt.Sprint(s.num)
}

func (s someStruct) GetType() string {
	return "someStruct"
}

func worker(id int, tasks <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.New(rand.NewSource(time.Now().Unix()))

	println("start worker ", id)

	for task := range tasks {
		randomNumber := rand.Intn(2) + 1
		time.Sleep(time.Duration(randomNumber) * time.Second)

		switch v := task.(type) {
		case int:
			fmt.Printf("Worker %d received int: %d\n", id, v)
		case string:
			fmt.Printf("Worker %d received string: %s\n", id, v)
		case structWithPresentation:
			fmt.Printf("Worker %d received %s: %s\n", id, v.GetType(), v.String())
		default:
			fmt.Printf("Worker %d received unknown type without presentation\n", id)
		}
	}
}

func dataProcessor(dataSlice []interface{}, numWorkers int) {
	var wg sync.WaitGroup
	dataChan := make(chan interface{}, len(dataSlice))

	println("start data processor")

	for i := 0; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	go func() {
		for {
			for _, value := range dataSlice {
				dataChan <- value
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	<-stopChan // Ожидание сигнала завершения программы
	fmt.Println("\nReceived interrupt signal, shutting down...")

	close(dataChan) // Закрытие канала
	wg.Wait()       // Ожидание завершения всех воркеров

	println("end data processor")
}

func taskN4() {
	dataSlice := []someStruct{{1}, {2}, {3}, {4}}

	interfaceSlice := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		interfaceSlice[i] = v
	}

	dataProcessor(interfaceSlice, 2)
}
