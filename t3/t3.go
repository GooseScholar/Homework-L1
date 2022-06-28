package main

import (
	"fmt"
	"time"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2*2+3*3+4*4….) с
использованием конкурентных вычислений.
*/

func main() {
	//Входные данные
	array := []int{2, 4, 6, 8, 10}
	sum := 0

	//Поэлеменьный расчет суммы квадратов элементов массива через цикл
	for _, a := range array {
		go sumSq(a, &sum)
	}

	time.Sleep(1 * time.Second)

	fmt.Println(sum)
}

func sumSq(a int, sum *int) {
	*sum += a * a
}
