// Package i2c provides low level control over the linux i2c bus.
//
// Before usage you should load the i2c-dev kernel module
//
//      sudo modprobe i2c-dev
//
// Each i2c bus can address 127 independent i2c devices, and most
// linux systems contain several buses.
package i2c

import (
	"fmt"
	"os"
	"syscall"
)

// I2C represents a connection to an i2c device.
type I2C struct {
	rc *os.File
}

// New opens a connection to an i2c device.
func NewI2C(addr uint8, bus int) (*I2C, error) {
	f, err := os.OpenFile(fmt.Sprintf("/dev/i2c-%d", bus), os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, f.Fd(), 0x0703, uintptr(addr), 0, 0, 0); err != 0 {
		return nil, err
	}
	this := &I2C{rc: f}
	return this, nil
}

// Write sends buf to the remote i2c device. The interpretation of
// the message is implementation dependant.
func (this *I2C) Write(buf []byte) (int, error) {
	return this.rc.Write(buf)
}

func (this *I2C) Close() error {
	return this.rc.Close()
}
