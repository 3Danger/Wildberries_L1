package main

import (
	"context"
	"fmt"
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
	_, _ = fmt.Scan(&n)
	ch := make(chan int)

	ctx, cancelFunc := context.WithCancel(context.Background())
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		cancelFunc()
	}()

	for i := 0; i < n; i++ {
		go Worker(ctx, i, ch)
	}
	if n > 0 {
		WriterFunc(ctx, ch)
		fmt.Println(">>>>>>>>>>>>>>> SIGINT Handled! >>>>>>>>>>>>>>>")
	} else {
		fmt.Println("wrong input")
	}
	close(ch)
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

func Worker(ctx context.Context, id int, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-ch:
			fmt.Println("id:", id, " Data:", v)
		}
	}
}
