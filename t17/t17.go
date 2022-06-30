package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	//Слайс должен быть отсортирован в порядке возрастания(убываения) перед началом бинарного поиска
	a := []int{1, 2, 6, 9, 13, 22, 28, 37, 42, 49}
	//Значение которое необходимо найти
	x := 6
	//Бинарный поиск
	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("Найдено число %d его индекс равен %d в %v\n", x, i, a)
	} else {
		fmt.Printf("Число %d не найден в %v\n", x, a)
	}
}
