package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&counter, 1)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
