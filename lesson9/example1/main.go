package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		inGoroutine := 1
		ch1 <- inGoroutine
		fromMain := <-ch2
		fmt.Println("goroutine:", inGoroutine, fromMain)
		close(ch1) // Clean up: close ch1
		close(ch2) // Clean up: close ch2
	}()

	var fromGoroutine int
	inMain := 2
	done := false
	for !done {
		select {
		case ch2 <- inMain:
			done = true
		case fromGoroutine = <-ch1:
			ch2 <- inMain
			done = true
		}
	}

	fmt.Println("main:", inMain, fromGoroutine)
}
