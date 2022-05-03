package main

import (
	"errors"
	"fmt"
)

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

/*
	Бинарный поиск возможен только в отсортированном массиве!
*/

func main() {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr = append(arr, i-50)
	} // Отсортированный массив от -50 до +50

	// Ищем под каким индексом находится значение `33` в массиве
	index, err := FindIndex(arr, 33)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("index:", index, "\nvalue:", arr[index])
	}

}

func FindIndex(arr []int, value int) (index int, ok error) {
	if len(arr) < 1 { // если слайс пустой, значит не нашли ничего
		return -1, errors.New("not found")
	}
	index = len(arr) >> 1   // Обращаемся к индексу в середине слайса
	if arr[index] > value { // если значение больше элемента посередине, то ищем в правой половине слайса, рекурсивно
		return FindIndex(arr[:index], value)
	} else if arr[index] < value { // иначе если значение меньше значения посередине то ищем в левой половине слайса, рекурсивно
		findIndex, err := FindIndex(arr[index+1:], value)
		if err != nil {
			return -1, err
		}
		return index + findIndex + 1, nil
	}
	return index, nil // Иначе мы нашли значение
}
