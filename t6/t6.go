package main

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

//Первый способ - ожидание сигнала завершения
func main() {
	quit := make(chan bool)
	for i := 0; i >= 0; i++ {
		go func(i int) {

			select {
			case <-quit:
				return
			default:

			}
		}(i)
	}
	//..
	quit <- true
}

/* //Второй способ отсановку по прошествии времени
func main() {

	to := time.After(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-to:
				fmt.Println("Время истекло")
				done <- true
				return

			}
		}
		}()

		<-done
	}
*/

/* //Третий способ ожидание контекста
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var count int32 = 1
	for i := 1; i > 0; i++ {
		go func(i int, ctx context.Context) {
			defer cancel()
			atomic.AddInt32(&count, 1)
		}(i, ctx)
	}
}
*/
