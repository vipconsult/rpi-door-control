package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/zserge/hid"
)

func loop(device hid.Device, addr string) {
	if err := device.Open(); err != nil {
		fmt.Println(os.Stderr, err.Error())
		return
	}

	defer device.Close()

	udpAddr, _ := net.ResolveUDPAddr("udp4", addr)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	for {
		dur, _ := time.ParseDuration("60000ms")
		buf, err := device.Read(-1, dur)

		if err == nil {
			line := string(buf[:])
			if len(line) > 0 {
				line = strings.Trim(line, "\x00\r\n#")
				log.Println(line)
				conn.Write([]byte(line))
			}
		} else {
			fmt.Println(os.Stderr, err.Error())
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s ip:port\n", os.Args[0])
		fmt.Println()
		return
	}

	hid.UsbWalk(func(device hid.Device) {
		loop(device, os.Args[1])
	})
}
