package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет,
	что все символы в строке уникальные
	(true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.

	Например:
		abcd — true
		abCdefAaf — false
		aabcd — false
*/

func main() {
	input := "Введите что-ни-будь:"
	fmt.Println(input)
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			return
		}
		fmt.Println(input+" -", CheckUnique(&input))
	}
}

func CheckUnique(input *string) bool {
	data := []rune(strings.ToLower(*input))
	l := len(data)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if data[j] == data[i] {
				return false
			}
		}
	}
	return true
}
