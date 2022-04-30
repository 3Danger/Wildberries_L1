package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать структуру-счетчик,
	которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

type CounterStruct struct {
	sync.WaitGroup
	sync.Mutex
	Count int
}

func (c *CounterStruct) Increment() {
	c.Lock()
	// Критическая секция к которой
	// доступ в один момент времени
	// должен быть только у одного потока
	c.Count++
	c.Unlock()
}

func GoroutineOne(counterStruct *CounterStruct) {
	// Инкрементируем
	counterStruct.Increment()
	// Отмечаемся что завершили работу
	counterStruct.Done()
}

func usingMutex(n int) {
	counter := &CounterStruct{sync.WaitGroup{}, sync.Mutex{}, 0}
	counter.Add(n)
	for i := 0; i < n; i++ {
		go GoroutineOne(counter)
	}
	counter.Wait()
	fmt.Println("counted:", counter.Count)

}

func main() {
	usingMutex(1000000)
}
