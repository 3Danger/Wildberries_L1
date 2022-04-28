package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func InputScanTime() (tm time.Duration) {
	fmt.Println("Enter num deadline second")
	scan, err := fmt.Scan(&tm)
	if err != nil || scan <= 0 || tm <= 0 {
		log.Fatalln("Wrong argument")
	}
	tm *= time.Second
	return tm
}

func main() {
	tm := InputScanTime()
	ctxTimeout, cancelFuncTm := context.WithTimeout(context.Background(), tm)
	defer cancelFuncTm()
	ctxClosable, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	ch := make(chan int)
	go Receiver(ctxClosable, ch)
	Sender(ctxTimeout, cancelFunc, ch)
	fmt.Println("Time expired", tm)
}

func Receiver(ctx context.Context, data <-chan int) {
	for {
		select {
		case d := <-data:
			fmt.Println(d)
		case <-ctx.Done():
			return
		}
	}
}

func Sender(ctx context.Context, cancelFunc context.CancelFunc, data chan<- int) {
	for {
		select {
		case <-ctx.Done():
			cancelFunc()
			close(data)
			return
		default:
			data <- rand.Int()
		}
	}
}
