package main

import (
	"fmt"
	"log"
	"net"
)

const (
	server = "0.0.0.0"
	port   = "2345"
)

func main() {
	fmt.Println("UDP Server...")
	address := server + ":" + port
	fmt.Println("Server starting at  ", address)

	//conn, err := net.Listen("udp", address)

	laddress, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalln("Received error while creating udp server", err)
	}
	conn, err := net.ListenUDP("udp", laddress)

	if err != nil {
		log.Fatalln("Received error while creating udp server", err)
	}

	for {
		data := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(data)
		if err != nil {
			fmt.Println("Received error")
			continue
		}
		fmt.Println("read ", n, "Bytes from ", addr)

		response := "Received:" + string(data)

		n, err = conn.WriteTo([]byte(response), addr)
		if err != nil {
			fmt.Println("Received error while writing...")
			continue
		}
		fmt.Println("Wrote ", n, "bytes...")
	}
}
