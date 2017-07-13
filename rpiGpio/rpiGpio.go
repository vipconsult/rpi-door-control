package rpiGpio

import (
	"io/ioutil"
	"os"
	"time"
)

const sysfs string = "/sys/class/gpio/"
const sysfsGPIOenable string = sysfs + "export"
const sysfsGPIOdisable string = sysfs + "unexport"

type Control struct {
	Pin   string
	Delay time.Duration
}

func NewControl(pin string, delay time.Duration) *Control {
	return &Control{
		Pin:   pin,
		Delay: delay,
	}
}

func (c *Control) enablePin() error {
	if _, err := os.Stat(sysfs + "gpio" + c.Pin); os.IsNotExist(err) {
		if _, err := os.Stat(sysfsGPIOenable); os.IsNotExist(err) {
			return err
		}

		if err := ioutil.WriteFile(sysfsGPIOenable, []byte(c.Pin), 0644); err != nil {
			return err
		}
		if err := ioutil.WriteFile(sysfs+"gpio"+c.Pin+"/direction", []byte("out"), 0644); err != nil {
			return err
		}
	}
	return nil
}

func (c *Control) disablePin() {
	if _, err := os.Stat(sysfs + "gpio" + c.Pin); os.IsNotExist(err) {
		return
	}

	ioutil.WriteFile(sysfsGPIOdisable, []byte(c.Pin), 0644)
}

func (c *Control) Start() error {
	if err := c.enablePin(); err != nil {
		return err
	}
	if err := ioutil.WriteFile(sysfs+"gpio"+c.Pin+"/value", []byte("1"), 0644); err != nil {
		return err
	}
	time.Sleep(c.Delay)
	if err := ioutil.WriteFile(sysfs+"gpio"+c.Pin+"/value", []byte("0"), 0644); err != nil {
		return err
	}
	return nil
}
