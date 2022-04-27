package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Написать программу, которая конкурентно рассчитает
	значение квадратов чисел взятых из массива (2,4,6,8,10)
	и выведет их квадраты в stdout.
*/

//nums := []int{2,4,6,8,10}

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	SolutionMutex(&array)
	SolutionChannel(&array)
	SolutionWaitGroup(&array)
}

func SolutionMutex(nums *[5]int) {
	mut := sync.Mutex{}
	fmt.Println("SolutionMutex")
	mut.Lock()
	for _, v := range *nums {
		go func(n int, m *sync.Mutex) {
			m.Lock()
			fmt.Println(n * n)
			m.Unlock()
		}(v, &mut)
	}
	mut.Unlock()
	// Знаю что Sleep() это худший способ ждать потоки,
	// но это только для демонстрации способа с минимальным читабельным кодом
	time.Sleep(time.Millisecond * 50)
}

func SolutionChannel(array *[5]int) {
	ch := make(chan int)
	fmt.Println("SolutionChannel")

	for _, v := range *array {
		go func(v int) {
			ch <- v * v
		}(v)
	}

	for i := 0; i < len(*array); i++ {
		fmt.Printf("%d\n", <-ch)
	}
}

func SolutionWaitGroup(array *[5]int) {
	wg := new(sync.WaitGroup)

	fmt.Println("SolutionWaitGroup")
	wg.Add(len(*array))
	for _, v := range *array {
		go func(v int) {
			defer wg.Done()
			fmt.Printf("%d\n", v*v)
		}(v)
	}
	wg.Wait()
}
