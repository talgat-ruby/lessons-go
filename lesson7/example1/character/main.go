package character

import (
	"fmt"
)

type Character struct {
	Name       string
	Age        int
	Occupation string
}

func (c *Character) info() {
	fmt.Printf("name: %s, age: %d, occupation: %s\n", c.Name, c.Age, c.Occupation)
}
