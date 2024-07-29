package main

import (
	"fmt"
)

func main() {
	withContinue()
	withGoto()
}

func withContinue() {
	duo := []string{"SpongeBob", "Patrick"}
outer:
	for _, character := range duo {
		for i, r := range character {
			if r == 'B' {
				continue outer
			}
			fmt.Println(i, r, string(r))
		}
		fmt.Println(character)
	}
}

func withGoto() {
	a := 10
	b := 20
	goto skip
skip:
	c := 30
	fmt.Println(a, b, c)
	if c > a {
		goto inner
	}
inner:
	fmt.Println("a is less than a")
}
