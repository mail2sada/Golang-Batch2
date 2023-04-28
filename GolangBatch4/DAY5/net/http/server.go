package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Demo: Http server")

	route := http.NewServeMux()

	route.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("You have hit 0.0.0.0:5151"))
		res.WriteHeader(http.StatusOK)
	})

	route1 := http.NewServeMux()

	route1.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("You have hit 0.0.0.0:3232"))
		res.WriteHeader(http.StatusOK)
	})

	go http.ListenAndServe("0.0.0.0:5151", route)

	http.ListenAndServe("0.0.0.0:3232", route1)
}
