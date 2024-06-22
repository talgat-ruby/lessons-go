package main

import "fmt"

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	for i := range countTo(10) {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
}
