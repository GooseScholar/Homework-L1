package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	do("Это строка")
	do(67)
	do(true)
	do(67.2)
}

/*Метод распознавания типа переменной. Сhannel невозможно передать в переменной типа пустой интерфейс.
В следствии этого принято решение опустить распознавание типа данных channel в методе распознавания.*/
func do(i interface{}) {

	switch t := i.(type) {
	case int:
		fmt.Printf("тип данных int: %d\n", t)
	case string:
		fmt.Printf("тип данных string: %s\n", t)
	case bool:
		fmt.Printf("тип данных bool: %v\n", t)
	default:
		fmt.Printf("тип данных не распознан: %v\n", t)
	}

}
