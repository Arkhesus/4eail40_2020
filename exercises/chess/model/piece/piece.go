package piece

import (
	"fmt"
)

//TODO Implement game pieces here

type Piece interface {
	fmt.Stringer
	Col Color.Code
	Repr string

}

type MethodPiece struct {
	fmt.Stringer
	Color() player.Color
	Move(isCapteur bool) map[coord.ChessCoordinates]bool
	GetInitPlace() []
	GetRepr() string
}

func (e *Piece) GetInitPlace() []int {
	return e.InitPlace
}

func (e *Piece) GetRepr() string {
	return e.Repr
}

func (e *Piece) String() string {
	return e.Col.Code + e.Repr + Reset.Code
}