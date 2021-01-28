// +build js,wasm

package gogame_test

import (
	"syscall/js"
	"testing"

	"example.com/gowithgo/gogame"
)

var board = js.Global().Call("eval", `
	[
		[0,0,0,0,0],
		[0,0,0,0,0],
		[0,0,0,0,0],
		[0,0,0,0,0],
		[0,0,0,0,0],
	]
`)

func TestGetNumLiberties(t *testing.T) {
	ans := gogame.GetNumLiberties(0, 0, board, true)

	if ans != 2 {
		t.Errorf("getNumLiberties(0, 0, *EmptyBoard*, true) = %d; want 2", ans)
	}
}
