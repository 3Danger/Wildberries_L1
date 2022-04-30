package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(22+32+42….)
	с использованием конкурентных вычислений.
*/

//nums := []int{2,4,6,8,10}

func main() {
	array := []int{2, 4, 6, 8, 10}
	SolutionChannel(array)
	SolutionWaitGroup(array)
	SolutionAtomic(array)
}

func SolutionChannel(array []int) {
	var sum int
	ch := make(chan int)
	fmt.Println("SolutionChannel")

	go func(ar []int) {
		defer close(ch)
		for _, v := range ar {
			ch <- v * v
		}
	}(array)

	for v := range ch {
		sum += v
	}
	fmt.Println("sum :=", sum)
}

func SolutionWaitGroup(array []int) {
	var sum int
	wg := new(sync.WaitGroup)

	fmt.Println("SolutionWaitGroup")
	wg.Add(len(array))
	for _, v := range array {
		go func(v int) {
			defer wg.Done()
			sum += v * v
		}(v)
	}
	wg.Wait()
	fmt.Println("sum :=", sum)
}

func SolutionAtomic(array []int) {
	var sum int64
	wg := sync.WaitGroup{}
	wg.Add(len(array))
	fmt.Println("SolutionAtomic")
	for _, v := range array {
		go func(vi int) {
			atomic.AddInt64(&sum, int64(vi*vi))
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println("sum :=", sum)
}
