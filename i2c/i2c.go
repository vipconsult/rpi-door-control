package i2c

import (
	"fmt"
	"os"
	"syscall"
)

type I2C struct {
	rc *os.File
}

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

func (this *I2C) Write(buf []byte) (int, error) {
	return this.rc.Write(buf)
}

func (this *I2C) Close() error {
	return this.rc.Close()
}
