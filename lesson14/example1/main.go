package main

import (
	"fmt"
	"time"
)

// Exercise: Producer-Consumer with Rate-Limited Processing
//
// Goal: Implement a producer-consumer pattern where:
//
// · The producer generates random numbers and sends them into a buffered channel.
// · The consumer processes each number from the channel but can only process a limited number of items per second.

var producerDelay = 200 * time.Millisecond
var consumerDelay = 500 * time.Millisecond
var jobsNum = 10

func producer() {
	ticker := time.NewTicker(producerDelay)
	defer ticker.Stop()

	for i := range jobsNum {
		<-ticker.C
		num := i + 1
		fmt.Printf("Produced and sent: %d\n", num)
	}
}

func consumer() {
	ticker := time.NewTicker(consumerDelay)
	defer ticker.Stop()

	// TODO: receive values from producer
	for num := range jobsNum {
		<-ticker.C
		fmt.Printf("Consumed: %d\n", num)
	}
}

func main() {
	go producer()

	go consumer()
}
