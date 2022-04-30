package main

import (
	"fmt"
	"time"
)

/*
	Реализовать собственную функцию sleep.
*/

func main() {
	now := time.Now()
	Sleep1(2 * time.Second)
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	Sleep2(2 * time.Second)
	fmt.Println(time.Now().Sub(now))
}

func Sleep1(tm time.Duration) {
	<-time.After(tm)
}

func Sleep2(tm time.Duration) {
	tick := time.Tick(tm / 100)
	for i := 0; i < 100; i++ {
		<-tick
	}
}
