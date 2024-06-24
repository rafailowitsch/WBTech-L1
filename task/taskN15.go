package task

import (
	"fmt"
	"strings"
)

var justString string

func TaskN15() {
	someFunc()
	fmt.Println(justString)
}

// функция createHugeString создает строку размером size символов
func createHugeString(size int) string {
	repeatedString := strings.Repeat("a", size)
	return repeatedString
}

func someFunc() {
	v := createHugeString(1 << 10) // создаем большую строку

	// корректное преобразование строки в срез рун и использование первых 100 рун
	// мы создаем новую строку на основе среза рун
	// соответственно сборщик мусора при окончании функции освобождает память из под переменной v
	justString = string([]rune(v)[:100])
}
