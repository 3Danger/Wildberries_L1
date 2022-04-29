package main

import "fmt"

/*
	Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	arr1 := []int{9, 2, 16, 14, 12, 7, 11, 6, 4, 3, 10}
	arr2 := []int{12, 0, 9, 1, 4, 13, 2, 5, 6, 8, 16, 3}
	//arr2 := []int{22, 23, 54, 3453, 232, 23, 2, 42, 4, 24234} // 2, 4
	fmt.Println(*oneWay(&arr1, &arr2))
}

func oneWay(arr1, arr2 *[]int) *[]int {

	// Хочу что бы первая ссылка на массив
	// указывала на минимальный по размеру из двух
	// об этом чуть ниже
	if len(*arr1) > len(*arr2) {
		arr1, arr2 = arr2, arr1
	}

	// Задаю максимально возможный Capacity из минимального размера массива
	// Потому что итоговый результат точно не будет больше минимального размера массива
	// Таким образом гарантировано избежим реалокацию памяти
	result := make([]int, 0, len(*arr1))

	// Поскольку в golang отсутствует контейнер SET что имеется в C++ то использую MAP
	// в качестве значения для map использую заглушку struct{}{}
	// struct{}{} весит 0 байтов,
	intersection := make(map[int]struct{})

	for i := 0; i < len(*arr1); i++ {
		intersection[(*arr1)[i]] = struct{}{}
	}
	for i := 0; i < len(*arr2); i++ {
		_, ok := intersection[(*arr2)[i]]
		if ok {
			result = append(result, (*arr2)[i])
		}
	}
	return &result
}
