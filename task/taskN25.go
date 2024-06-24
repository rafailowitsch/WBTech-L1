package task

import (
	"fmt"
	"time"
)

// функция sleep приостанавливает выполнение программы на заданное количество секунд
func sleep(sec int64) {
	currTime := time.Now().Unix()
	reqTime := currTime + sec

	for currTime < reqTime {
		currTime = time.Now().Unix()
	}
}

// функция taskN25 демонстрирует использование функции sleep
func TaskN25() {
	fmt.Print("Hello ")
	sleep(5)
	fmt.Println("world!")
}
