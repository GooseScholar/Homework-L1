package main

import (
	"fmt"
	"math"
)

/* ЗАДАЧА 8

Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

*/

func main() {
	//переменная, можно задать значение
	var a int = 7
	//номер бита
	i := 1
	//установить в 1 - true, в 0 - false
	b := false
	if b == true {
		//установка i-ого бита числа "a" в 1
		fmt.Println(a | int(math.Pow(2, float64(i-1))))
	} else {
		//установка i-ого бита числа "a" в 0
		fmt.Println(a &^ int(math.Pow(2, float64(i-1))))
	}

}
