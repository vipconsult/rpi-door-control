package rpiGpio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

var (
	gpioPins = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 23, 24, 25, 26, 27}
)

const sysfs string = "/sys/class/gpio/"
const sysfsGPIOenable string = sysfs + "export"
const sysfsGPIOdisable string = sysfs + "unexport"

// DefaultDelay  used as a default if not explicitly set
const DefaultDelay = 2

// DefaultPin used as a default if not explicitly set
const DefaultPin = "18"

//NewControl the constructor with some defaults
func NewControl() *Control {
	return &Control{
		Type:  "timer",
		Pin:   DefaultPin,
		Delay: DefaultDelay * time.Second,
	}
}

// Control holds all configuration
type Control struct {
	Type  string
	Pin   string
	Delay time.Duration
}

// SetType is the controller type setter
func (c *Control) SetType(d string) error {
	switch d {
	case "timer":
		return nil
	case "toggle":
		c.Type = d
	default:
		return errors.New("Invalid control type:" + d)
	}
	return nil
}

//SetPin the pin on gpio that willbe controlled
func (c *Control) SetPin(d string) error {
	for _, v := range gpioPins {
		if strconv.Itoa(v) == d {
			c.Pin = d
			return nil
		}

	}
	sort.Ints(gpioPins)
	return errors.New(fmt.Sprintf("Invalid GPIO pin number:%v, choose one of :%v", d, gpioPins))
}

// SetDelay delay between enable and disable timer
func (c *Control) SetDelay(d string) error {
	if t, err := time.ParseDuration(d); err == nil {
		c.Delay = t
		return nil
	}
	return errors.New(fmt.Sprintf("Invalid time delay format :%v (use 1ms, 1s, 1m, 1h)", d[0]))
}

func (c *Control) enablePin() error {
	// enable if not already enabled
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
		log.Printf("Output %v ready for work!", c.Pin)
	}
	return nil
}

func (c *Control) disablePin() {
	if _, err := os.Stat(sysfs + "gpio" + c.Pin); os.IsNotExist(err) {
		// it is already disabled so nothing else to do, bail out
		return
	}

	err := ioutil.WriteFile(sysfsGPIOdisable, []byte(c.Pin), 0644)
	if err != nil {
		log.Printf("Oops can't disable pin %v because %v", c.Pin, err)
	}
	log.Printf("Disabled pin %v", c.Pin)
}

// StartTimer enable and then disable a pin output using a set delay
func (c *Control) StartTimer(ch chan string) error {
	if err := c.enablePin(); err != nil {
		log.Printf("I couldn't enable pin %v, because %v", c.Pin, err)
		return err
	}
	if err := ioutil.WriteFile(sysfs+"gpio"+c.Pin+"/value", []byte("1"), 0644); err != nil {
		return err
	}
	r := fmt.Sprintf("Pin %v got 'HIGH' on drugs for %v seconds", c.Pin, c.Delay)
	log.Printf(r)
	if ch != nil {
		ch <- r
	}
	time.Sleep(c.Delay)
	if err := ioutil.WriteFile(sysfs+"gpio"+c.Pin+"/value", []byte("0"), 0644); err != nil {
		return err
	}
	log.Printf("pin %v is laid 'LOW'", c.Pin)
	return nil
}

// Toggle between high and low state
func (c *Control) Toggle(ch chan string) error {
	if err := c.enablePin(); err != nil {
		log.Printf("I couldn't enable pin %v, because %v", c.Pin, err)
	}

	d, err := ioutil.ReadFile(sysfs + "gpio" + c.Pin + "/value")
	if err != nil {
		log.Printf("Oh boy can't read the status of pin	%v becasue I don't have my glasses and %v", c.Pin, err)
	}

	if string(d) == "1\n" {
		if err := ioutil.WriteFile(sysfs+"gpio"+c.Pin+"/value", []byte("0"), 0644); err != nil {
			return err
		}
		r := fmt.Sprintf("pin %v just got 'LOW' on selfesteam", c.Pin)
		log.Printf(r)
		ch <- r
		return nil
	}
	if err := ioutil.WriteFile(sysfs+"gpio"+c.Pin+"/value", []byte("1"), 0644); err != nil {
		return err
	}
	r := fmt.Sprintf("pin %v just got 'HIGH' on drugs", c.Pin)
	log.Printf(r)
	ch <- r
	return nil
}
