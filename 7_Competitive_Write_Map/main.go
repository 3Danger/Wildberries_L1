package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.
*/

type MapMutex struct {
	sync.Mutex
	data map[int]int
}

func main() {
	const COUNT = 1000
	dataMap := &MapMutex{sync.Mutex{}, make(map[int]int)}
	group := sync.WaitGroup{}
	group.Add(COUNT)

	// Конкурентная запись
	for i := 0; i < COUNT; i++ {
		go func(v int) {
			dataMap.Lock()
			dataMap.data[v] = v
			dataMap.Unlock()
			group.Done()
		}(i)
	}

	//Ждем завершения всех потоков
	group.Wait()

	// Чтение
	for k, v := range dataMap.data {
		fmt.Printf("key %d\tval %d\n", k, v)
	}
	fmt.Println("\nsize of map", len(dataMap.data))
}
