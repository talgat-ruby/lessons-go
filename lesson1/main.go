package main

import (
	"fmt"

	"github.com/talgat-ruby/lessons-go/lesson1/greet"
)

func main() {
	text := greet.Greet("Leila", 4)
	fmt.Println(text)
}
