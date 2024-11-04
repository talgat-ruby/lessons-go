package main

import (
	"fmt"
	"sync"
	"time"
)

// Exercise: Thread-Safe Data Cache with Expiration
//
// Goal: Implement a thread-safe cache that allows storing key-value pairs with expiration times.

const NumJobs = 5

var workerDelay = 500 * time.Millisecond

var keys = []string{
	"item1",
	"item2",
	"item3",
}

type Cache struct {
}

func worker(id int, cache *Cache, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, key := range keys {
		value, found := cache.Get(key)
		if found {
			fmt.Printf("Worker %d: Cache Hit %s -> %v\n", id, key, value)
		} else {
			fmt.Printf("Worker %d: Cache Miss %s expired\n", id, key)
		}
		time.Sleep(workerDelay)
	}
}

func main() {
	cache := Cache{}
	var wg sync.WaitGroup

	cache.Set(keys[0], "data1", 3*time.Second)
	cache.Set(keys[1], "data2", 1*time.Second)
	cache.Set(keys[2], "data3", 5*time.Second)

	for i := range NumJobs {
		wg.Add(1)
		go worker(i, &cache, &wg)
	}

	wg.Wait()
	fmt.Println("Cache operations completed.")
}
