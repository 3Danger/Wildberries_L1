package main

import (
	"fmt"
	"time"
)

/*
	Разработать конвейер чисел.
	Даны два канала:
		в первый пишутся числа (x) из массива,
		во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	now := time.Now()
	const SIZE = 100000
	numbers := make([]int, 0, SIZE)
	for i := 1; i <= SIZE; i++ {
		numbers = append(numbers, i)
	}
	chanX := make(chan int)
	go func() {
		for _, n := range numbers {
			chanX <- n
		}
		close(chanX)
	}()

	chanDoubledX := make(chan int)
	go func() {
		for {
			v, ok := <-chanX
			if ok {
				chanDoubledX <- v << 1
			} else {
				close(chanDoubledX)
				return
			}
		}
	}()

	for {
		dv, ok := <-chanDoubledX
		if ok {
			fmt.Println(dv)
		} else {
			break
		}
	}
	fmt.Println(time.Now().Sub(now))
}
