package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

const PORT = 3456

func main() {
	addr := "0.0.0.0:" + strconv.Itoa(PORT)
	fmt.Println("TCP client...")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln("Could not connect to server...", err)
	}
	var data = make([]byte, 2048)
	//for {

	conn.Write([]byte("Hello from TCP client..."))
	n, err := conn.Read(data)

	if err != nil {
		log.Println("We have received error")
		//continue
	}
	fmt.Println("Read ", n, "Bytes")
	fmt.Println(string(data))
}

//}
