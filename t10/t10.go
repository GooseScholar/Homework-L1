package main

import "fmt"

/*Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	//Входные данные
	input := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//Сорздание мапы для объединения значений
	m := make(map[int][]float32)
	//int(a) - отбрасывается дробная часть, int(a)/10 -вычисление кол-ва десятков, (int(a)/10)*10 - создание групп кратных 10
	for _, a := range input {
		m[(int(a)/10)*10] = append(m[(int(a)/10)*10], a)
	}

	fmt.Println(m)
}