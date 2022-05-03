package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func InputScanTime() (tm time.Duration) {
	fmt.Println("Enter num of second for deadline")
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

	m := sync.Mutex{}
	m.Lock()

	ch := make(chan int)
	go Receiver(&m, ch)
	Sender(ctxTimeout, ch)
	m.Lock()
	fmt.Println("Time expired", tm)
}

func Receiver(m *sync.Mutex, data <-chan int) {
	defer m.Unlock()
	for d := range data {
		fmt.Println(d)
	}
}

func Sender(ctx context.Context, data chan<- int) {
	for {
		select {
		case <-ctx.Done():
			close(data)
			return
		default:
			data <- rand.Int()
		}
	}
}
