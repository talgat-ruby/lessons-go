package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func(r *rand.Rand) {
		time.Sleep(time.Duration(r.Intn(500)+500) * time.Millisecond)
		fmt.Printf("First player buzzes\n")
		channel1 <- "Player 1 Buzzed"
	}(r)

	go func(r *rand.Rand) {
		time.Sleep(time.Duration(r.Intn(500)+500) * time.Millisecond)
		fmt.Printf("Second player buzzes\n")
		channel2 <- "Player 2 Buzzed"
	}(r)

	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
}
