package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func Reverse2(str string) string {
	texts := strings.Split(str, " ")
	start, end := 0, len(texts)-1
	swapper := reflect.Swapper(texts)
	for ; start < end; start, end = start+1, end-1 {
		swapper(start, end)
	}
	return strings.Join(texts, " ")
}

func Reverse(str string) string {
	texts := strings.Split(str, " ")
	start, end := 0, len(texts)-1
	for ; start < end; start, end = start+1, end-1 {
		texts[start], texts[end] = texts[end], texts[start]
	}
	return strings.Join(texts, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку:")
	for {
		scanner.Scan()
		reverse := Reverse(scanner.Text())
		fmt.Println(reverse)
	}
}
