package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
	//"strconv"
	"github.com/zserge/hid"
)

func loop(device hid.Device, addr string) {
	if err := device.Open(); err != nil {
		log.Println(os.Stderr, err.Error())
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
		buf, err := device.Read(-1, 10*time.Second)
		if (err == nil) {
			line := string(buf[:])
			//log.Println("LINE: " + line)
			//log.Println("LENGTH: " + strconv.Itoa(len(line)))
			if len(line) > 2 {
				line = strings.Trim(line, "\x00")
				line = "CHECK|" + line[1:] + "\x00"
				log.Println(line)
				conn.Write([]byte(line))
			}
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s ip:port\n", os.Args[0])
		return
	}

	hid.UsbWalk(func(device hid.Device) {
		loop(device, os.Args[1])
	})

}
