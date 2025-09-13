package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mtx   sync.Mutex
}

func (c *Counter) increase() {
	c.mtx.Lock()
	c.count += 1
	c.mtx.Unlock()
}

func main() {
	ctr := Counter{count: 0}
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			ctr.increase()
			fmt.Printf("я %d горутина \n", n)
		}(i)
	}
	wg.Wait()
	fmt.Println(ctr.count)
}
