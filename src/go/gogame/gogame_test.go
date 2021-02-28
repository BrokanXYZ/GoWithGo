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

	var attemptCaptureTests = []struct {
		board    [][]uint8
		target   gogame.Coord
		expected int
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
			1,
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
			0,
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
			3,
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
			6,
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
			3,
		},
		{
			[][]uint8{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 1, 1, 1},
				{0, 0, 0, 0, 1, 2, 2, 2, 2},
				{0, 0, 0, 0, 1, 2, 0, 1, 1},
			},
			gogame.NewCoord(6, 8),
			5,
		},
	}

	for _, test := range attemptCaptureTests {
		goGame := gogame.GoGame{Board: test.board, BoardSize: 9}
		stonesToBeCaptured := goGame.AttemptCapture(test.target.X, test.target.Y, true)
		ans := len(stonesToBeCaptured)
		if ans != test.expected {
			t.Errorf("\nGiven Board: \n%v \nTarget: %v\nAnswer: %v\nExpected: %v\nCaptured: %v",
				getPrettyBoard(test.board),
				test.target,
				ans,
				test.expected,
				stonesToBeCaptured,
			)
		}
	}
}

func TestCheckForKo(t *testing.T) {

	var checkForKoTests = []struct {
		board      [][]uint8
		initMoves  []gogame.Coord
		targetMove gogame.Coord
		expected   bool
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
			[]gogame.Coord{
				gogame.NewCoord(1, 0),
				gogame.NewCoord(2, 0),
				gogame.NewCoord(0, 1),
				gogame.NewCoord(1, 1),
				gogame.NewCoord(1, 2),
				gogame.NewCoord(2, 2),
				gogame.NewCoord(2, 3),
				gogame.NewCoord(3, 1),
				gogame.NewCoord(2, 1),
			},
			gogame.NewCoord(1, 1),
			true,
		},
	}

	/*for _, test := range attemptCaptureTests {
		goGame := gogame.GoGame{Board: test.board, BoardSize: 9}
		stonesToBeCaptured := goGame.AttemptCapture(test.target.X, test.target.Y, true)
		ans := len(stonesToBeCaptured)
		if ans != test.expected {
			t.Errorf("\nGiven Board: \n%v \nTarget: %v\nAnswer: %v\nExpected: %v\nCaptured: %v",
				getPrettyBoard(test.board),
				test.target,
				ans,
				test.expected,
				stonesToBeCaptured,
			)
		}
	}*/
}
