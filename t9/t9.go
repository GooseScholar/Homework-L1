package main

import (
	"fmt"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй —
результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	//Создаем два канала
	firstCh := make(chan int)
	secondCh := make(chan int)

	//Запись в первый канал
	go func() {
		for x := 0; x <= 10; x++ {
			firstCh <- x
		}
		close(firstCh)
	}()

	//Запись во второй канал
	go func() {
		for x := range firstCh {
			secondCh <- x * x
		}
		close(secondCh)
	}()

	//Вывот значений вт
	for x := range secondCh {
		fmt.Println(x)
	}
}
