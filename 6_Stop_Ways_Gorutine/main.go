package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

func routineContext(ctx context.Context, msg string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(msg, rand.Int())
		}
	}
}

func routineChannel(s <-chan struct{}, msg string) {
	for {
		select {
		case <-s:
			return
		default:
			fmt.Println(msg, rand.Int())
		}
	}
}

func ExampleWithCancel() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go routineContext(ctx, "1 - WithCancel")
	time.Sleep(time.Second)
	cancelFunc()
}

func ExampleWithTimeout() {
	ctxTime, cancelFuncTime := context.WithTimeout(context.Background(), time.Second)
	defer cancelFuncTime()
	go routineContext(ctxTime, "2 - WithTimeout")
	<-ctxTime.Done()
}

func ExampleWithDeadline() {
	deadline, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancelFunc()
	go routineContext(deadline, "3 - WithDeadline")
	<-deadline.Done()
}

func ExampleWithChannel() {
	cn := make(chan struct{})
	go routineChannel(cn, "4 - WithChannel")
	time.Sleep(time.Second)
	cn <- struct{}{}
}

func ExampleWithChannelSignal() {
	cn := make(chan struct{})
	sg := make(chan os.Signal)
	signal.Notify(sg, os.Interrupt)
	go routineChannel(cn, "5 - WithChannelSignal")
	<-sg
	cn <- struct{}{}
	fmt.Println("-----Canceled with signal-----")
}

func main() {
	// С помощью функции cancel() который получили с контекста, мы имеем возможность сообщить горутинам о прекращении работы.
	ExampleWithCancel()
	// Контекст в котором через заданное время будет послан сигнал о прекращении работы горутинам
	ExampleWithTimeout()
	// Тоже самое что и выше - но работать будет до наступления времени дедлайна
	ExampleWithDeadline()
	// Альтернативные способы с помощью каналов ...
	ExampleWithChannel()
	// и сигналов с каналами
	ExampleWithChannelSignal() // Здесь ждет сигнала SIGINT (ctrl + c)
}
