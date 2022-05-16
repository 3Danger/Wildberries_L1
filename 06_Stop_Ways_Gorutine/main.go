package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

func routineContext(ctx context.Context, wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	<-ctx.Done()
	fmt.Println("\rClosed", msg)
	//for {
	//	select {
	/*	case <-ctx.Done():*/
	//		return
	//	default:
	//	........
	//	}
	//}
}

func routineChannel(s <-chan struct{}, wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	<-s
	fmt.Println("\rClosed", msg)
	//for {
	//	select {
	/*	case <-s: */
	//		return
	//	default:
	//	........
	//	}
	//}
}

func ExampleWithCancel() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancelFunc := context.WithCancel(context.Background())
	go routineContext(ctx, &wg, "1 - WithCancel")
	time.Sleep(time.Second)
	cancelFunc()
	wg.Wait()
}

func ExampleWithTimeout() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctxTime, cancelFuncTime := context.WithTimeout(context.Background(), time.Second)
	defer cancelFuncTime()
	go routineContext(ctxTime, &wg, "2 - WithTimeout")
	<-ctxTime.Done()
	wg.Wait()
}

func ExampleWithDeadline() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	deadline, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancelFunc()
	go routineContext(deadline, &wg, "3 - WithDeadline")
	<-deadline.Done()
	wg.Wait()
}

func ExampleWithCtxSignal() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("\tpress CTRL+C to close NotifyContext goroutine")
	deadline, cancelFunc := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancelFunc()
	go routineContext(deadline, &wg, "4 with context signal CTRL+C - NotifyContext")
	<-deadline.Done()
	wg.Wait()
}

func ExampleWithChannel() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	cn := make(chan struct{})
	go routineChannel(cn, &wg, "5 - WithChannel")
	time.Sleep(time.Second)
	cn <- struct{}{}
	wg.Wait()
}

func ExampleWithChannelSignal() {
	fmt.Println("\tpress CTRL+C to close last goroutine")
	wg := sync.WaitGroup{}
	wg.Add(1)
	cn := make(chan struct{})
	sg := make(chan os.Signal)
	signal.Notify(sg, os.Interrupt)
	go routineChannel(cn, &wg, "6 with signal CTRL+C - WithChannelSignal")
	<-sg
	cn <- struct{}{}
	wg.Wait()
}

func main() {
	// С помощью функции cancel() который получили с контекста, мы имеем возможность сообщить горутинам о прекращении работы.
	ExampleWithCancel()
	// Контекст в котором через заданное время будет послан сигнал о прекращении работы горутинам
	ExampleWithTimeout()
	// Тоже самое что и выше - но работать будет до наступления времени дедлайна
	ExampleWithDeadline()
	// Тоже контекст, но привязанный к сигналам, а данном случае (ctrl + c)
	ExampleWithCtxSignal()
	// Альтернативные способы с помощью каналов ...
	ExampleWithChannel()
	// и сигналов с каналами
	ExampleWithChannelSignal() // Здесь ждет сигнала SIGINT (ctrl + c)
}
