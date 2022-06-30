package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func entry(c chan int) {
	for i := 1; i > 0; i++ {
		c <- i
	}

	close(c)
}

func main() {
	//Ввести время работы программы (в секундах)
	N := 5
	timeout := time.After(time.Duration(N) * time.Second)
	c := make(chan int, 1)

	go entry(c)

	for {
		select {
		case <-timeout:
			fmt.Println("Время истекло")
			close(c)
			return
		default:
			val, ok := <-c
			if ok {
				fmt.Println(val)
			}
		}
	}
}
