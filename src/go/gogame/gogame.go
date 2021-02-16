package gogame

import (
	"fmt"
)

// GoGame ...
type GoGame struct {
	BoardSize int
	Board     [][]uint8
}

// Coord ...
type Coord struct {
	X int
	Y int
}

// Intersection Enum
const (
	NoStone = iota
	BlackStone
	WhiteStone
)

// NewGoGame creates new instance of GoGame
func NewGoGame(boardSize int) (GoGame, error) {

	if boardSize < 1 {
		return GoGame{}, fmt.Errorf("boardSize should be > 0")
	}

	newBoard := make([][]uint8, boardSize)
	for i := range newBoard {
		newBoard[i] = make([]uint8, boardSize)
	}

	return GoGame{Board: newBoard, BoardSize: boardSize}, nil
}

// NewCoord creates a new coordinate object. Allows for easy thinking in (x,y) terms
func NewCoord(col int, row int) Coord {
	return Coord{X: row, Y: col}
}

// GetPotentialNumLiberties returns the number of liberties that a stone would have
// if it were placed at the target location
func (goGame *GoGame) GetPotentialNumLiberties(x int, y int, isBlackTurn bool) int {

	numLiberties := 0
	visited := make(map[Coord]bool)
	toVisit := []Coord{{X: x, Y: y}}
	sameStone := BlackStone

	if !isBlackTurn {
		sameStone = WhiteStone
	}

	for len(toVisit) > 0 {
		current := toVisit[0]
		visited[current] = true
		toVisit = toVisit[1:]

		haveVisitedUp := visited[Coord{X: current.X, Y: current.Y - 1}]
		if current.Y-1 >= 0 && !haveVisitedUp {
			up := goGame.Board[current.X][current.Y-1]

			if up == uint8(sameStone) {
				toVisit = append(toVisit, Coord{X: current.X, Y: current.Y - 1})
			} else if up == NoStone {
				numLiberties++
			}
		}

		haveVisitedDown := visited[Coord{X: current.X, Y: current.Y + 1}]
		if current.Y+1 < goGame.BoardSize && !haveVisitedDown {
			down := goGame.Board[current.X][current.Y+1]

			if down == uint8(sameStone) {
				toVisit = append(toVisit, Coord{X: current.X, Y: current.Y + 1})
			} else if down == NoStone {
				numLiberties++
			}
		}

		haveVisitedLeft := visited[Coord{X: current.X - 1, Y: current.Y}]
		if current.X-1 >= 0 && !haveVisitedLeft {
			left := goGame.Board[current.X-1][current.Y]

			if left == uint8(sameStone) {
				toVisit = append(toVisit, Coord{X: current.X - 1, Y: current.Y})
			} else if left == NoStone {
				numLiberties++
			}
		}

		haveVisitedRight := visited[Coord{X: current.X + 1, Y: current.Y}]
		if current.X+1 < goGame.BoardSize && !haveVisitedRight {
			right := goGame.Board[current.X+1][current.Y]

			if right == uint8(sameStone) {
				toVisit = append(toVisit, Coord{X: current.X + 1, Y: current.Y})
			} else if right == NoStone {
				numLiberties++
			}
		}
	}

	return numLiberties
}

// PlaceStone attempts to place stone at given (x,y) Coords on the board
func (goGame *GoGame) PlaceStone(col int, row int, isBlackTurn bool) error {

	target := NewCoord(col, row)
	intersection := &goGame.Board[target.X][target.Y]

	// 1. space is empty
	if *intersection != NoStone {
		return fmt.Errorf("there is already a stone at position (%v,%v)", target.X, target.Y)
	}

	// 2. stone will capture

	// 3. stone will have a liberty
	potentialLiberties := goGame.GetPotentialNumLiberties(target.X, target.Y, isBlackTurn)
	if potentialLiberties == 0 {
		return fmt.Errorf("stone would have no liberties and not capture")
	}

	// 4. stone will not recreate former board

	if isBlackTurn {
		*intersection = BlackStone
	} else {
		*intersection = WhiteStone
	}

	return nil
}
