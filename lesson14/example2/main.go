package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Exercise: Dynamic Task Manager with Timeout
//
// Goal: Implement a "task manager" where:
//
// · A set of workers process tasks from a buffered task channel.
// · Each worker processes tasks independently, but a timeout will occur if no task is received within a certain time.
// · The task manager can dynamically receive or stop tasks based on input.

var jobsNum = 20
var workersNum = 3

var workerTimeout = 2 * time.Second
var producerTimeout = 1 * time.Second

func worker(id int, rnd *rand.Rand, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		processTime := time.Duration(rnd.Intn(2000)) * time.Millisecond
		fmt.Printf("Worker %d: Processing task %d %dms\n", id, task, processTime.Milliseconds())
		time.Sleep(processTime)
	}

	fmt.Printf("Worker %d: Task channel closed, exiting.\n", id)
}

func producer(rnd *rand.Rand, tasks chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range jobsNum {
		jobNum := i + 1
		tasks <- jobNum
		fmt.Printf("Producer: Sent a job %d\n", jobNum)
		processTime := time.Duration(rnd.Intn(500)) * time.Millisecond
		time.Sleep(processTime)
	}
	close(tasks)
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)

	tasks := make(chan int, 5)
	var wg sync.WaitGroup

	for i := range workersNum {
		wg.Add(1)
		go worker(i+1, rnd, tasks, &wg)
	}

	wg.Add(1)
	go producer(rnd, tasks, &wg)

	wg.Wait()

	fmt.Println("All workers have completed.")
}
