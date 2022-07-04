package main

import "fmt"

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

//Первый способ - остановка по сигналу завершения
func main() {
	quit := make(chan bool)
	go func() {
		fmt.Println("В горутине")
		for {
			select {
			case <-quit:
				return
			default:

				//Что-нибудь делаем
			}
		}
	}()
	//..
	fmt.Println("Перед выходом")
	quit <- true
	fmt.Println("После выхода")
}

/*
//Второй способ: остановка по таймеру
func main() {

	timeIsOver := time.After(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-timeIsOver:
				fmt.Println("Время истекло")
				done <- true
				return

			}
		}
	}()

	<-done
}
*/
/*
//Третий способ ожидание таймера контекста
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		for {
			fmt.Println("мы тут")
			select {
			case <-ctx.Done():
				fmt.Println("Время истекло 1")
				return
			default:
				fmt.Println("Мы тут 2")
			}

		}
	}(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("Время истекло 2")
}
*/
