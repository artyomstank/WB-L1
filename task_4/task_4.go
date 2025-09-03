package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// worker слушает задачи и проверяет контекст
func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // сигнал завершения
			fmt.Printf("воркер %d завершает работу (получен сигнал отмены)\n", id)
			return
		case j, ok := <-jobs:
			if !ok { // канал закрыт -> работы больше нет
				fmt.Printf("воркер %d: канал закрыт, завершаюсь\n", id)
				return
			}
			fmt.Printf("воркер %d получил задачу: %d\n", id, j)
			time.Sleep(2 * time.Second) // имитация работы
			fmt.Printf("воркер %d завершил задачу: %d\n", id, j)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("укажите количество воркеров в качестве аргумента.")
		fmt.Println("пример: go run task_4.go 3")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("количество воркеров должно быть положительным числом.")
		return
	}

	jobs := make(chan int, 10)

	// создаём контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// обрабатываем Ctrl+C
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("\nполучен сигнал Ctrl+C, завершаем...")
		cancel()    // оповещаем воркеров
		close(jobs) // закрываем отправку задач
	}()

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go worker(ctx, w, jobs, &wg)
	}

	// отправка задач
	for j := 1; j <= 30; j++ {
		select {
		case <-ctx.Done(): // если контекст уже завершён
			fmt.Println("отмена отправки задач")
			break
		case jobs <- j:
			fmt.Printf("главная горутина отправила задачу: %d\n", j)
		}
		time.Sleep(time.Second)
	}

	close(jobs) // закрываем отправку задач
	wg.Wait()
	fmt.Println("все воркеры завершили работу, программа завершена.")
}
