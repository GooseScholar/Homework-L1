package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 1
	b := 2
	a, b = b, a
	fmt.Printf("a=%d, b=%d \n", a, b)
}
