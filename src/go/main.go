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
	//boardSize := board.Index(0).Length()
	moveIsLegal := true

	// 1. space is empty
	if target != None {
		moveIsLegal = false
		result["error"] = "There is already a stone there!"
	}

	// 2. stone will capture
	if moveIsLegal {

	}

	// 3. stone will have a liberty

	// WIP if adj intersection is same color as stone being placed,
	// then must check if connection has a liberty

	if moveIsLegal {
		var numLiberties int
		var adjIntersections []int

		if !board.Index(y + 1).IsUndefined() {
			adjIntersections = append(adjIntersections, board.Index(y+1).Index(x).Int())
		}

		if !board.Index(y - 1).IsUndefined() {
			adjIntersections = append(adjIntersections, board.Index(y-1).Index(x).Int())
		}

		if !board.Index(y).Index(x + 1).IsUndefined() {
			adjIntersections = append(adjIntersections, board.Index(y).Index(x+1).Int())
		}

		if !board.Index(y).Index(x - 1).IsUndefined() {
			adjIntersections = append(adjIntersections, board.Index(y).Index(x-1).Int())
		}

		for _, intersection := range adjIntersections {
			if intersection == None {
				numLiberties++
			}
		}

		if numLiberties == 0 {
			moveIsLegal = false
			result["error"] = "Stone would not have any liberties"
		}
	}
	// 4. stone will not recreate former board
	if moveIsLegal {

	}

	if moveIsLegal {
		if isBlackTurn {
			board.Index(y).SetIndex(x, BlackStone)
		} else {
			board.Index(y).SetIndex(x, WhiteStone)
		}
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
