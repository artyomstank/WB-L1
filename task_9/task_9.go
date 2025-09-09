package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 8, 9, 10}

	chX := make(chan int)
	ch2X := make(chan int)
	fmt.Println("Конвейер начал работу...")

	// запись чисел в канал chX
	go func() {
		for _, x := range nums {
			chX <- x
		}
		close(chX)
	}()

	// умножение на 2
	go func() {
		for x := range chX {
			ch2X <- x * 2
		}
		close(ch2X)
	}()

	// чтение результатов
	for result := range ch2X {
		fmt.Println(result)
	}

	fmt.Println("Конвейер завершил работу.")
}
