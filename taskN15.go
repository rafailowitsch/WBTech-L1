package main

import (
	"fmt"
	"strings"
)

var justString string

// функция taskN15 демонстрирует работу функции someFunc и проблему срезов строк
func taskN15() {
	someFunc()
	fmt.Println(justString)
}

// функция createHugeString создает строку размером size символов
func createHugeString(size int) string {
	repeatedString := strings.Repeat("a", size)
	return repeatedString
}

// функция someFunc демонстрирует проблему срезов строк
func someFunc() {
	v := createHugeString(1 << 10) // создаем большую строку

	// корректное преобразование строки в срез рун и использование первых 100 рун
	justString = string([]rune(v)[:100])
}
