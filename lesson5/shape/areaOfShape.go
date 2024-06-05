package shape

import (
	"fmt"
)

func areaOfShape(shaper Areaable) {
	fmt.Println(shaper.Area())
}
