package main

import (
	"fmt"
	"net"
	"os"
	"github.com/chzyer/readline"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s ip:port\n", os.Args[0])
		return
	}

	udpAddr, err := net.ResolveUDPAddr("udp4", os.Args[1])
	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}
	
	conn, err := net.DialUDP("udp", nil, udpAddr)
	
	if err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}
	
	defer conn.Close()
	
	var completer = readline.NewPrefixCompleter(
		readline.PcItem("OPEN"),
	)
	rl, err := readline.NewEx(&readline.Config{
			Prompt:       "> ",
			AutoComplete: completer,
	})
	
	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		line += "\x00"
		conn.Write([]byte(line))
	}
}
