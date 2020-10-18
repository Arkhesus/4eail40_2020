package model

import (
	"fmt"
)

//TODO Implement game pieces here

type MethodPiece interface {
	fmt.Stringer
}

type Piece struct {
	Represent  string
	Directions [][]string
	Quantity   int
	Player     string
}

func (e Piece) String() string {
	return e.Represent
}
