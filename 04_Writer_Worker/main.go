package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные
	данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
*/

func main() {
	var nWorker int
	wg := sync.WaitGroup{}
	fmt.Println("Enter number of workers")
	_, ok := fmt.Scan(&nWorker)
	if ok != nil || nWorker <= 0 {
		log.Fatalln("Wrong input")
	}
	wg.Add(nWorker)
	ch := make(chan int)
	ctx, cancelFunc := context.WithCancel(context.Background())
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		cancelFunc()
	}()

	for i := 0; i < nWorker; i++ {
		go ReaderWorker(ctx, &wg, i, ch)
	}
	WriterFunc(ctx, &wg, ch)
	fmt.Println(">>> SIGINT Handled! >>>")
}

func WriterFunc(ctx context.Context, wg *sync.WaitGroup, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			close(ch)
			return
		default:
			ch <- rand.Int()
		}
	}
}

func ReaderWorker(ctx context.Context, wg *sync.WaitGroup, id int, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case v := <-ch:
			fmt.Println("id:", id, " Data:", v)
		}
	}
}
