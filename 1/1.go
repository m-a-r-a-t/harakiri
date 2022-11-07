/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от
родительской структуры Human (аналог наследования).*/

package main

import "fmt"

type Human struct {
	height int32
	weight int32
}

type Action struct {
	// Встраивание структуры Human в Action
	Human
}

// Определение метода Move для структуры Human
func (h Human) Move() {
	fmt.Println("Human moving !")
}

// Определение метода GetWeight для структуры Human
func (h Human) GetWeight() int32 {
	return h.weight
}

func main() {

	human := Human{
		weight: 100,
		height: 200,
	}

	action := &Action{Human: human}

	action.Move()
	weight := action.GetWeight()
	fmt.Println(weight)
}
