package task

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// интерфейс structWithPresentation определяет методы для структур, имеющих строковое представление
type structWithPresentation interface {
	String() string
	GetType() string
}

// структура someStruct содержит одно поле num
type someStruct struct {
	num int
}

// метод String возвращает строковое представление someStruct
func (s someStruct) String() string {
	return fmt.Sprint(s.num)
}

// метод GetType возвращает тип структуры
func (s someStruct) GetType() string {
	return "someStruct"
}

// функция worker выполняет задачи из канала tasks и выводит результат в stdout
// воркеры работают до тех пор, пока открыт канал
func worker(id int, tasks <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.New(rand.NewSource(time.Now().Unix()))

	fmt.Println("start worker ", id)

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

// функция dataProcessor запускает воркеров для обработки данных из dataSlice
func dataProcessor(ctx context.Context, dataSlice []interface{}, numWorkers int) {
	var wg sync.WaitGroup
	dataChan := make(chan interface{}, len(dataSlice))
	ctx1, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	go func() {
		defer close(dataChan)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				for _, value := range dataSlice {
					select {
					case dataChan <- value:
					case <-ctx1.Done():
						return
					case <-ctx.Done():
						return
					}
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		fmt.Println("\nTimeout reached, shutting down...")
	case <-stopChan:
		fmt.Println("\nReceived interrupt signal, shutting down...")
		cancel()

	}

	wg.Wait() // ожидание завершения всех воркеров
}

func TaskN4() {
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	dataSlice := []someStruct{{1}, {2}, {3}, {4}}

	interfaceSlice := make([]interface{}, len(dataSlice))
	for i, v := range dataSlice {
		interfaceSlice[i] = v
	}

	dataProcessor(ctx, interfaceSlice, 2)
}
