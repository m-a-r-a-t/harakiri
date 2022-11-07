/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/


package main

import (
	"fmt"
)

var justString string

func main() {
	fmt.Println("а"[1:])
	someFunc()
}

func someFunc() {
	v := createHugeString(1 << 10)
	runes := []rune(v) // переводим строку в руны так как в строке могут быть символы юникода
	ind := 100
	fmt.Println(len(v))
	if ind <= len(runes)-1 { // проверяем доступен ли такой индекс в слайсе
		justString := runes[:100]
		fmt.Println("String:", string(justString))

	} else {
		fmt.Println("Граница слайса больше чем кол-во символов в строке")
	}

}
// функция создающая длинную юникод строку
func createHugeString(n int) string { 
	var s string
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			s += "Ж"
		} else {
			s += "B"
		}
	}
	return s
}
