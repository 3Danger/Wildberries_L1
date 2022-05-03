package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные
	данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
*/

func main() {
	c := make(chan int)
	defer close(c)
	N := flag.Int("n", 3, "number of workers")
	flag.Parse()

	for i := 0; i < *N; i++ {
		go func(id int, c chan int) {
			for data := range c {
				time.Sleep(time.Second)
				fmt.Println("Worker id:", id, ", data:", data)
			}
		}(i, c)
	}

	for {
		data := rand.Int()
		c <- data
	}
}
