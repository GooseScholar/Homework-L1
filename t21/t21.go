package main

import "fmt"

/*Реализовать паттерн «адаптер» на любом примере. (Неизвестный турист приехал в Европу)*/

//Интерфейс туриста socket.go
type socket interface {
	plugIntoEuropeanSocket()
}

//Турестический код tourist.go
type tourist struct {
}

func (t *tourist) plugIntoSocket(s socket) {
	fmt.Println("Турист вставляет зарядку в европейскую розетку.")
	s.plugIntoEuropeanSocket()
}

// Европейская зарядка european.go
type european struct {
}

func (e *european) plugIntoEuropeanSocket() {
	fmt.Println("Европейская зарядка подключилась")
}

//Неизвестная зарядка unknow.go
type unknown struct{}

func (u *unknown) plugIntoUnknownSocket() {
	fmt.Println("Неизвестный турист подключает свою зарядку")
}

//Адаптер unknownAdapter.go
type unknownAdapter struct {
	unknownSocket *unknown
}

func (u *unknownAdapter) plugIntoEuropeanSocket() {
	fmt.Println("Адаптер позволил поключить зарядку неизвестного стандарта к европейской розетке")
	u.unknownSocket.plugIntoUnknownSocket()
}

//main
func main() {
	tourist := &tourist{}
	european := &european{}

	tourist.plugIntoSocket(european)

	unknownSocket := &unknown{}
	unknownAdapter := &unknownAdapter{
		unknownSocket: unknownSocket,
	}

	tourist.plugIntoSocket(unknownAdapter)
}
