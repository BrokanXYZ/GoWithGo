package gogame_test

import (
	"fmt"
	"testing"

	"github.com/BrokanXYZ/GoWithGo/gogame"
)

// Intersection Enum
const (
	None = iota
	BlackStone
	WhiteStone
)

func getPrettyBoard(board [][]uint8) string {
	prettyBoard := "   "
	boardSize := len(board[0])

	for i := 0; i < boardSize; i++ {
		prettyBoard += fmt.Sprintf("%v ", i)
	}

	for rowIndex, row := range board {

		prettyBoard += fmt.Sprintf("\n%v [", rowIndex)

		for colIndex, intersection := range row {
			if colIndex == len(row)-1 {
				prettyBoard += fmt.Sprintf("%v", intersection)
			} else {
				prettyBoard += fmt.Sprintf("%v ", intersection)
			}

		}

		prettyBoard += "]"
	}

	return prettyBoard
}

func TestGetPotentialNumLiberties(t *testing.T) {

	var goGame gogame.GoGame
	var ans int

	var numLibertiesTests = []struct {
		board    [][]uint8
		target   gogame.Coord
		expected int
	}{
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(0, 0),
			2,
		},
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0, 0},
				{0, 0, 0, 2, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(4, 4),
			0,
		},
		{
			[][]uint8{
				{0, 1, 0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(0, 0),
			2,
		},
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 1, 2, 0, 2, 1, 0, 0},
				{0, 0, 1, 2, 2, 2, 1, 0, 0},
				{0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(4, 4),
			5,
		},
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 1, 1, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(1, 6),
			16,
		},
	}

	for _, test := range numLibertiesTests {
		goGame = gogame.GoGame{Board: test.board, BoardSize: 9}
		ans = goGame.GetPotentialNumLiberties(test.target.X, test.target.Y, true)
		if ans != test.expected {
			t.Errorf("\nBoard: \n%v \nTarget: %v\nAnswer: %v\nExpected: %v",
				getPrettyBoard(test.board),
				test.target,
				ans,
				test.expected,
			)
		}
	}
}
