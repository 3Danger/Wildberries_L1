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

func WriteX(chanX chan<- int, numbers []int) {
	for _, n := range numbers {
		chanX <- n
	}
	close(chanX)
}

func WriteDoubledX(chanDoubledX chan int, chanX <-chan int) {
	for {
		v, ok := <-chanX
		if ok {
			chanDoubledX <- v << 1
		} else {
			close(chanDoubledX)
			return
		}
	}
}

func main() {
	// Засекаем время
	now := time.Now()

	// Инициализируем слайс
	const SIZE = 333
	numbers := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		numbers[i] = i + 1
	}
	// Запускаем горутину для записи слайса в канал
	chanX := make(chan int)
	go WriteX(chanX, numbers)

	// Запускаем горутину для записи из канала в другой канал с 2x
	chanDoubledX := make(chan int)
	go WriteDoubledX(chanDoubledX, chanX)
	// В теории можно использовать один и тот же канал, но не стоит так делать мне кажется.

	//Читаем на результат
	for dv := range chanDoubledX {
		fmt.Println(dv)
	}
	// Смотрим сколько времени потребовалось:)
	fmt.Println("took time:", time.Now().Sub(now))
}
