// Package model contains the gameplay logic for the game of chess
package model

import (
	"fmt"
)

type Board interface {
	fmt.Stringer
	Initialize()
	Move(from, to []int) int
}

type Board8x8 struct {
	Table [][]Piece
}

func (b *Board8x8) Initialize() {
	for row := 0; row < 8; row++ {
		line := make([]Piece, 0, 8) //crée une liste de type Pièce rempli de 0 taille 8
		for col := 0; col < 8; col++ {
			line = append(line, Piece{Represent: "_"})
		}
		b.Table = append(b.Table, line)
	}
}

func (b *Board8x8) String() string {
	toprint := ""
	for row := range b.Table {
		if row == 0 {
			toprint += "| |A|B|C|D|E|F|G|H\n"
		}
		for col := range b.Table[row] {
			if col == 0 {
				toprint += fmt.Sprintf("|%d|", row+1)
			}
			toprint += fmt.Sprintf(b.Table[row][col].String())
			toprint += "|"
		}
		toprint += "\n"
	}

	return toprint
}
