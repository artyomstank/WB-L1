package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 1. завершение по условию
func condExit(iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("горутина с входом по условию начинает работу.")
	for i := 0; i < iterations; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("условие выполнено, горутина завершилась.")
}

// 2. авершение по сигналу через канал
func doneCh(done <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("горутина с выходом по сигналу из канала начинает работу.")
	<-done
	fmt.Println("сигнал получен из канала, горутина завершилась.")
}

// 3. завершение по отмене контекста
func ctxDone(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("горутина с отменой контекста начинает работу.")
	<-ctx.Done()
	fmt.Println("контекст отменен, горутина завершилась.")
}

// 4. завершение по закрытию канала
func closeCh(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("горутина с закрытием канала начинает работу.")
	for range tasks {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("канал закрыт, горутина завершилась.")
}

// 5. завершение с помощью runtime.Goexit()
func runtExit(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("горутина с вызовом runtime.Goexit() завершена.")
	fmt.Println("горутина с вызовом runtime.Goexit() начинает работу.")
	runtime.Goexit()
}

// 6. завершение через панику (recover)
func panicExit(wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("паника обработана")
		}
		wg.Done()
	}()
	fmt.Println("горутина с паникой начинает работу.")
	panic("сообщение паники")
}

func main() {
	var wg sync.WaitGroup

	//завершение по условию
	wg.Add(1)
	go condExit(5, &wg)
	wg.Wait()

	//авершение по сигналу через канал
	doneChan := make(chan bool)
	wg.Add(1)
	go doneCh(doneChan, &wg)
	doneChan <- true
	wg.Wait()

	//завершение по отмене контекста
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go ctxDone(ctx, &wg)
	cancel()
	wg.Wait()

	// завершение по закрытию канала
	tasksChan := make(chan int)
	wg.Add(1)
	go closeCh(tasksChan, &wg)
	tasksChan <- 1
	tasksChan <- 2
	tasksChan <- 3
	close(tasksChan)
	wg.Wait()

	//завершение с помощью runtime.Goexit()
	wg.Add(1)
	go runtExit(&wg)
	wg.Wait()

	//завершение через панику (recover)
	wg.Add(1)
	go panicExit(&wg)
	wg.Wait()
}
