package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/krasi-georgiev/rpiGpio"
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

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	len, _, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
	recieved := string(buf[:len])
	log.Println(recieved)

	match := "OPEN|"
	if strings.HasPrefix(recieved, match) {
		go func() {
			t := rpiGpio.NewControl()
			t.SetType("timer")
			t.SetDelay("5s")
			t.SetPin("18")
			if err := t.StartTimer(nil); err != nil {
				fmt.Println(os.Stderr, err.Error())
			}
		}()
	} else {
		log.Println("Recieved string:", recieved, ", doesn't include:", match)
	}

}
