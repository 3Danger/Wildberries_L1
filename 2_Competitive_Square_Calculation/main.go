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
	SolutionChannel(array)
	SolutionWaitGroup(array)
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
