package main

import (
	"fmt"
	"sort"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	array := []int{-13, 4, -5, -5, -4, 3, 1, 2, 5, 7, -1, 0}
	//Быстрая сорнировка
	sort.Ints(array)
	fmt.Println(array)
}
