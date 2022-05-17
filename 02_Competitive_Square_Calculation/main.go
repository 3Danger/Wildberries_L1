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
	SolutionChannel(array) // Решение с использованием каналов
	SolutionMutexWG(array) // решение с Mutex и WaitGroup
}

func SolutionMutexWG(nums []int) {
	mut := sync.Mutex{}
	// size: Запоминаем количество будущих горутин для Wait
	wg := sync.WaitGroup{}
	wg.Add(len(nums))
	fmt.Println("SolutionMutexWG")

	for _, v := range nums {
		go func(n int, m *sync.Mutex) {
			// Использую mutex для того
			// что бы была реальная конкуренция,
			// без mutex при печати результата
			// может случиться такое что несколько потоков
			// в один момент времени захотят вывести к примеру:
			// "123\n" и "456\n" в stdout, но вместо этого
			// может напечатать "1425\n3\n" c одновременным вызовом fmt.Println()
			// в то же время мы этим добились конкуренции.
			m.Lock()
			fmt.Println(n * n)
			m.Unlock()
			wg.Done()
		}(v, &mut)
	}
	wg.Wait()
}

func SolutionChannel(array []int) {
	ch := make(chan int, 0)
	fmt.Println("SolutionChannel")

	// Вычисление и запись
	for _, v := range array {
		go func(c chan<- int, v int) {
			// Каналы потоко-безопасны
			c <- v * v
		}(ch, v)
	}

	// Чтение из канала результатов
	for i := 0; i < len(array); i++ {
		fmt.Printf("%d\n", <-ch)
	}
	close(ch)
}
