package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	//валидация ввода
	if len(os.Args) < 2 {
		fmt.Println("укажите количество секунд в качестве аргумента.")
		fmt.Println("пример: go run task_5.go 4")
		return
	}
	// читаем аргумент
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("время работы должно быть положительным числом.")
		return
	}

	ch := make(chan int)
	timeoutDuration := time.Duration(n) * time.Second
	timeoutChan := time.After(timeoutDuration)

	fmt.Printf("программа работает(сек)  %v.\n", timeoutDuration)

	// отправка
	go func() {
		for i := 1; ; i++ {
			select {
			case <-timeoutChan: // если время вышло, закрываем канал и выходим
				close(ch)
				return
			case ch <- i:
				time.Sleep(time.Second)
			}
		}
	}()

	//получение значений
	for val := range ch {
		fmt.Printf("получено значение: %d\n", val)
	}

	fmt.Println("таймаут, завершение програмы.")
}
