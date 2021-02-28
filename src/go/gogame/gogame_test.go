package gogame_test

import (
	"fmt"
	"reflect"
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
	red := "\033[31m"
	blue := "\033[34m"
	yellow := "\033[33m"
	resetColor := "\033[0m"

	for i := 0; i < boardSize; i++ {
		prettyBoard += fmt.Sprintf("%v%v%v ", yellow, i, resetColor)
	}

	for rowIndex, row := range board {

		prettyBoard += fmt.Sprintf("\n%v%v%v [", yellow, rowIndex, resetColor)

		for colIndex, intersection := range row {

			color := ""
			switch intersection {
			case BlackStone:
				color = red
			case WhiteStone:
				color = blue
			}

			if colIndex == len(row)-1 {
				prettyBoard += fmt.Sprintf("%v%v%v", color, intersection, resetColor)
			} else {
				prettyBoard += fmt.Sprintf("%v%v%v ", color, intersection, resetColor)
			}

		}

		prettyBoard += "]"
	}

	return prettyBoard
}

func TestGetNumLiberties(t *testing.T) {

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
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 1, 2, 1, 2, 1, 0, 0},
				{0, 0, 1, 2, 2, 2, 1, 0, 0},
				{0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(4, 4),
			5,
		},
	}

	for _, test := range numLibertiesTests {
		goGame = gogame.GoGame{Board: test.board, BoardSize: 9}
		ans = goGame.GetNumLiberties(test.target.X, test.target.Y, true)
		if ans != test.expected {
			t.Errorf("\nGiven Board: \n%v \nTarget: %v\nAnswer: %v\nExpected: %v",
				getPrettyBoard(test.board),
				test.target,
				ans,
				test.expected,
			)
		}
	}
}

func TestAttemptCapture(t *testing.T) {

	var goGame gogame.GoGame
	var ans bool

	var attemptCaptureTests = []struct {
		board         [][]uint8
		target        gogame.Coord
		expected      bool
		expectedBoard [][]uint8
	}{
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 2, 1, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(2, 4),
			true,
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 2, 1, 0, 0, 0, 0},
				{0, 0, 1, 2, 1, 0, 0, 0, 0},
				{0, 0, 1, 2, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(2, 4),
			false,
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 2, 1, 0, 0, 0, 0},
				{0, 0, 1, 2, 1, 0, 0, 0, 0},
				{0, 0, 1, 2, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 1, 2, 1},
				{0, 0, 0, 0, 0, 0, 1, 2, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 2},
				{0, 0, 0, 0, 0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			gogame.NewCoord(8, 1),
			true,
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0, 0},
				{0, 0, 1, 2, 2, 1, 0, 0, 0},
				{0, 0, 1, 2, 0, 2, 1, 0, 0},
				{0, 0, 0, 1, 2, 2, 1, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0, 0},
			},
			gogame.NewCoord(4, 6),
			true,
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 1, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 1, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0, 0},
			},
		},
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 1, 2, 2, 1, 0, 0, 0},
				{0, 0, 1, 2, 0, 2, 1, 0, 0},
				{0, 0, 0, 1, 2, 2, 1, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0, 0},
			},
			gogame.NewCoord(4, 6),
			true,
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 1, 2, 2, 1, 0, 0, 0},
				{0, 0, 1, 2, 0, 0, 1, 0, 0},
				{0, 0, 0, 1, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0, 0},
			},
		},
	}

	for _, test := range attemptCaptureTests {
		goGame = gogame.GoGame{Board: test.board, BoardSize: 9}
		ans = goGame.AttemptCapture(test.target.X, test.target.Y, true)
		boardIsCorrect := reflect.DeepEqual(goGame.Board, test.expectedBoard)
		if ans != test.expected || !boardIsCorrect {
			t.Errorf("\nGiven Board: \n%v \nTarget: %v\nAnswer: %v\nExpected: %v\nFinal Board: \n%v \nExpected Board: \n%v",
				getPrettyBoard(test.board),
				test.target,
				ans,
				test.expected,
				getPrettyBoard(goGame.Board),
				getPrettyBoard(test.expectedBoard),
			)
		}
	}
}
