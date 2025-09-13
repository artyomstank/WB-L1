package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 8, 9, 10}

	chX := make(chan int)
	ch2X := make(chan int)

	fmt.Println("Конвейер начал работу...")

	var wg sync.WaitGroup

	// запись чисел в канал chX
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, x := range nums {
			chX <- x
		}
		close(chX)
	}()

	// умножение на 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range chX {
			ch2X <- x * 2
		}
		close(ch2X)
	}()

	// чтение результатов
	for result := range ch2X {
		fmt.Println(result)
	}

	// ждём завершения всех горутин
	wg.Wait()

	fmt.Println("Конвейер завершил работу.")
}
