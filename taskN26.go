package main

import "fmt"

func taskN26() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(checkUnique(str1), checkUnique(str2), checkUnique(str3))
}

func checkUnique(str string) bool {
	m := make(map[rune]bool)

	for _, char := range str {
		if !m[char] {
			m[char] = true
		} else {
			return false
		}
	}
	return true
}
