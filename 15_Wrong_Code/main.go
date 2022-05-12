package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
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

/*
	Мы сожем проверить это с помощью пакета reflect,
	что бы убедиться в каких случаях передается тот же самый указатель на байты,
	что может вызвать проблему в виде не используемого мусора, который висит в памяти.
*/

func main() {
	var justString string

	v := createHugeString(1 << 10)
	fmt.Println(unsafe.Pointer((*(*reflect.StringHeader)(unsafe.Pointer(&v))).Data))

	// Решение: клонировать
	justString = strings.Clone(v[:100])
	//justString = v[:100]
	fmt.Println(unsafe.Pointer((*(*reflect.StringHeader)(unsafe.Pointer(&justString))).Data))
	fmt.Printf(justString)
}

func createHugeString(size int) string {
	return strings.Repeat("/", size)
}
