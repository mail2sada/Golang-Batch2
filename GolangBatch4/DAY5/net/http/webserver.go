/*

 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Details struct {
	Slno int    `json:"SlNo"`
	Name string `json:"Name"`
}

var Info = []Details{}

func handleRootGet(res http.ResponseWriter, req *http.Request, data Details) {

	for _, val := range Info {
		if val.Slno == data.Slno {
			data, _ := json.MarshalIndent(val, "", "  ")
			res.Write(data)
			res.WriteHeader(http.StatusOK)
			return
		}
	}
	res.Write([]byte("Data not found"))
	res.WriteHeader(http.StatusNotFound)
}
func handleRootPost(res http.ResponseWriter, req *http.Request, data Details) {
	res.Write([]byte("We have hit handleRootPost"))
	Info = append(Info, data)
	res.Write([]byte(fmt.Sprintln("SlNo:", data.Slno, "is stored")))
	res.WriteHeader(http.StatusOK)
}
func handleRootPut(res http.ResponseWriter, req *http.Request, data Details) {
	res.Write([]byte("We have hit handleRootPut"))
}
func handleRootDelete(res http.ResponseWriter, req *http.Request, data Details) {
	res.Write([]byte("We have hit handleRootDelete"))
}
func HandleRoot(res http.ResponseWriter, req *http.Request) {

	contentType := req.Header["Content-Type"]
	fmt.Println("Content Type is ", contentType)
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Read error", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	d := Details{}
	err = json.Unmarshal(data, &d)
	if err != nil {
		fmt.Println("Got error while unmarshaling json", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if contentType[0] != "application/json" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	switch req.Method {
	case "GET":
		handleRootGet(res, req, d)
	case "POST":
		handleRootPost(res, req, d)
	case "PUT":
		handleRootPut(res, req, d)
	case "DELETE":
		handleRootDelete(res, req, d)
	default:
		res.WriteHeader(http.StatusBadRequest)
	}

}

func main() {

	fmt.Println("WebServer::")

	serverAddress := "0.0.0.0:5151"

	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/GetAll", HandleGetAll)

	http.ListenAndServe(serverAddress, nil)
}

func HandleGetAll(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		res.Write([]byte("Only GET is allowed"))
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := json.MarshalIndent(Info, "", " ")
	if err != nil {
		fmt.Println("Got error while marshling")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.Write(data)
	res.WriteHeader(http.StatusOK)
}
