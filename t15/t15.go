package main

import "fmt"

/*
К каким негативным последствиям может привести данный фрагмент кода,
и как это исправить?
Приведите корректный пример реализации.
*/

/*Все верно работает, если все символы однобайтовые(например только англиский текст), иначе нужны руны

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/
var justString string

func main() {
	someFunc()
}

func someFunc() {
	//Сгенирировалась какая то строка, будем считать, что она 1024 символа (2^10)
	v := "g32ooprj`908he19082h0R2 CR2R[32TG453HALOIFWH812934Y`8910`2I10928`1HE209JEW`-12JW90`32EJH902`HJEW21JW	OMOPQWJ	JE09J29`0JS902Q`JSM2`90JE M290`EJ[0-0d0-dj32jedeeWDK20EU-90-Ъ032-АКО13ТРВ2ЁУ-0Ш1ЁБЪЦ-У Ш20	-ВОГЦЙавоцй а934а90Ц3РТАКца0-3оъ0каъ-АК0-Ъ3ГО2	3-К50Н6Ш89054ОНПТУФЫКЩЫПОК90ПГО4290	К34КАЗЩШ3УАЭК	9Щ2Х09H231Y89	RHE3H	20R4ghe$ tG,L"
	//Мы не знаем сколько байт каждый символ, поэтому превращаем большую строку из стринг в []rune, теперь каждый символ 4 байта
	r := []rune(v)
	//Для руны можно спокойно брать первые 100 символов и Никогда не будет обрезанных символов
	justString = string(r[:100])
	fmt.Println(justString)
}
