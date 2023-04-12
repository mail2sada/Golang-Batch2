package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

const PORT = 3456

func handleConnection(conn net.Conn) {
	var data = make([]byte, 2048)
	conn.Read(data)
	fmt.Println(string(data))
	n, err := conn.Write([]byte("Welcome to TCP server"))
	fmt.Println(err, n)
}

func main() {
	fmt.Println("Starting TCP Server...")
	addr := "0.0.0.0:" + strconv.Itoa(PORT)
	fmt.Println("Server Address is ", addr)
	listn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Could not listen on port 3456:", err)
	}
	for {
		conn, err := listn.Accept()

		if err != nil {
			log.Println("Failed to connect")
		}
		go handleConnection(conn)
	}

}
