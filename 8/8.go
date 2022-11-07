/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.*/

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {

	var num int64 = 152
	newNum, err := replaceNumberBit(num, 5, 1)

	if err != nil {
		panic(err)
	}

	fmt.Println("Num", num)
	fmt.Println("Changed num", newNum)

}

func replaceNumberBit(n int64, position int, binaryNum int) (int64, error) {

	var newNumInBinaryFormat string

	if binaryNum == 1 || binaryNum == 0 {
	} else {
		return -1, errors.New("binaryNum must be 1 or 0 !")
	}

	numInBinaryFormat := strconv.FormatInt(n, 2)

	if position >= len(numInBinaryFormat) {
		return -1, errors.New("the position must be less than or equal to the number of bits in the number")
	}

	fmt.Println(numInBinaryFormat)
	binaryNumLength := len(numInBinaryFormat)
	newNumInBinaryFormat = numInBinaryFormat[0:binaryNumLength-1-position] + strconv.Itoa(binaryNum) + numInBinaryFormat[binaryNumLength-position:]
	fmt.Println(newNumInBinaryFormat)

	newDecimalNum, err := strconv.ParseInt(newNumInBinaryFormat, 2, 64)

	if err != nil {
		fmt.Println(err)
	}

	return newDecimalNum, nil
}
