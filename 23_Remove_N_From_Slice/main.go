package main

import (
	"errors"
	"fmt"
)

/*
	Удалить i-ый элемент из слайса.
*/

func main() {
	arr := make([]int, 20, 30)
	for i := 0; i < 20; i++ {
		arr[i] = i
	}

	fmt.Printf("before:\nlen: %d, cap: %d - %v\n\n", len(arr), cap(arr), arr)

	arr, _ = Remove(arr, 4)
	fmt.Printf("after Remove:\nlen: %d, cap: %d - %v\n\n", len(arr), cap(arr), arr)

	arr, _ = RemoveWithResize(arr, 6)
	fmt.Printf("after RemoveWithResize:\nlen: %d, cap: %d - %v\n\n", len(arr), cap(arr), arr)
}

func Remove[T any](arr []T, n int) ([]T, error) {
	l := len(arr) - 1
	if n < l && n >= 0 {
		return append(arr[:n], arr[n+1:]...), nil
	} else if n == l {
		return arr[:n], nil
	}
	return nil, errors.New("index out of range")
}

func RemoveWithResize[T any](arr []T, n int) ([]T, error) {
	arr, ok := Remove(arr, n)
	if ok != nil {
		return arr, ok
	}
	l := float32(len(arr))
	c := float32(cap(arr))
	if c/l > 1.333 {
		arr2 := make([]T, int(l))
		copy(arr2, arr)
		return arr2, nil
	}
	return arr, nil
}
