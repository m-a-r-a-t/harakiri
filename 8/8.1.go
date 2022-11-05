package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func main() {

	var num int64 = 92233720368
	start := time.Now().UnixNano()
	newNum, err := replaceNumberBit(num, 1, 1)
	fmt.Println("Time", time.Now().UnixNano()-start)
	if err != nil {
		panic(err)
	}

	fmt.Println("Num", num)
	fmt.Println("Changed num", newNum)

}

func replaceNumberBit(d int64, position int64, replaceWithNum int64) (int64, error) {

	if replaceWithNum == 1 || replaceWithNum == 0 {
	} else {
		return d, errors.New("binaryNum must be 1 or 0 !")
	}

	var i int64 = 0
	isChanged := false

	var binaryNum int64 = 0
	var digitOfNumber int64 = 1
	for d != 0 {
		fmt.Println(digitOfNumber)
		mod := d % 2
		d = d / 2

		if i == position {
			binaryNum += replaceWithNum * digitOfNumber
			isChanged = true
		} else {
			binaryNum += mod * digitOfNumber
		}

		digitOfNumber *= 10
		i++
	}

	var newInt int64
	var k float64 = 0.0
	for binaryNum != 0 {
		fmt.Println(binaryNum, binaryNum%10, int64(math.Pow(2, k))*(binaryNum%10))

		mode := binaryNum % 10
		binaryNum = binaryNum / 10
		newInt += mode * int64(math.Pow(2, k))
		k++
	}

	fmt.Println(k)
	// newInt, _ := strconv.ParseInt(strconv.FormatInt(binaryNum, 10), 2, 64)

	if !isChanged {
		return newInt, errors.New("the position must be less than or equal to the number of bits in the number")
	}

	return newInt, nil
}
