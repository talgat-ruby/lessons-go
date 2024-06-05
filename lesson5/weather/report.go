package weather

import (
	"time"
)

type Report struct {
	Temperature
	Location
	createdAt time.Time
}

func NewReport(temperature Temperature, location Location, createdAt time.Time) Report {
	return Report{
		Temperature: temperature,
		Location:    location,
		createdAt:   createdAt,
	}
}

func (r *Report) Info() {
	r.Temperature.Info()
	r.Location.Info()
}
