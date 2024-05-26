package main

import (
	"fmt"

	"github.com/talgat-ruby/lessons-go/lesson3/stack"
)

func main() {
	var s stack.Stack[int]
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Contains(10))
	fmt.Println(s.Contains(5))
}
