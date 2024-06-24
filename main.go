package main

import (
	"fmt"
	"os"
	"strconv"
	t "wbtech/task"
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

	//массив функций для вызова
	tasks := []func(){
		t.TaskN1, t.TaskN2, t.TaskN3, t.TaskN4, t.TaskN5, t.TaskN6,
		func() { t.TaskN7a(); t.TaskN7b() }, // объединение t.TaskN7a и t.TaskN7b
		t.TaskN8, t.TaskN9, t.TaskN10, t.TaskN11, t.TaskN12, t.TaskN13,
		t.TaskN14, t.TaskN15, t.TaskN16, t.TaskN17, t.TaskN18, t.TaskN19,
		t.TaskN20, t.TaskN21, t.TaskN22, t.TaskN23, t.TaskN24, t.TaskN25, t.TaskN26,
	}

	if taskNumber >= 1 && taskNumber <= len(tasks) {
		tasks[taskNumber-1]()
	} else {
		fmt.Println("Unknown task number")
	}
}
