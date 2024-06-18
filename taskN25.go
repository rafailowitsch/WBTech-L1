package main

import (
	"fmt"
	"time"
)

func sleep(sec int64) {
	currTime := time.Now().Unix()
	reqTime := currTime + sec - 1

	for currTime <= reqTime {
		currTime = time.Now().Unix()
	}
}

func taskN25() {
	fmt.Print("Hello ")
	sleep(5)
	fmt.Println("world!")
}
