package greet

import (
	"fmt"
)

func Greet(name string, table int) string {
	return fmt.Sprintf("Hello %s! Your table is %d.", name, table)
}
