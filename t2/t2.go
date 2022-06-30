package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

//Конкурентно Вызывает функцию расчета и вызова квадратов элементов массива,
//для ожидания окончания выполнения горутин используется sync.WaitGroup
func main() {
	var wg sync.WaitGroup
	array := [5]int{2, 4, 6, 8, 10}
	wg.Add(len(array))

	for _, a := range array {
		go sqIf(a, &wg)
	}
	wg.Wait()
}

//Вывод квадрата числа в консоль
func sqIf(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(a * a)
}
