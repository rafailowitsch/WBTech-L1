package main

import (
	"fmt"
	"strings"
)

// функция taskN26 демонстрирует проверку уникальности символов в строках
func taskN26() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(checkUnique(str1)) // true
	fmt.Println(checkUnique(str2)) // false
	fmt.Println(checkUnique(str3)) // false
}

// функция checkUnique проверяет, что все символы в строке уникальны
func checkUnique(str string) bool {
	m := make(map[rune]bool)
	str = strings.ToLower(str) // преобразуем строку к нижнему регистру для регистронезависимой проверки

	for _, char := range str {
		if m[char] {
			return false
		}
		m[char] = true
	}
	return true
}
