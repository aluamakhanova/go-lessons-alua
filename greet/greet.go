package greet

import (
	"fmt"
	"go-lessons-alua/greet/input"
)

func Greet(input input.Input) string {
	return fmt.Sprintf("Hello, %s! Your table is %d.", input.Name, input.Table)
}
