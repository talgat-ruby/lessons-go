package shape

import (
	"fmt"
)

func perimeterOfShape(shaper Perimeterable) {
	fmt.Println(shaper.Perimeter())
}
