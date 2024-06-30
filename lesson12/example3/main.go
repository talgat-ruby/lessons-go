package main

import "fmt"

func printTable(n int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", i, n, n*i)
	}

}

func main() {
	for number := 2; number <= 12; number++ {
		go printTable(number)
	}
}
