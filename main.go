package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a task number")
		return
	}

	taskNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid task number")
		return
	}

	// массив функций для вызова
	tasks := []func(){
		taskN1, taskN2, taskN3, taskN4, taskN5, taskN6,
		func() { taskN7a(); taskN7b() }, // объединение taskN7a и taskN7b
		taskN8, taskN9, taskN10, taskN11, taskN12, taskN13,
		taskN14, taskN15, taskN16, taskN17, taskN18, taskN19,
		taskN20, taskN21, taskN22, taskN23, taskN24, taskN25, taskN26,
	}

	if taskNumber >= 1 && taskNumber <= len(tasks) {
		tasks[taskNumber-1]()
	} else {
		fmt.Println("Unknown task number")
	}
}
