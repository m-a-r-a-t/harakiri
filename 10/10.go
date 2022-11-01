package main

import "fmt"

func main() {
	tempratures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int64][]float64)

	for _, temp := range tempratures {
		div := int64(temp/10) * 10
		if _, ok := groups[div]; ok {
			groups[div] = append(groups[div], temp)

		} else {
			groups[div] = []float64{temp}
		}

	}

	fmt.Println(groups)

}
