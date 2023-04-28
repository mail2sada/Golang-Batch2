/*
root GET, PUT, POST, DELETE, TRACE
root/a GET, PUT
root/b GET PUT
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	server = "0.0.0.0"
	port   = "5151"
)

func main() {
	fmt.Println("Demo Webserver")

	serverAddress := server + ":" + port
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("We have got a request @ root")
		fmt.Println("Method invoked is", req.Method)
		res.Write([]byte("Webserver @ root:" + serverAddress))
	})
	http.HandleFunc("/user", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("We have got a request @ /user")
		fmt.Println("Method invoked is", req.Method)
		res.Write([]byte("Webserver @ /user:" + serverAddress))
	})
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatalln("Failed to create server", err)
	}

}
