package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

//Первый способ map+RWMutex
type Example struct {
	mx   sync.RWMutex
	Data map[int]int
}

//Создание map
func NewExample() *Example {
	return &Example{
		Data: make(map[int]int),
	}
}

//Запись данных в map
func (c *Example) Put(id int, o int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.Data[id] = o
}

//Получение данных из map
func (c *Example) Get(id int) (o int, b bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	o, b = c.Data[id]
	return
}

//Второй способ  sync.Map

func main() {
	//sync.Map
	var example sync.Map
	//запись
	example.Store("example", false)
	//чтение
	v, ok := example.Load("example")
	var val interface{}
	if ok {
		val = v
	}

	fmt.Println(val)

}
