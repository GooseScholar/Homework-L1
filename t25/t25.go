package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {
	fmt.Println(time.Now())
	sleep(1)
	fmt.Println(time.Now())
	sleep(3)
	fmt.Println(time.Now())
	sleep(2)
	fmt.Println(time.Now())
	sleep(7)
	fmt.Println(time.Now())
	sleep(1)
	fmt.Println(time.Now())
}

func sleep(n int) {
	timeout := time.After(time.Duration(n) * time.Second)
	c := make(chan int, 1)

	//Цикл for ждет спецсигнала - timeout, после чего завершается
	for {
		select {
		case <-timeout:
			fmt.Printf("Времени прошло: %v\n", timeout)
			close(c)
			return
		}
	}
}
