package main

import (
	"fmt"
	"strings"
)

// функция reverseWords переворачивает слова в строке
func reverseWords(str string) string {
	words := strings.Fields(str) // разбиваем строку на слова

	l := len(words)
	mid := l / 2

	// меняем местами слова
	for i := 0; i < mid; i++ {
		words[i], words[l-1-i] = words[l-1-i], words[i]
	}

	return strings.Join(words, " ") // объединяем слова обратно в строку
}

// функция taskN19 демонстрирует переворачивание слов в строке
func taskN20() {
	str := "snow dog sun moon"
	fmt.Println("String before reverse words:", str)

	str = reverseWords(str)
	fmt.Println("String after reverse words:", str)
}
