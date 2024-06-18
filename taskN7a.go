// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu sync.RWMutex // испольузем RWMutex, чтобы поддерживать конкуретное чтение из мапы
	m  map[string]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]interface{}),
	}
}

func (cm *ConcurrentMap) Set(key string, value interface{}) {
	cm.mu.Lock() // используем обычный метод блокировки мьютекса, чтобы заблокировать чтение
	defer cm.mu.Unlock()

	cm.m[key] = value
}

func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	cm.mu.RLock() // используем метод Read Lock, который позволяет нескольким горутинам считывать значение из мапы
	defer cm.mu.RUnlock()

	value, ok := cm.m[key]

	return value, ok
}

func (cm *ConcurrentMap) GetKeyValue() map[string]interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	copyMap := make(map[string]interface{}) // делаем копию мапы, чтобы соблюдать инкапсуляцию
	for key, value := range cm.m {
		copyMap[key] = value
	}

	return copyMap
}

func (cm *ConcurrentMap) Delete(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	delete(cm.m, key)
}

func taskN7a() {
	cm := NewConcurrentMap()

	cm.Set("foo", "bar")
	cm.Set("bar", "baz")
	cm.Set("baz", "foo")

	value, ok := cm.Get("foo")
	if ok {
		fmt.Println("Value:", value)
	}

	m := cm.GetKeyValue()
	fmt.Println(m)

	cm.Delete("foo")

	m = cm.GetKeyValue()
	fmt.Println(m)
}
