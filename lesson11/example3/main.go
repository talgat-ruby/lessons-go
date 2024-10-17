package main

import (
	"context"
	"fmt"
)

func countTo(ctx context.Context, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				fmt.Println("done")
				return
			}
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancelable, _ := context.WithCancel(ctx)
	defer cancel()
	for i := range countTo(cancelable, 10) {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
