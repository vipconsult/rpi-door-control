package main

import (
	"time"

	"github.com/krasi-georgiev/door/display"
)

func main() {

	display.Success()
	time.Sleep(1 * time.Second)
	display.Error()

}
