package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(22+32+42….)
	с использованием конкурентных вычислений.
*/

//nums := []int{2,4,6,8,10}

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	SolutionMutex(&array)
	SolutionChannel(&array)
	SolutionWaitGroup(&array)
	SolutionAtomic(&array)
}

func SolutionMutex(nums *[5]int) {
	var sum int
	mut := sync.Mutex{}
	fmt.Println("SolutionMutex")
	mut.Lock()
	for _, v := range *nums {
		go func(n int, m *sync.Mutex) {
			m.Lock()
			sum += n * n
			m.Unlock()
		}(v, &mut)
	}
	mut.Unlock()
	// Знаю что Sleep() это худший способ ждать потоки,
	// но это только для демонстрации способа с минимальным, читабельным кодом
	time.Sleep(time.Millisecond * 50)
	fmt.Println("sum :=", sum)
}

func SolutionChannel(array *[5]int) {
	var sum int
	ch := make(chan int)
	fmt.Println("SolutionChannel")

	go func(ar *[5]int) {
		defer close(ch)
		for _, v := range *ar {
			ch <- v * v
		}
	}(array)

	for v := range ch {
		sum += v
	}
	fmt.Println("sum :=", sum)
}

func SolutionWaitGroup(array *[5]int) {
	var sum int
	wg := new(sync.WaitGroup)

	fmt.Println("SolutionWaitGroup")
	wg.Add(len(*array))
	for _, v := range *array {
		go func(v int) {
			defer wg.Done()
			sum += v * v
		}(v)
	}
	wg.Wait()
	fmt.Println("sum :=", sum)
}

func SolutionAtomic(array *[5]int) {
	var sum int64
	wg := sync.WaitGroup{}
	wg.Add(len(*array))
	fmt.Println("SolutionAtomic")
	for _, v := range *array {
		go func(vi int) {
			atomic.AddInt64(&sum, int64(vi*vi))
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("sum :=", sum)
}
