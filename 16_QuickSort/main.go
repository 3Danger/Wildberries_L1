package main

import "fmt"

/*
	Реализовать быструю сортировку массива
	(quicksort) встроенными методами языка.
*/

func main() {
	arr := []int{3, 2, 6, 10, 3, 8, 7, 3, 1, 0, 5}
	fmt.Println("Before:", arr)
	QuickSort(arr)
	fmt.Println("After: ", arr)
}

func QuickSort(arr []int) {
	// Берем некий индекс в качестве опорной точки (Взял последний чтоб не передвигать в конец)
	pivot := len(arr) - 1
	// Если мы имеем меньше 2‑х элементов, то выходим из рекурсии
	if pivot < 1 {
		return
	}
	// left будет в качестве стартового индекса
	left := 0
	// Все что меньше pivot будем класть в left
	// и переходить к следующему индексу, и так не доходя до pivot
	for i, _ := range arr {
		if arr[i] < arr[pivot] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	// перемещаем pivot на свое законное место на следующую ячейку от left
	arr[left], arr[pivot] = arr[pivot], arr[left]
	// Разделяем на рекурсию, на левую и правую часть, не трогая часть pivot
	QuickSort(arr[left+1:])
	QuickSort(arr[:left])
}
