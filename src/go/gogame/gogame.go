// +build js,wasm

package gogame

import (
	"fmt"
	"syscall/js"
)

// GoGame ...
type GoGame struct {
	BoardSize int
	Board     [][]uint8
}

// Intersection Enum
const (
	None = iota
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

// GetNumLiberties ...
func GetNumLiberties(x int, y int, board js.Value, isBlackTurn bool) int {
	return 0
}

// PlaceStone attempts to place stone at given (x,y) coords on the board
func (goGame *GoGame) PlaceStone(x int, y int, isBlackTurn bool) error {
	// result := map[string]interface{}{
	// 	"updatedBoard": board,
	// 	"error":        nil,
	// }

	target := goGame.Board[y][x]

	// 1. space is empty
	if target != None {
		return fmt.Errorf("there is already a stone at position (%v,%v)", x, y)
	}

	// 2. stone will capture

	// 3. stone will have a liberty

	// WIP if adj intersection is same color as stone being placed,
	// then must check if connection has a liberty

	// if moveIsLegal {
	// 	var numLiberties int
	// 	var adjIntersections []int

	// 	if !board.Index(y + 1).IsUndefined() {
	// 		adjIntersections = append(adjIntersections, board.Index(y+1).Index(x).Int())
	// 	}

	// 	if !board.Index(y - 1).IsUndefined() {
	// 		adjIntersections = append(adjIntersections, board.Index(y-1).Index(x).Int())
	// 	}

	// 	if !board.Index(y).Index(x + 1).IsUndefined() {
	// 		adjIntersections = append(adjIntersections, board.Index(y).Index(x+1).Int())
	// 	}

	// 	if !board.Index(y).Index(x - 1).IsUndefined() {
	// 		adjIntersections = append(adjIntersections, board.Index(y).Index(x-1).Int())
	// 	}

	// 	for _, intersection := range adjIntersections {
	// 		if intersection == None {
	// 			numLiberties++
	// 		}
	// 	}

	// 	if numLiberties == 0 {
	// 		moveIsLegal = false
	// 		result["error"] = "Stone would not have any liberties"
	// 	}
	// }

	// 4. stone will not recreate former board

	if isBlackTurn {
		goGame.Board[y][x] = BlackStone
	} else {
		goGame.Board[y][x] = WhiteStone
	}

	return nil
}
