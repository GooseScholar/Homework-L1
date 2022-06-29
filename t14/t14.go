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
	do(make(chan int))
	do(make(chan string))
	do(make(chan bool))
	do(make(chan float32))

}

func do(i interface{}) {

	switch t := i.(type) {
	case int:
		fmt.Printf("тип данных int: %d\n", t)
	case string:
		fmt.Printf("тип данных string: %s\n", t)
	case bool:
		fmt.Printf("тип данных bool: %v\n", t)
	case chan int:
		fmt.Printf("тип данных chan int\n")
	case chan string:
		fmt.Printf("тип данных chan string\n")
	case chan bool:
		fmt.Printf("тип данных chan bool\n")
	default:
		fmt.Printf("тип данных не распознан\n")
	}

}
