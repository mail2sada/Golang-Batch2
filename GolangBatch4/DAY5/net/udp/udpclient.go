package main

import (
	"fmt"
	"net"
	"os"
)

const (
	server = "0.0.0.0"
	port   = "2345"
)

func main() {

	fmt.Println("Udp client...")
	address := server + ":" + port

	conn, err := net.Dial("udp", address)

	if err != nil {
		fmt.Println("Received error", err)
		os.Exit(1)
	}

	var messages = []string{"Hello from client ", "1234", "abc", `fmt.Println("Udp client...")
	address := server + ":" + port

	conn, err := net.Dial("udp", address)

	if err != nil {
		fmt.Println("Received error", err)
		os.Exit(1)
	}

	var messages = []string{"Hello from client ", "1234", "abc" }
	conn.Write([]byte("Hello from client"))`}

	for _, msg := range messages {
		n, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Received err", err)
			continue
		}
		fmt.Println("Worte ", n, "Bytes")
		data := make([]byte, 1024)
		n, err = conn.Read(data)
		if err != nil {
			fmt.Println("Received err while reading", err)
			continue
		}
		fmt.Println("Read ", n, "Bytes", string(data))
	}
}
