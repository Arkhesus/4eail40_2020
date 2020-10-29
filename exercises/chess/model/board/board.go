// Package model contains the gameplay logic for the game of chess
package board

import (
	"fmt"

	"../coord"
	"../piece"
)

type IBoard interface {
	fmt.Stringer
	Initialize()
	MovePiece(from, to coord.ChessCoordinates) error
	PieceAt(at coord.ChessCoordinates) piece.Piece
	PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error
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

func (b *Board8x8) PieceAt(at coord.ChessCoordinates) piece.Piece {
	x := from.Coord(0);
	y := from.Coord(1);

	if b[x][y] != nil {
		return b[x][y]
	}
	else {
		return nil
	}

}

func (b *Board8x8) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {
	x := from.Coord(0);
	y := from.Coord(1);
	piece := b[x][y];

	return PlacePieceAt(piece, to)
}

func (b *Board8x8) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	x := at.Coord(0);
	y := at.Coord(1);

	if b.PieceAt(at) == nil {
		b[x][y] = piece;
		return nil
	} else {
		return fm.Errorf("Destination occupied by other piece");
	}
}
