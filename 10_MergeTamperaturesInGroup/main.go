package main

import (
	"fmt"
	"sort"
)

/*
	Дана последовательность температурных колебаний:
		-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов.
	Последовательность в подмножноствах не важна.

	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	temperatures := []float64{-25.4, -15, -16.6, -19, -20, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	subsets := GetSubsets(&temperatures)
	for key, values := range *subsets {
		fmt.Printf("key %d: {\n", key)
		for _, v := range values {
			fmt.Printf("\tvalues %.1f\n", v)
		}
		fmt.Println("},")
	}
}

func GetThreshold(f float64) int {
	res := int(f) - (int(f) % 10)
	if f < 0 {
		res -= 10
	}
	return res
}

func GetSubsets(temp *[]float64) *map[int][]float64 {
	sort.Float64s(*temp)
	min := GetThreshold((*temp)[0])
	result := make(map[int][]float64)
	for _, t := range *temp {
		if (min + 10) <= int(t) {
			min = GetThreshold(t)
		}
		result[min] = append(result[min], t)
	}
	return &result
}
