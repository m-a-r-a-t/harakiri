package main

import (
	"fmt"
	"math/big"
)

func main() {

	var num1, _ = new(big.Int).SetString("43554432321", 10)
	var num2, _ = new(big.Int).SetString("33554432", 10)

	fmt.Println("Умножение:", Mul(*num1, *num2))
	fmt.Println("Деление:", Div(*num1, *num2))
	fmt.Println("Складывание:", Sum(*num1, *num2))
	fmt.Println("Вычитание:", Sub(*num1, *num2))

}

func Mul(a, b big.Int) *big.Int {

	return new(big.Int).Mul(&a, &b)
}

func Sum(a, b big.Int) *big.Int {

	return new(big.Int).Add(&a, &b)
}

func Div(a, b big.Int) *big.Int {
	return new(big.Int).Div(&a, &b)
}

func Sub(a, b big.Int) *big.Int {
	return new(big.Int).Sub(&a, &b)
}
