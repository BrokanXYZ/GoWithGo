// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"example.com/gowithgo/gogame"
)

func placeStoneWrapper() js.Func {
	wrapperFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		x := args[0].Int()
		y := args[1].Int()
		board := args[2]
		isBlackTurn := args[3].Bool()
		result := gogame.PlaceStone(x, y, board, isBlackTurn)
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
