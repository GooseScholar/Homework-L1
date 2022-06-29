package main

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	// остановка горутины отправкой сигнала в специальный канал
	count := 1
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				count++
			}
		}
	}()
	quit <- true
}

/* Данный код прерывает все горутины
ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
*/
