package jswrappers

import (
	"fmt"
	"syscall/js"

	"github.com/BrokanXYZ/GoWithGo/gogame"
)

var goGame gogame.GoGame

func newGoGameWrapper() js.Func {
	wrapperFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result := map[string]interface{}{
			"board": nil,
			"error": nil,
		}
		boardSize := args[0].Int()

		var err error
		goGame, err = gogame.NewGoGame(boardSize)

		if err != nil {
			result["error"] = err.Error()
		} else {
			initialBoardEvalString := fmt.Sprintf("new Array(%[1]v).fill(0).map(() => new Array(%[1]v).fill(0))", boardSize)
			result["board"] = js.Global().Call("eval", initialBoardEvalString)
		}

		return result
	})
	return wrapperFunc
}

func tryMoveWrapper() js.Func {
	wrapperFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		result := map[string]interface{}{
			"board": nil,
			"error": nil,
		}
		col := args[0].Int()
		row := args[1].Int()
		isBlackTurn := args[2].Bool()

		err := goGame.TryMove(col, row, isBlackTurn)

		if err != nil {
			result["error"] = err.Error()
		} else {
			result["board"] = getJsBoard(goGame.Board, goGame.BoardSize)
		}

		return result
	})
	return wrapperFunc
}

func getJsBoard(goBoard [][]uint8, boardSize int) js.Value {
	jsBoardEvalString := "["

	for i := 0; i < boardSize; i++ {
		jsBoardEvalString += "["
		for j := 0; j < boardSize; j++ {
			jsBoardEvalString += fmt.Sprintf("%v,", goBoard[i][j])
		}
		jsBoardEvalString += "],"
	}

	jsBoardEvalString += "]"

	return js.Global().Call("eval", jsBoardEvalString)
}

// RegisterCallbacks assigns all wrapper functions to the Javascript Global object
func RegisterCallbacks() {
	js.Global().Set("newGoGame", newGoGameWrapper())
	js.Global().Set("tryMove", tryMoveWrapper())
}
