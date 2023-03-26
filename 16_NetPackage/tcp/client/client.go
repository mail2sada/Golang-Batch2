package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func readFromStdin() []byte {
	var reader io.Reader = os.Stdin
	data := make([]byte, 512)
	fmt.Print("Sent:")
	reader.Read(data)
	//fmt.Println(string(data))

	return data
}

func handleWrite(conn net.Conn) {
	for {
		conn.Write(readFromStdin())
	}
}

func main() {
	fmt.Println("Chat client...")

	conn, err := net.Dial("tcp", "0.0.0.0:5221")

	if err != nil {
		log.Fatalln("Failed to connect exiting...")
	}

	go handleWrite(conn)

	for {
		data := make([]byte, 512)
		n, err := conn.Read(data)

		if n == 0 || err != nil {
			log.Print("Failed to read")
			continue
		}
		fmt.Println("\nReceived:" + string(data))
	}
}
