package board

import (
	"testing"

	"../coord"
)

type mockCoord int;

// Coord returns x if n==0, y if n==1
func (c mockCoord) Coord(n int) (int, error) {
	return int(c), nil;
}

func (c mockCoord) String() string {
	return "1";
}


func TestBoard8x8_PieceAt(t *testing.T){
	board := Board8x8{};
	coord := coord.Newcartesian(0,0)
	type args struct {
		at coord.ChessCoordinates
	}
	tests := []struct {
		name string
		b *Board8x8
		args args
		want piece.Piece
	}{
		"test A1",
		board,
		args{at: coord},
		nil,

	}	
	for _, tt := range tests {
		t.Run(tt.at,func(t *testing.T){
			if got := tt.b.PieceAt(tt.args.at); !reflect.Deepequal(got, tt.want) {
				t.Errorf("PieceAt() = %v, wan %v", got, tt.want)
			}
		})
	}
}

func TestClassic_PlacePieceAt(t *testing.T) {
	b := Board8x8{}
	coordin := coord.NewCartesian(0, 0)
	x, _ := coord.Coord(0)
	y, _ := coord.Coord(1)
	class[x][y] = pi
	type args struct {
		piece  piece.Piece
		at coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		b       *Boad8x8
		args    args
		wantErr bool // true if there is an error, false if there is not.
	}{
		{
			"Test pass",
			class,
			args{piece: "_", at: coord.NewCartesian(4, 4)},
			false,
		},
		{
			"Test failed",
			class,
			args{p: "_", at: coord},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.PlacePieceAt(tt.args.piece, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("PlacePieceAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestBoard8x8_MovePiece(t *testing.T) {
	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		b       *Board8x8
		args    args
		wantErr bool
	}{
		{
			"testmock",
			&Classic{},
			args{mockCoord(0), mockCoord(1)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.b.MovePiece(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Board8x8.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}