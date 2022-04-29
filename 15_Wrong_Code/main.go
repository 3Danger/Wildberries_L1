package main

import (
	"fmt"
	"strings"
)

/*
	К каким негативным последствиям может привести
	данный фрагмент кода, и как это исправить?
	Приведите корректный пример реализации.

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func mainExample() {
	someFunc()
}
*/

/*
	В этом примере если мы используем лишь первые 100 элементов из выделенных 1024, то 924 элементов не будут удалены,
	и будут висеть в памяти, сборщик мусора удалит это лишь тогда - когда justString тоже выйдет из области видимости.
*/

func main() {
	var justString string

	v := createHugeString(1 << 10)
	// Решение: клонировать
	justString = strings.Clone(v[:100])
	fmt.Printf(justString)
}

func createHugeString(size int) string {
	return strings.Repeat("/", size)
}
