package main

import (
	"fmt"
	"strings"
)

/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	//входные данные
	input := "snow dog sun"
	fmt.Printf("До: %s\n", input)
	//разбиваем строку input по элементу пробел " " на отдельные
	message := strings.Fields(input)
	//переворачиваем слова в строке
	reverse(message)
	//объединяем срез в строку и вставляем пробел между словами
	output := strings.Join(message, " ")
	fmt.Printf("После: %s\n", output)

}

func reverse(message []string) {
	for i, j := 0, len(message)-1; i < j; i, j = i+1, j-1 {
		message[i], message[j] = message[j], message[i]
	}
}
