package main

import (
	"fmt"
)

func main() {
	simpsons := []string{"Homer", "Marge", "Bart", "Liza", "Maggy"}
	parents := simpsons[0:2]
	children := simpsons[2:5]

	simpsons[2] = "Seymour Butz"

	parents = append(parents, "Abe")

	fmt.Println(simpsons, parents, children)
}
