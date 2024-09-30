package main

import (
	"fmt"

	"github.com/talgat-ruby/lessons-go/lesson7/example2/family"
)

func main() {
	simpsonsAssets := family.NewAssets(2, 1)
	simpsons := family.NewFamily("Simpsons", 5, simpsonsAssets)
	smiths := family.NewFamily("Smiths", 5, family.NewAssets(3, 1))
	fmt.Println(simpsons, smiths)
}
