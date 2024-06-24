package task

import (
	"fmt"
	"sync"
)

// структура ConcurrentMap содержит мапу и мьютекс для обеспечения конкурентного доступа
type ConcurrentMap struct {
	mu sync.RWMutex // используем RWMutex для поддержки конкурентного чтения
	m  map[string]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[string]interface{}),
	}
}

// метод Set устанавливает значение в мапу с блокировкой на запись
func (cm *ConcurrentMap) Set(key string, value interface{}) {
	cm.mu.Lock() // используем обычный метод блокировки мьютекса, чтобы заблокировать чтение
	defer cm.mu.Unlock()

	cm.m[key] = value
}

// метод Get получает значение из мапы с блокировкой на чтение
func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	cm.mu.RLock() // используем метод Read Lock, который позволяет нескольким горутинам считывать значение из мапы
	defer cm.mu.RUnlock()

	value, ok := cm.m[key]

	return value, ok
}

// метод GetKeyValue возвращает копию мапы
func (cm *ConcurrentMap) GetKeyValue() map[string]interface{} {
	cm.mu.RLock()
	defer cm.mu.RUnlock()

	copyMap := make(map[string]interface{}) // делаем копию мапы для инкапсуляции
	for key, value := range cm.m {
		copyMap[key] = value
	}

	return copyMap
}

// метод Delete удаляет значение из мапы с блокировкой на запись
func (cm *ConcurrentMap) Delete(key string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	delete(cm.m, key)
}

// функция taskN7a создает и использует экземпляр ConcurrentMap
func TaskN7a() {
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
