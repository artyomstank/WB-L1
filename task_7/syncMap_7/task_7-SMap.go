package main

import (
	"fmt"
	"strconv"
	"sync"
)

// встроенная мапа
func main() {
	var safeMap sync.Map
	var wg sync.WaitGroup
	const numGoroutines = 1000

	fmt.Printf("запись в sync.Map %d элементов...\n", numGoroutines)

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(n int) {
			defer wg.Done()
			key := "key_" + strconv.Itoa(n)
			safeMap.Store(key, n)
		}(i)
	}

	wg.Wait()

	//кол-во элементов
	count := 0
	safeMap.Range(func(key, value any) bool {
		count++
		return true
	})

	fmt.Printf("в sync.Map хранится %d элементов.\n", count)
}
