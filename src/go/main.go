package main

import (
	"fmt"
	"syscall/js"
)

func placeStone(x int, y int, board js.Value) map[string]interface{} {
	fmt.Printf("(%v,%v) = %v\n", x, y, board.Index(x).Index(y))

	result := map[string]interface{}{
		"updatedBoard": board,
		"error":        nil,
	}

	if board.Index(y).Index(x).Int() == 0 {
		board.Index(y).SetIndex(x, 1)
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
		result := placeStone(x, y, board)
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
