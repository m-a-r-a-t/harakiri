package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	c:= math.big
	a := int64(math.Pow(2, 63))
	b := int64(math.Pow(2, 63))

	fmt.Println(a * b)
}
