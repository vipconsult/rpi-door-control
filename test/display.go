package main

import (
	"log"

	"github.com/krasi-georgiev/door/i2c"
)

// init
// clear
// ok
// fail

func main() {
	// Create new connection to I2C bus on 2 line with address 0x27
	i2c, err := i2c.NewI2C(0x70, 1)
	if err != nil {
		log.Fatal(err)
	}
	// Free I2C connection on exit
	defer i2c.Close()

	i2c.Write([]byte{0x01, 0x00})
	i2c.Write([]byte{0x81, 0x00})
	i2c.Write([]byte{0x03, 0xFF})
	// i2c.Write([]byte{0x03, 0xFF})
	// i2c.Write([]byte{0x03, 0xFF})
	// i2c.Write([]byte{0x03, 0xFF})
}
