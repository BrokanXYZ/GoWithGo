package gogame

import (
	"fmt"
	"strings"
)

// GoGame ...
type GoGame struct {
	BoardSize              int
	Board                  [][]uint8
	BlackPreviousBoardHash string
	WhitePreviousBoardHash string
	LastTurnWasPassed      bool
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

	newGoGame := GoGame{
		Board:                  newBoard,
		BoardSize:              boardSize,
		BlackPreviousBoardHash: "",
		WhitePreviousBoardHash: "",
		LastTurnWasPassed:      false,
	}

	return newGoGame, nil
}

// NewCoord creates a new coordinate object. Allows for easy thinking in (x,y) terms
func NewCoord(col int, row int) Coord {
	return Coord{X: row, Y: col}
}

func (goGame *GoGame) attemptCaptureGroup(x int, y int, isBlackTurn bool, capturedStones *map[Coord]bool) {
	toVisit := []Coord{{X: x, Y: y}}
	oppStone := WhiteStone

	if !isBlackTurn {
		oppStone = BlackStone
	}

	for len(toVisit) > 0 {
		current := toVisit[0]
		(*capturedStones)[current] = true
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

func (goGame *GoGame) getBoardHash() string {
	var hash strings.Builder
	hash.Grow(goGame.BoardSize * goGame.BoardSize)

	for _, row := range goGame.Board {
		for _, inter := range row {
			hash.WriteByte(inter)
		}
	}

	return hash.String()
}

// CaptureStones removes stones from specified coordinates
func (goGame *GoGame) CaptureStones(stonesToBeCaptured []Coord) {
	for _, coord := range stonesToBeCaptured {
		goGame.Board[coord.X][coord.Y] = NoStone
	}
}

// SetStone forcefully sets the target intersection on the board to the current
// player's stone.
func (goGame *GoGame) SetStone(x int, y int, isBlack bool) {
	if isBlack {
		goGame.Board[x][y] = BlackStone
	} else {
		goGame.Board[x][y] = WhiteStone
	}
}

// CheckForKo returns true if the rule of Ko has been violated at the target location.
// The move at the target location must have already been made.
func (goGame *GoGame) CheckForKo(x int, y int, isBlackTurn bool) bool {
	boardHash := goGame.getBoardHash()
	previousPlayerBoardHash := goGame.BlackPreviousBoardHash
	if !isBlackTurn {
		previousPlayerBoardHash = goGame.WhitePreviousBoardHash
	}

	if isBlackTurn {
		goGame.BlackPreviousBoardHash = boardHash
	} else {
		goGame.WhitePreviousBoardHash = boardHash
	}

	return boardHash == previousPlayerBoardHash
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

// AttemptCapture makes an attempt to capture stones as a result of placing a stone
// at the specified location. Ultimately, the board is not modified. The returned
// slice contains the stones that would be captured.
func (goGame *GoGame) AttemptCapture(x int, y int, isBlackTurn bool) []Coord {
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
				if goGame.GetNumLiberties(adjInter.X, adjInter.Y, !isBlackTurn) == 0 {
					goGame.attemptCaptureGroup(adjInter.X, adjInter.Y, isBlackTurn, &capturedStones)
				}
			}
		}
	}

	var result []Coord
	for k := range capturedStones {
		result = append(result, k)
	}

	return result
}

// TryPlaceStone attempts to place stone at given (x,y) Coords on the board.
// The stone will be successfully placed if it is a legal move. Otherwise,
// an error will be returned specifying the reason. If the move will capture
// stones, then they will be.
func (goGame *GoGame) TryPlaceStone(col int, row int, isBlackTurn bool) error {

	target := NewCoord(col, row)
	targetInter := &goGame.Board[target.X][target.Y]

	// 1. space is empty
	if *targetInter != NoStone {
		return fmt.Errorf("there is already a stone at position (%v,%v)", target.X, target.Y)
	}

	// 2. stone will have a liberty or would capture
	potentialLiberties := goGame.GetNumLiberties(target.X, target.Y, isBlackTurn)
	stonesToBeCaptured := goGame.AttemptCapture(target.X, target.Y, isBlackTurn)

	if potentialLiberties == 0 && len(stonesToBeCaptured) == 0 {
		return fmt.Errorf("stone would have no liberties and not capture")
	}

	// set stone and capture
	goGame.SetStone(target.X, target.Y, isBlackTurn)
	goGame.CaptureStones(stonesToBeCaptured)

	// 3.Ko:  one may not play in such a way as to recreate the board
	// 		  position following one's previous move.
	koViolated := goGame.CheckForKo(target.X, target.Y, isBlackTurn)
	if koViolated {
		// revert board changes
		*targetInter = NoStone

		oppStone := WhiteStone
		if !isBlackTurn {
			oppStone = BlackStone
		}

		for _, coord := range stonesToBeCaptured {
			goGame.Board[coord.X][coord.Y] = uint8(oppStone)
		}

		return fmt.Errorf("Ko: one may not play in such a way as to recreate the board position following one's previous move")
	}

	// Stone was successfully placed. Turn was not passed.
	goGame.LastTurnWasPassed = false

	return nil
}

// PassTurn passes the current player's turn.
// If the other player passed their previous turn, then this will trigger the end of the game.
// Returned: (isGameOver, playerOneScore, playerTwoScore)
func (goGame *GoGame) PassTurn(isBlackTurn bool) (bool, int, int) {
	isGameOver := false
	playerOneScore, playerTwoScore := 0, 0

	if goGame.LastTurnWasPassed {
		isGameOver = true
		// TODO: Score board
	} else {
		goGame.LastTurnWasPassed = true
	}

	return isGameOver, playerOneScore, playerTwoScore
}
