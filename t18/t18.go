package main

import (
	"log"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

//Первый счетчик работает через atomic
type Counter1 struct {
	count int32
}

func NewCounter1() *Counter1 {
	return &Counter1{
		count: 0,
	}
}

func (c *Counter1) Counter1Print() int32 {
	return c.count
}

//Второй счетчик работае через RWMutex
type Counter2 struct {
	sync.RWMutex
	count int
}

func NewCounter2() *Counter2 {
	return &Counter2{
		count: 0,
	}
}

func (c *Counter2) Counter2Add() {
	c.RLock()
	defer c.RUnlock()
	c.count++
}

func (c *Counter2) Counter2Print() int {
	return c.count
}

func main() {
	C1 := NewCounter1()
	C2 := NewCounter2()

	wg := new(sync.WaitGroup)
	//Количество горутин
	for i := 0; i < 31; i++ {
		wg.Add(1)
		go DoSomething(C1, C2, wg)
	}

	wg.Wait()
	log.Printf("Первый счетчик: %d", C1.Counter1Print())
	log.Printf("Второй счетчик: %d", C2.Counter2Print())
}

func DoSomething(c1 *Counter1, c2 *Counter2, wg *sync.WaitGroup) {

	go func(c1 *Counter1, c2 *Counter2, wg *sync.WaitGroup) {
		defer wg.Done()
		atomic.AddInt32(&c1.count, 1)
		c2.Counter2Add()
	}(c1, c2, wg)

}
