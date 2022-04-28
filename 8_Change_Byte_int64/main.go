package main

import (
	"errors"
	"fmt"
	"log"
)

/*
	Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var n int8
	var data int64
	fmt.Println("Enter N of byte: [1..64]")
	for {
		_, ok := fmt.Scan(&n)
		if ok != nil {
			log.Fatalln(ok)
		}

		ok = ChangeByte(&data, n)
		if ok != nil {
			log.Fatalln(ok)
		}
		fmt.Print("value: ", data, "\nbytes: ")
		PrintBytes(data)
	}
}

func ChangeByte(data *int64, n int8) (ok error) {
	if n > 64 || n < 0 {
		return errors.New("range out of size int64")
	}
	*data ^= 1 << (n - 1)
	return nil
}

func PrintBytes(data int64) {
	for i := 63; i >= 0; i-- {
		if (data & (1 << i)) == 0 {
			print("0")
		} else {
			print("1")
		}
	}
	print("\n")
}
