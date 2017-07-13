package main

import (
	"time"

	"../display"
)

func main() {

	display.Success()
	time.Sleep(1 * time.Second)
	display.Error()

}
