package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2*2+3*3+4*4….) с
использованием конкурентных вычислений.
*/

//Для ожидания завершения всех горутин используется sync.WaitGroup
func main() {
	var wg sync.WaitGroup
	array := []int32{2, 4, 6, 8, 10}
	wg.Add(len(array))
	var sum int32 = 0

	//Поэлеменьный расчет суммы квадратов элементов массива через цикл и горутины
	for _, a := range array {
		go func(a int32) {
			defer wg.Done()
			atomic.AddInt32(&sum, a*a)
		}(a)
	}

	wg.Wait()
	fmt.Println(sum)
}
