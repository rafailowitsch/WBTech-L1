package task

import (
	"fmt"
	"sync"
)

// функция routineCalc рассчитывает квадраты чисел из слайса и выводит их в stdout
func routineCalc(slice []int) {
	var wg sync.WaitGroup

	// запускаем горутину для каждого числа в слайсе
	for _, num := range slice {
		wg.Add(1)
		go func(num int) {
			defer wg.Done() // уменьшаем счетчик горутин по завершении
			fmt.Println(num, ": ", num*num)
		}(num)
	}
	wg.Wait() // ждем завершения всех горутин
}

func TaskN2() {
	routineCalc([]int{2, 4, 6, 8, 10})
}
