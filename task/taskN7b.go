package task

import (
	"fmt"
	"sync"
)

// функция PrintSyncMap выводит все ключи и значения из sync.Map
func PrintSyncMap(cm *sync.Map) {
	cm.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
		return true
	})
}

// функция taskN7b демонстрирует использование sync.Map для конкурентной записи данных
func TaskN7b() {
	var m sync.Map

	m.Store("foo", "bar")
	m.Store("bar", "baz")
	m.Store("baz", "foo")

	value, ok := m.Load("foo")
	if ok {
		fmt.Println("Value:", value)
	}

	PrintSyncMap(&m)

	m.Delete("foo")

	PrintSyncMap(&m)
}
