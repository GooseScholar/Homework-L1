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
	//Введите значения переменных строки (любые целочисленные значения в десятичной системе)
	valueA := "90485789743533425524534643543"
	valueB := "429493324342342232134436"

	a := new(big.Int)
	a.SetString(valueA, 10)
	b := new(big.Int)
	b.SetString(valueB, 10)
	fmt.Printf("Первая переменная: %v\nВторая  пременная: %v\n", a, b)
	plus(a, b)
	subtract(a, b)
	multiply(a, b)
	divide(a, b)

}

func plus(a, b *big.Int) {
	fmt.Printf("Результат  сложения: %v\n", big.NewInt(0).Add(a, b))
}

func subtract(a, b *big.Int) {
	fmt.Printf("Результат вычитания: %v\n", big.NewInt(0).Sub(a, b))
}

func multiply(a, b *big.Int) {
	fmt.Printf("Результат умножения: %v\n", big.NewInt(0).Mul(a, b))
}

func divide(a, b *big.Int) {
	fmt.Printf("Результат   деления: %v\n", big.NewInt(0).Div(a, b))
}
