package utils

import "fmt"

type Coordinate struct {
	X int
	Y int
}

func (coord Coordinate) String() string {
	return fmt.Sprintf("x: %d, y: %d", coord.X, coord.Y)
}
