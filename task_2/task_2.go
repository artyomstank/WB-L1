package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, obj := range arr {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x * x)
		}(obj)
	}
	wg.Wait()
}
