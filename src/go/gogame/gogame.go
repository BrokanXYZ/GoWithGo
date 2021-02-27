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

// GetNumLiberties returns the number of liberties that a stone would or does have.
// It doesn't matter if the stone exists yet or not. The color of the stone is assumed
// by the param isBlackTurn.
func (goGame *GoGame) GetNumLiberties(x int, y int, isBlackTurn bool) int {

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

		adjIntersections := []Coord{
			{X: current.X, Y: current.Y - 1},
			{X: current.X, Y: current.Y + 1},
			{X: current.X - 1, Y: current.Y},
			{X: current.X + 1, Y: current.Y},
		}

		for _, adjInter := range adjIntersections {
			coordIsInBoardRange := adjInter.Y >= 0 &&
				adjInter.X >= 0 &&
				adjInter.Y < goGame.BoardSize &&
				adjInter.X < goGame.BoardSize

			if coordIsInBoardRange && !visited[adjInter] {
				inter := goGame.Board[adjInter.X][adjInter.Y]
				if inter == uint8(sameStone) {
					toVisit = append(toVisit, adjInter)
				} else if inter == NoStone {
					numLiberties++
				}
			}
		}
	}

	return numLiberties
}

func (goGame *GoGame) captureGroup(x int, y int, isBlackTurn bool, capturedStones *map[Coord]bool) {
	toVisit := []Coord{{X: x, Y: y}}
	oppStone := WhiteStone

	if !isBlackTurn {
		oppStone = BlackStone
	}

	for len(toVisit) > 0 {
		current := toVisit[0]
		(*capturedStones)[current] = true
		goGame.Board[current.X][current.Y] = NoStone
		toVisit = toVisit[1:]

		adjIntersections := []Coord{
			{X: current.X, Y: current.Y - 1},
			{X: current.X, Y: current.Y + 1},
			{X: current.X - 1, Y: current.Y},
			{X: current.X + 1, Y: current.Y},
		}

		for _, adjInter := range adjIntersections {
			coordIsInBoardRange := adjInter.Y >= 0 &&
				adjInter.X >= 0 &&
				adjInter.Y < goGame.BoardSize &&
				adjInter.X < goGame.BoardSize

			if coordIsInBoardRange && !(*capturedStones)[adjInter] {
				inter := goGame.Board[adjInter.X][adjInter.Y]
				if inter == uint8(oppStone) {
					toVisit = append(toVisit, adjInter)
				}
			}
		}
	}
}

// AttemptCapture ...
func (goGame *GoGame) AttemptCapture(x int, y int, isBlackTurn bool) bool {
	capturedStones := make(map[Coord]bool)
	sameStone := BlackStone
	oppStone := WhiteStone
	if !isBlackTurn {
		sameStone = WhiteStone
		oppStone = BlackStone
	}

	// Temporarily place stone to check for captures
	goGame.Board[x][y] = uint8(sameStone)
	defer func() { goGame.Board[x][y] = NoStone }()

	adjIntersections := []Coord{
		{X: x, Y: y - 1},
		{X: x, Y: y + 1},
		{X: x - 1, Y: y},
		{X: x + 1, Y: y},
	}

	for _, adjInter := range adjIntersections {
		coordsIsInBoardRange := adjInter.Y >= 0 &&
			adjInter.X >= 0 &&
			adjInter.Y < goGame.BoardSize &&
			adjInter.X < goGame.BoardSize

		if coordsIsInBoardRange && !capturedStones[adjInter] {
			stone := goGame.Board[adjInter.X][adjInter.Y]
			if stone == uint8(oppStone) {
				if goGame.GetNumLiberties(adjInter.X, adjInter.Y, isBlackTurn) == 0 {
					goGame.captureGroup(adjInter.X, adjInter.Y, isBlackTurn, &capturedStones)
				}
			}
		}
	}

	return len(capturedStones) > 0
}

// PlaceStone attempts to place stone at given (x,y) Coords on the board
func (goGame *GoGame) PlaceStone(col int, row int, isBlackTurn bool) error {

	target := NewCoord(col, row)
	intersection := &goGame.Board[target.X][target.Y]

	// 1. space is empty
	if *intersection != NoStone {
		return fmt.Errorf("there is already a stone at position (%v,%v)", target.X, target.Y)
	}

	// 2. stone will have a liberty or would capture
	potentialLiberties := goGame.GetNumLiberties(target.X, target.Y, isBlackTurn)
	doesCapture := goGame.AttemptCapture(target.X, target.Y, isBlackTurn)

	if potentialLiberties == 0 && !doesCapture {
		return fmt.Errorf("stone would have no liberties and not capture")
	}

	// 3. stone will not recreate former board
	//	TODO!

	// Place stone
	if isBlackTurn {
		*intersection = BlackStone
	} else {
		*intersection = WhiteStone
	}

	return nil
}
