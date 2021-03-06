package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg2     sync.WaitGroup
)

func main() {
	runtime.GOMAXPROCS(1)
	wg2.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg2.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter(id int) {
	defer wg2.Done()

	for count := 0; count < 2; count++ {
		value := counter

		runtime.Gosched()

		value++
		counter = value
	}
}
