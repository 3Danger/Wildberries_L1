package main

import (
	"fmt"
	"sync"
)

/*
	Написать программу, которая конкурентно рассчитает
	значение квадратов чисел взятых из массива (2,4,6,8,10)
	и выведет их квадраты в stdout.
*/

//nums := []int{2,4,6,8,10}

func main() {
	array := []int{2, 4, 6, 8, 10}
	SolutionMutex(array)
	SolutionChannel(array)
	SolutionWaitGroup(array)
}

func SolutionMutex(nums []int) {
	mut := sync.RWMutex{}
	size := len(nums)
	waitChan := make(chan struct{}, size)
	fmt.Println("SolutionMutex")

	for _, v := range nums {
		go func(n int, m *sync.RWMutex) {
			m.Lock()
			fmt.Println(n * n)
			m.Unlock()
			waitChan <- struct{}{}
		}(v, &mut)
	}
	for size > 0 {
		<-waitChan
		size--
	}
	close(waitChan)
}

func SolutionChannel(array []int) {
	ch := make(chan int)
	fmt.Println("SolutionChannel")

	for _, v := range array {
		go func(v int) {
			ch <- v * v
		}(v)
	}

	for i := 0; i < len(array); i++ {
		fmt.Printf("%d\n", <-ch)
	}
}

func SolutionWaitGroup(array []int) {
	wg := new(sync.WaitGroup)

	fmt.Println("SolutionWaitGroup")
	wg.Add(len(array))
	for _, v := range array {
		go func(v int) {
			fmt.Printf("%d\n", v*v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}
