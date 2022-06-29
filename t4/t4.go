package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

func worker(messagesCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message, ok := <-messagesCh
		if !ok {
			return
		}
		d := time.Duration(message) * time.Millisecond
		time.Sleep(d)
		fmt.Fprintln(os.Stdout, message)
		fmt.Printf("Worker: %d\n", message)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	messagesCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(messagesCh, wg)
	}

	for i := 1; i > 0; i++ {
		messagesCh <- i
		time.Sleep(500 * time.Microsecond)
	}

	close(messagesCh)
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	N := 5
	wg.Add(N)
	go pool(&wg, N, 50)

	wg.Wait()
}
