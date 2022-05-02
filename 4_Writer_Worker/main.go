package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные
	данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
*/

func main() {
	var n int
	fmt.Println("Enter number of workers")
	_, ok := fmt.Scan(&n)
	if ok != nil || n <= 0 {
		log.Fatalln("Wrong input")
	}
	ch := make(chan int)

	ctx, cancelFunc := context.WithCancel(context.Background())
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		cancelFunc()
	}()

	for i := 0; i < n; i++ {
		go ReaderWorker(ctx, i, ch)
	}
	WriterFunc(ctx, ch)
	close(ch)
	fmt.Println(">>> SIGINT Handled! >>>")
}

func WriterFunc(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.Int()
		}
	}
}

func ReaderWorker(ctx context.Context, id int, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-ch:
			fmt.Println("id:", id, " Data:", v)
		}
	}
}
