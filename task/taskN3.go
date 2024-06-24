package task

import (
	"fmt"
	"sync"
)

// функция getSumSquares рассчитывает сумму квадратов чисел из слайса с использованием горутин
func getSumSquares(slice []int) int {
	var wg sync.WaitGroup
	var sum int = 0
	var mu sync.Mutex // добавляем мьютекс для защиты доступа к переменной sum

	for _, num := range slice {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()         // блокируем доступ к sum
			defer mu.Unlock() // разблокируем доступ к sum после выполнения
			sum += num * num
		}(num)
	}
	wg.Wait()

	return sum
}

func TaskN3() {
	ss := getSumSquares([]int{2, 3, 4, 5})
	fmt.Println(ss)
}
