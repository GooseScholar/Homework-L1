package main

import (
	"fmt"
)

/*Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.*/

func main() {
	//Входные данные
	input := []rune("главрыба")
	fmt.Println(string(input))
	//Разбираем строку на слайс, где каждому индексу приравнивается только один символ строки input
	//Переворачиваем слайс
	swap(input)
	//Сложение слайса в строку и забись в переменную output
	output := string(input)
	fmt.Println(output)

}

func swap(a []rune) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
