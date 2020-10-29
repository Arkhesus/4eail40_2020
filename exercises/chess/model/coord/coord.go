package coord

import (
	"fmt"
)

type ChessCoordinates interface {
	fmt.Stringer
	Coord(n int) (int, error)
}
type Cartesian struct {
	x, y int
}

func NewCartesian(x, y int) Cartesian {
	return Cartesian{x, y}
}

func (c Cartesian) Coord(n int) (int, error) {
	switch n {
	case 0:
		return c.x, nil
	case 1:
		return c.y, nil
	}

	return 0, fmt.Errorf("unknown coor component %d", n)
}

func (c Cartesian) String() string {
	return fmt.Sprintf("%c%d", 65+c.y, c.x+1)
}
