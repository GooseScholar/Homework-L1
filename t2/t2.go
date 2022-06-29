package main

import (
	"fmt"
	"os"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	var wg sync.WaitGroup
	//Входные данные
	array := [5]int{2, 4, 6, 8, 10}
	wg.Add(len(array))

	for _, a := range array {
		go sqIf(a, &wg)
	}
	wg.Wait()
}

//вывод квадрата числа в stdout
func sqIf(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(os.Stdout, a*a)
}
