package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func Reverse(texts []string) string {
	start, end := 0, len(texts)-1
	for start < end {
		texts[start], texts[end] = texts[end], texts[start]
		start++
		end--
	}
	return strings.Join(texts, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку:")
	for {
		scanner.Scan()
		reverse := Reverse(strings.Split(scanner.Text(), " "))
		fmt.Println(reverse)
	}
}
