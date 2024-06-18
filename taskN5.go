package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func taskN5() {
	duration := time.Duration(30) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	values := []int{1, 2, 3, 4, 5}
	ch := make(chan int, 2)

	go func(values []int) {
		defer close(ch)
		for {
			for _, val := range values {
				ch <- val
				fmt.Printf("Msg sent: %d\n", val)
			}
		}
	}(values)

	go func(messages <-chan int) {
		rand.New(rand.NewSource(time.Now().Unix()))
		for msg := range messages {
			randomNumber := rand.Intn(180) + 60
			time.Sleep(time.Duration(randomNumber) * time.Millisecond)
			fmt.Printf("Msg [%d] processed\n", msg)
		}
	}(ch)

	for {
		select {
		case <-ctx.Done():
			println("program interrupted")
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

/*
	a = 3
	b = 4

	a = 3
	b = b + a = 7



*/
