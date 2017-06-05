package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s [port]\n", os.Args[0])
		return
	}

	udpAddr, err := net.ResolveUDPAddr("udp4", ":"+os.Args[1])
	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	
	for {
		var buf [512]byte
		len, _, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		
		log.Println(string(buf[:len]))
	}
}
