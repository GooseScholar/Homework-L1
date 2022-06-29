package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20.
*/

func main() {
	//Введите значения переменных, от от –9 223 372 036 854 775 808 до 9 223 372 036 854 775 807

	a := big.NewInt(9048578974353342552)
	b := big.NewInt(4294933243423424436)
	fmt.Printf("Первая переменная: %v\nВторая пременная:  %v\n", a, b)
	plus(a, b)
	subtract(a, b)
	multiply(a, b)
	divide(a, b)

}

func plus(a, b *big.Int) {
	fmt.Printf("Результат сложения:  %v\n", big.NewInt(0).Add(a, b))
}

func subtract(a, b *big.Int) {
	fmt.Printf("Результат вычитания: %v\n", big.NewInt(0).Sub(a, b))
}

func multiply(a, b *big.Int) {
	fmt.Printf("Результат умножения: %v\n", big.NewInt(0).Mul(a, b))
}

func divide(a, b *big.Int) {
	fmt.Printf("Результат деления:   %v\n", big.NewInt(0).Div(a, b))
}
