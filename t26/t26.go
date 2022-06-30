package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	var a, b, c string = "abcd", "abCdefAaf", "aabcd"
	fmt.Println(checkMessage(a))
	fmt.Println(checkMessage(b))
	fmt.Println(checkMessage(c))

}

//Работает только с латинскими буквами
func checkMessage(m string) (string, bool) {
	var bitset uint32 = 0
	for _, ch := range strings.ToLower(m) {
		if unicode.IsLetter(ch) {
			if index := ch - 'a'; bitset&(1<<index) == 0 {
				bitset |= 1 << index
			} else {
				return m, false
			}
		}
	}
	return m, true
}
