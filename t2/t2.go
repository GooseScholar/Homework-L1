package main

import (
	"fmt"
	"os"
	"time"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	//тестовый массив
	array := [5]int{2, 4, 6, 8, 10}

	for _, a := range array {
		go sqIf(a)
	}

	time.Sleep(1 * time.Second)
}

//вывод квадрата числа в stdout
func sqIf(a int) {
	fmt.Fprintln(os.Stdout, a*a)
}
