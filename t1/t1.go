package main

import "fmt"

/* Задача 1
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
*/

type Human struct {
	name string
	age  int
	sex  string
}

func (h *Human) Say() {
	fmt.Println("Say my name")
}

type Action struct {
	Human
	id int
}

func main() {
	Ex := new(Action)
	Ex.Say()
}
