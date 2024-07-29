package main

import (
	"fmt"
)

func main() {
	fibNums := []int{1, 1, 2, 3, 5, 8}
	for i, v := range fibNums {
		fmt.Println(i, v)
	}

	simpsons := map[string]string{
		"father":    "Homer",
		"mother":    "Marge",
		"son":       "Bart",
		"daughter1": "Liza",
		"daughter2": "Maggie",
	}
	for role, name := range simpsons {
		fmt.Println(role, name)
	}

	samples := []string{"Batman", "J🤪ker", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
	}

	for i := range 10 {
		fmt.Println(i) // 0 1 2 3 4 5 6 7 8 9
	}
}
