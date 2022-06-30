package main

import "fmt"

/* Задача 1
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
*/

//Структура Human
type Human struct {
	name string
	age  int
	sex  string
}

//Функция которую умеет делать Human
func (h *Human) Say() {
	fmt.Println("Say my name")
}

//Структура Action которой передали имя структуры Human
type Action struct {
	Human
	id int
}

func main() {
	//Теперь она умеет делать ту же функцию что и хумен
	Ex := new(Action)
	Ex.Say()
}
