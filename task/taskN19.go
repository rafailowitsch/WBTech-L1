package task

import "fmt"

// функция reverseString переворачивает строку
func reverseString(str string) string {
	runes := []rune(str) // преобразуем строку в срез рун

	l := len(runes)
	mid := l / 2

	// меняем местами символы
	for i := 0; i < mid; i++ {
		runes[i], runes[l-1-i] = runes[l-1-i], runes[i]
	}

	return string(runes) // преобразуем срез рун обратно в строку
}

// функция taskN18 демонстрирует переворачивание строки
func TaskN19() {
	str := "главрыба"
	fmt.Println("String before reverse:", str)

	str = reverseString(str)
	fmt.Println("String after reverse:", str)
}
