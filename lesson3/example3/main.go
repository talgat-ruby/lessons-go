package main

import (
	"fmt"
)

func main() {
	fibNums := []int{1, 1, 2, 3, 5, 8}
	for _, v := range fibNums {
		v += v
		fmt.Println(v)
	}
	fmt.Println(fibNums)
}
