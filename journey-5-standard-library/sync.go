package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mutex sync.Mutex

func increment() {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
