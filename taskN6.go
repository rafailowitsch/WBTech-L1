// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func taskN6() {
	var wg sync.WaitGroup

	// завершение горутины используя context.WithTimeout
	{
		wg.Add(1)
		fmt.Println("1. terminating the goroutine using context.WithTimeout")
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		go func() {
			defer wg.Done()
			fmt.Println("goroutine 1 started")
			for {
				select {
				case <-ctx.Done():
					fmt.Println("goroutine 1 finished")
					return
				default:
					time.Sleep(time.Second)
				}
			}
		}()

		wg.Wait()
	}

	// завершение горутины используя context.WithCancel
	{
		wg.Add(1)
		fmt.Println("2. terminating the goroutine using context.WithCancel")
		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			defer wg.Done()
			fmt.Println("goroutine 2 started")
			for {
				select {
				case <-ctx.Done():
					fmt.Println("goroutine 2 finished")
					return
				default:
					time.Sleep(time.Second)
				}
			}
		}()

		time.Sleep(2 * time.Second)
		cancel()
		wg.Wait()
	}

	// завершение горутины используя канал
	{
		wg.Add(1)
		fmt.Println("3. terminating the goroutine using channel of bool type")
		done := make(chan bool)

		go func() {
			defer wg.Done()
			fmt.Println("goroutine 3 started")
			for {
				select {
				case <-done:
					fmt.Println("goroutine 3 finished")
					return
				default:
					time.Sleep(time.Second)
				}
			}
		}()

		time.Sleep(time.Second)
		done <- true
		wg.Wait()
	}

	// завершение горутины используя закрытый канал
	{
		wg.Add(1)
		fmt.Println("4. terminating the goroutine using a closed channel")
		ch := make(chan int)

		go func() {
			defer wg.Done()
			fmt.Println("goroutine 4 started")
			for range ch {
				time.Sleep(500 * time.Millisecond)
			}
			fmt.Println("goroutine 4 finished")
		}()

		time.Sleep(time.Second)
		close(ch)
		wg.Wait()
	}

	// функция внутри гоуритны выполняется
	{
		wg.Add(1)
		fmt.Println("5. execution of the goroutine until the internal function is executed")

		go func() {
			defer wg.Done()
			fmt.Println("goroutine 5 started")
			for i := 1; i < 3; i++ {
				time.Sleep(time.Second)
			}
			fmt.Println("goroutine 5 finished")
		}()

		wg.Wait()
	}
}
