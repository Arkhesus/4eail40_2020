package state

import (
	"fmt"
	"../../model/board"
)

//TODO Implement game state related elements here

type IState interface {
	fmt.Stringer
	ShowState() string
}

type State8x8 struct {
	CurrentBoard  IBoard
	PreviousState *IState
	Player        string
	LastMove      []string
	ActionNumber  int
}

func (state State8x8) ShowState() string {
	return "This is the ne state"
}

func (state State8x8) String() string {
	return fmt.Sprintf("**********\nAction: %d \nPlayer: %s\nMove: %s\nGame: \n \n%s", state.ActionNumber, state.Player, state.LastMove, state.CurrentBoard.String())
}
