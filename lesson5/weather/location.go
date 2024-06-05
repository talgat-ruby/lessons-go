package weather

import (
	"fmt"
)

type Location struct {
	longitude float64
	latitude  float64
}

func NewLocation(longitude, latitude float64) Location {
	return Location{longitude, latitude}
}

func (l *Location) Info() {
	fmt.Printf("current locations are: %f and %f\n", l.longitude, l.latitude)
}
