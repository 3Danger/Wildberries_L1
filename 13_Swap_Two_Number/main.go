package main

import (
	"fmt"
	"sync/atomic"
)

/*
	Поменять местами два числа без создания временной переменной.
*/

func main() {
	var a, b int64 = 999, 121
	fmt.Printf("before: a = %d, b = %d\n", a, b)

	// Первый способ
	a, b = b, a
	fmt.Printf("after:  a = %d, b = %d\n", a, b)

	// Второй способ
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("after:  a = %d, b = %d\n", a, b)

	// Третий способ
	b = atomic.SwapInt64(&a, b)
	fmt.Printf("after:  a = %d, b = %d\n", a, b)

}
