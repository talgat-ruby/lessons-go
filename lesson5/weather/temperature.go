package weather

import (
	"fmt"
)

type Temperature struct {
	min, max int
}

func NewTemperature(min, max int) Temperature {
	return Temperature{min, max}
}

func (t *Temperature) Info() {
	fmt.Printf("the max and min temps are: %d and %d\n", t.max, t.min)
}
