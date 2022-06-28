package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	//конструктор точек
	point1 := NewPoint(7, 4)
	point2 := NewPoint(53, 13)

	//расчет дистанции с округлением до целых
	distance := math.Round(math.Sqrt(float64(sq(point1.x-point2.x) + sq(point1.y-point2.y))))

	fmt.Println(distance)

}

//функция возведения в квадрат
func sq(a int) int {
	return a * a
}
