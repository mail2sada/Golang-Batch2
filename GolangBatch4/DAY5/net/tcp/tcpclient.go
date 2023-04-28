package main

import (
	"fmt"
	"log"
	"net"
)

const (
	serverAddress = "0.0.0.0"
	port          = "1234"
)

func main() {
	fmt.Println("Tcp Client")
	address := serverAddress + ":" + port
	fmt.Println("Connecting to Address", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln("Received error while connecting to server, ", address, err)
	}
	defer conn.Close()
	messages := []string{"Hello from client", "This is first message", "wating for some messsages", "exit"}
	for {

		for _, msg := range messages {
			n, err := conn.Write([]byte(msg))
			if err != nil {
				fmt.Println("Received error while writing", err)
				break
			}
			fmt.Println("Wrote ", n, "Bytes")
			data := make([]byte, 1024)
			n, err = conn.Read(data)
			if err != nil {
				fmt.Println("Received error while reading...")
				break
			}
			fmt.Println("Read ", n, "Bytes")
			fmt.Println(string(data))
		}
	}

}
