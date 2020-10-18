package model

import (
	"fmt"
)

//TODO Implement game state related elements here

type State interface {
	fmt.Stringer
	ShowState() string
}

type State8x8 struct {
	CurrentBoard  Board8x8
	PreviousState *State8x8
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
