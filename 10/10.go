/*Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.*/

package main

import "fmt"

func main() {
	tempratures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int64][]float64)

	for _, temp := range tempratures {
		div := int64(temp/10) * 10 // определяем группу у каждого числа
		if _, ok := groups[div]; ok {
			groups[div] = append(groups[div], temp)

		} else {
			groups[div] = []float64{temp}
		}

	}

	fmt.Println(groups)

}
