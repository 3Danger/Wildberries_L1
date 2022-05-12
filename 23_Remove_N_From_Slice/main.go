package main

import (
	"errors"
	"fmt"
)

/*
	Удалить i-ый элемент из слайса.
*/

func main() {
	arr := make([]int, 10, 30)
	for i := range arr {
		arr[i] = i
	}

	fmt.Printf("len: %d, cap: %d - %v До изменений\n\n", len(arr), cap(arr), arr)

	arr, _ = RemoveWithPreservationOrder(arr, 7)
	fmt.Printf("len: %d, cap: %d - %v\tУдалени с сохранением порядка\n", len(arr), cap(arr), arr)

	arr, _ = Remove(arr, 4)
	fmt.Printf("len: %d, cap: %d - %v\tБыстрое удаление c нарушением порядка\n", len(arr), cap(arr), arr)

	arr, _ = RemoveAndResize(arr, 2)
	fmt.Printf("len: %d, cap: %d - %v\tУдаление и подгонка capacity\n", len(arr), cap(arr), arr)

	fmt.Printf("\nlen: %d, cap: %d - %v\tПосле изменений\n\n", len(arr), cap(arr), arr)
}

func Remove[T any](arr []T, n int) ([]T, error) {
	l := len(arr) - 1
	if n > l && n <= 0 || arr == nil {
		return nil, errors.New("index out of range")
	}
	arr[n], arr[len(arr)-1] = arr[len(arr)-1], arr[n]
	return arr[:len(arr)-1], nil
}

func RemoveWithPreservationOrder[T any](arr []T, n int) ([]T, error) {
	l := len(arr) - 1
	if n > l && n <= 0 || arr == nil {
		return nil, errors.New("index out of range")
	} else if n == l {
		return arr[:n], nil
	}
	return append(arr[:n], arr[n+1:]...), nil
}

func RemoveAndResize[T any](arr []T, n int) ([]T, error) {
	arr, ok := Remove(arr, n)
	if ok != nil {
		return arr, ok
	}
	return Resize(arr), nil
}

func Resize[T any](arr []T) []T {
	l := float32(len(arr))
	c := float32(cap(arr))
	if c/l > 1.333 {
		arr2 := make([]T, int(l))
		copy(arr2, arr)
		return arr2
	}
	return arr
}
