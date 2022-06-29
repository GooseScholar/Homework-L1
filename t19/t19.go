package main

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.*/

func main() {
	//Входные данные
	input := "главрыба"
	fmt.Println(input)
	//Разбираем строку на слайс, где каждому индексу приравнивается только один символ строки input
	array := strings.Split(input, "")
	//Переворачиваем слайс
	swap(array)
	//Сложение слайса в строку и забись в переменную output
	output := strings.Join(array, "")
	fmt.Println(output)

}

func swap(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
