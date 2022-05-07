package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные
	данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.

	Программа должна завершаться по нажатию Ctrl+C.
	Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	c := make(chan int)

	N := flag.Int("n", 3, "number of workers")
	flag.Parse()

	wg := sync.WaitGroup{}
	wg.Add(*N)
	notifyContext, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	for i := 0; i < *N; i++ {
		go Worker(&wg, i, c)
	}
	Producer(notifyContext, &wg, c)
}

func Producer(ctx context.Context, wg *sync.WaitGroup, ch chan<- int) {
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\tSignal ctrl+c handled")
			close(ch)
			wg.Wait()
			fmt.Println("\tMain thread are closed")
			break loop
		default:
			data := rand.Int()
			ch <- data
		}
	}
}

func Worker(wg *sync.WaitGroup, id int, c chan int) {
	defer wg.Done()
	for data := range c {
		fmt.Println("Worker id:", id, ", data:", data)
		time.Sleep(time.Millisecond * 300)
	}
	fmt.Printf("Worker ID: %d, closed\n", id)
}
