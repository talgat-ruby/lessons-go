package main

import (
	"fmt"
)

//func main() {
//	var s stack.Stack[int]
//	s.Push(10)
//	s.Push(20)
//	s.Push(30)
//	fmt.Println(s.Contains(10))
//	fmt.Println(s.Contains(5))
//}

type num uint

func main() {
	var sum1 num = 104
	var sum2 int = 222
	minSum := Min(sum1, sum2) // error num does not satisfy uint | int | int64 (possibly missing ~ for uint in uint | int | int64)
	fmt.Println(minSum)       // 104
}

func Min[T ~uint | ~int | ~int64, P ~uint | ~int | ~int64](a T, b P) T {
	c := Convert[P, T](b)
	if a < c {
		return a
	}
	return c
}

func Convert[T ~uint | ~int | ~int64, P ~uint | ~int | ~int64](a T) P {
	return P(a)
}
