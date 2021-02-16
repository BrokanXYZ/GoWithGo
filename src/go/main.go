package main

import (
	"fmt"

	"github.com/BrokanXYZ/GoWithGo/jswrappers"
)

func main() {
	fmt.Println("Go WASM is intialized!")
	jswrappers.RegisterCallbacks()
	<-make(chan bool)
}
