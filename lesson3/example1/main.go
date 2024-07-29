package main

import (
	"fmt"
	"math/rand"
)

// NOTE shadow true

func main() {
	//slog.Info("Hello")
	//fmt.Println(slog)
	//slog := "oops"
	//fmt.Println(slog)

	//fmt.Println(true)
	//true := 42
	//fmt.Println(true)

	n := rand.Intn(10)

	if n < 5 {
		fmt.Println("That's too low:", n)
		return
	}

	if n > 5 {
		fmt.Println("That's too big:", n)
		return
	}

	fmt.Println("Perfect:", n)
}
