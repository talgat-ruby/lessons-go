package main

import (
	"fmt"

	"github.com/talgat-ruby/lessons-go/lesson7/env-vars/config"
)

func main() {
	conf := config.NewConfig()
	fmt.Printf("%+v\n", conf)
}
