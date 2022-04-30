package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Разработать программу, которая переворачивает
	подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/

func Reverse(bt []rune) string {
	start, end := 0, len(bt)-1
	for start < end {
		bt[start], bt[end] = bt[end], bt[start]
		start++
		end--
	}
	return string(bt)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку:")
	for {
		scanner.Scan()
		reverse := Reverse([]rune(scanner.Text()))
		fmt.Println(reverse)
	}
}
