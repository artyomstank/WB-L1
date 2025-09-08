package main

import (
	"fmt"
	"strconv"
	"sync"
)

// наша структура с мьютексом
type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// запись в map
func (m *SafeMap) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

// чтение длины map
func (m *SafeMap) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.data)
}

func main() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup
	const numGoroutines = 1000

	fmt.Printf("Запись в карту %d элементов...\n", numGoroutines)

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			key := "key_" + strconv.Itoa(n)
			safeMap.Set(key, n)
		}(i)
	}

	wg.Wait()

	fmt.Printf("В карте хранится %d элементов.\n", safeMap.Len())
}
