package main

import (
	"log"

	"github.com/krasi-georgiev/door/i2c"
)

func main() {
	i2c, err := i2c.NewI2C(0x70, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	i2c.Write([]byte{0x21, 0x00})
	i2c.Write([]byte{0x81, 0x00})

	for i := 0; i <= 0x0f; i++ {
		i2c.Write([]byte{byte(i), 0x00})
	}

	i2c.Write([]byte{0x00, 0x18})
	i2c.Write([]byte{0x02, 0x18})
	i2c.Write([]byte{0x04, 0x18})
	i2c.Write([]byte{0x06, 0x99})
	i2c.Write([]byte{0x08, 0xDB})
	i2c.Write([]byte{0x0a, 0x7E})
	i2c.Write([]byte{0x0c, 0x3C})
	i2c.Write([]byte{0x0e, 0x18})

}
