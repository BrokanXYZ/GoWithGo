package main

import (
	"fmt"
	"syscall/js"
)

const (
	None = iota
	BlackStone
	WhiteStone
)

func placeStone(x int, y int, board js.Value, isBlackTurn bool) map[string]interface{} {
	result := map[string]interface{}{
		"updatedBoard": board,
		"error":        nil,
	}

	target := board.Index(y).Index(x).Int()

	if target == None {
		if isBlackTurn {
			board.Index(y).SetIndex(x, BlackStone)
		} else {
			board.Index(y).SetIndex(x, WhiteStone)
		}
	} else {
		result["error"] = "There is already a stone there!"
	}

	return result
}

func placeStoneWrapper() js.Func {
	wrapperFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		x := args[0].Int()
		y := args[1].Int()
		board := args[2]
		isBlackTurn := args[3].Bool()
		result := placeStone(x, y, board, isBlackTurn)
		return result
	})
	return wrapperFunc
}

func registerCallbacks() {
	js.Global().Set("placeStone", placeStoneWrapper())
}

func main() {
	fmt.Println("Go WASM is intialized!")
	registerCallbacks()
	<-make(chan bool)
}
