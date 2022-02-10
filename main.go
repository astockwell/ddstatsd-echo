package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	portPtr := flag.Int("port", 8125, "UDP port #")
	flag.Parse()

	s, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%d", *portPtr)) // "udp", "udp4", "udp6"
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()

	log.Println("Started UDP server on port", *portPtr)
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("(%v) -> %v\n", addr, string(buffer[0:n-1]))
	}
}
