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
	//arr := []int{2, 4, 6, 8, 10}
	arr := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	SolutionMutex(arr)
	SolutionChannel(arr)
	SolutionAtomic(arr)
}

func SolutionMutex(nums []int) {
	var sum int
	mut := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(nums))
	fmt.Println("SolutionMutex")
	for _, v := range nums {
		go func(n int, m *sync.Mutex) {
			m.Lock()
			sum += n * n
			m.Unlock()
			wg.Done()
		}(v, &mut)
	}
	wg.Wait()
	fmt.Println("sum :=", sum)
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
