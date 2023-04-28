package main

import (
	"fmt"
	"log"
	"net"
)

const (
	serverIP      = "0.0.0.0"
	port          = "1234"
	maxConnection = 10
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	data := make([]byte, 1024)
	for {
		//data := make([]byte, 1024)

		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error while reading...")
			break
		}
		fmt.Println("Read ", n, "bytes...")
		response := "Received:" + string(data)
		fmt.Println(response)
		n, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Received error while writing..")
			break
		}
	}
}

func main() {

	fmt.Println("Demo Tcp Server")

	address := serverIP + ":" + port
	fmt.Println("Listnening @", address)

	listner, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Failed to listen..")
	}
	conCount := 0
	for {
		conn, err := listner.Accept()
		conCount++
		if err != nil {
			fmt.Println("Received error while accepting connection", err)
			continue
		}
		if conCount == maxConnection {
			
		}
		go handleConnection(conn)
	}

}
