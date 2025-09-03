package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("воркер %d получил задачу: %d\n", id, j)
		time.Sleep(2 * time.Second) // имитация работы
		fmt.Printf("воркер %d завершил задачу: %d\n", id, j)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("укажите количество воркеров в аргументе.")
		fmt.Println("пример: go run main.go 3")
		return
	}
	numWorkers, err := strconv.Atoi(os.Args[1])
	//валидация
	if err != nil || numWorkers <= 0 {
		fmt.Println("количество воркеров должно быть положительным числом.")
		return
	}

	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	//воркеры
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, &wg)
	}

	// отправка задач
	for j := 1; j <= 15; j++ {
		jobs <- j
		fmt.Printf("главная горутина отправила задачу: %d\n", j)
		time.Sleep(time.Second)
	}

	close(jobs) // закрываем канал -> воркеры завершатся
	wg.Wait()
	fmt.Println("все задачи обработаны, программа завершена.")
}
