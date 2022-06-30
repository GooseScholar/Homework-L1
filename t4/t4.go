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

func main() {
	//Если в канале будут миллионы сообщений, которые нужно ещё обработать, то завершение получится долгим.
	//Выбран метод прерывания без ожидания выполнения всех горутин.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	N := 30
	for i := 0; i < N; i++ {
		wg.Add(1)
	}
	go pool(&wg, N)

	wg.Wait()
}

//Запуск воркеров и запись в канал
func pool(wg *sync.WaitGroup, workers int) {
	messagesCh := make(chan int)

	for i := 1; i <= workers; i++ {
		go worker(messagesCh, wg, i)
	}

	for i := 1; i > 0; i++ {
		messagesCh <- i
	}

	close(messagesCh)
}

//Воркер
func worker(messagesCh <-chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for {
		message, ok := <-messagesCh
		if !ok {
			return
		}
		d := time.Duration(message) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("Worker %d: ", i)
		fmt.Fprintln(os.Stdout, message*message)
	}
}
