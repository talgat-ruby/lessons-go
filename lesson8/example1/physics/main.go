package physics

import (
	"fmt"

	"github.com/talgat-ruby/lessons-go/lesson7/example1/math"
	"github.com/talgat-ruby/lessons-go/lesson7/example1/physics/cool"
)

type Physic struct {
	Name string
}

func SecondLawOfNewton() int {
	m := &math.Math{
		Name: "Euler",
	}
	fmt.Println(m)
	return internal.Force(3, 4)
}
