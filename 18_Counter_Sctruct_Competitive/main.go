package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	Реализовать структуру-счетчик,
	которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

type CounterStruct struct {
	sync.Mutex
	sync.WaitGroup
	Count int
}

func main() {
	const NWorkers = 10
	counter := CounterStruct{sync.Mutex{}, sync.WaitGroup{}, 0}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	counter.Add(NWorkers)
	for i := 0; i < NWorkers; i++ {
		go CompetitiveFunc(ctx, &counter)
	}
	counter.Wait()
	fmt.Println("counted:", counter.Count)
}

func CompetitiveFunc(ctx context.Context, c *CounterStruct) {
	for {
		select {
		case <-ctx.Done():
			c.Done()
			return
		default:
			c.Lock()
			c.Count++
			c.Unlock()
		}
	}
}
