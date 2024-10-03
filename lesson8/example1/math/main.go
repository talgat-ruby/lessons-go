// math/main.go
package math

import (
	"fmt"

	"github.com/talgat-ruby/lessons-go/lesson7/example1/physics"
)

func Add(a, b int) int {
	hello()
	p := &physics.Physic{
		Name: "Name",
	}
	fmt.Println(p)
	return a + b
}

type Math struct {
	Name string
}
