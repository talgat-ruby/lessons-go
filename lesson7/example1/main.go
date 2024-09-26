package main

import (
	"fmt"
	"reflect"

	"github.com/talgat-ruby/lessons-go/lesson7/example1/character"
	"github.com/talgat-ruby/lessons-go/lesson7/example1/population"
)

type trademark struct {
	*character.Character
	*population.Population
	Name  string
	Owner string
}

type sampleStruct struct {
	Num1 *int   // 8
	Num2 *int   // 8
	Num3 *int   // 8
	Str1 string // 16
}

func main() {
	bender := &character.Character{
		Name:       "Bender Bending Rodríguez",
		Age:        4,
		Occupation: "worker",
	}

	_ = &trademark{
		Character: bender,
		Name:      "Bender Trademark",
		Owner:     "Disney",
	}

	a := 1
	b := &a

	s := sampleStruct{
		Num1: b,
		Num2: b,
		Num3: b,
		Str1: "",
	}

	fmt.Println(reflect.TypeOf(b).Size())
	fmt.Println(reflect.TypeOf("").Size())
	fmt.Println(reflect.TypeOf(s).Size())
	//benderTrademark.Population.info()
	//fmt.Println(benderTrademark.owner)
}
