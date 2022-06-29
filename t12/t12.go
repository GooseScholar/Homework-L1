package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	//Входные данные в виде строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	//Множество в виде мапы
	m := make(map[string]struct{})
	//Запись всех строк в множество (мапу). Строки входных данных являются ключами мапы.
	for _, str := range strings {
		m[str] = struct{}{}
	}
	fmt.Println(m)
}
