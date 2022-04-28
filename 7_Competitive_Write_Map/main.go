package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать конкурентную запись данных в map.
*/

type MapMutex struct {
	sync.RWMutex
	data map[int]int
}

func main() {
	const COUNT = 1000
	dataMap := &MapMutex{sync.RWMutex{}, make(map[int]int)}
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
	group.Wait()
	group.Add(COUNT)

	// Конкурентное чтение
	for i := 0; i < COUNT; i++ {
		go func(v int) {
			dataMap.RLock()
			fmt.Printf("key %d\tval %d\n", v, dataMap.data[v])
			dataMap.RUnlock()
			group.Done()
		}(i)
	}
	group.Wait()
	fmt.Println("len", len(dataMap.data))
}
