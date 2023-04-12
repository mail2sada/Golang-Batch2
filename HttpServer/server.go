package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Employee struct {
	ID   int    `json:"empid"`
	Name string `json:"name"`
	Dept string `json:"department"`
}

type GetEmpRequest struct {
	ID int `json:"empid"`
}

var empDetails []Employee

func init() {
	empDetails = []Employee{{ID: 1, Name: "Santhosh", Dept: "NGN"},
		{ID: 2, Name: "Ajay", Dept: "Radio"},
		{ID: 3, Name: "Vijay", Dept: "RAN"}}
}

func createEmployee(req *http.Request) error {
	fmt.Println("createEmployee")

	fmt.Println(*req)
	//reader := req.Body
	emp := Employee{}
	//data := make([]byte, 2048)
	body, _ := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	fmt.Println("Got error in reading...", err)
	// 	return errors.New("failed to read ")
	// }
	// fmt.Println("Received ", n, "bytes")
	err := json.Unmarshal(body, &emp)
	if err != nil {
		fmt.Println("Failed to Unmarshal")
		return fmt.Errorf("data not in required format %w", err)
	}
	empDetails = append(empDetails, emp)
	return nil
}

func getEmployee(req *http.Request) ([]byte, error) {
	reader := req.Body
	getRequest := GetEmpRequest{}
	data := make([]byte, 2048)

	data, err := ioutil.ReadAll(req.Body)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Got error in reading...", err)
	// 	return nil, errors.New("failed to read ")
	// }
	fmt.Println("Received ", n, "bytes")
	err = json.Unmarshal(data, &getRequest)
	if err != nil {
		fmt.Println("Failed to Unmarshal")
		return nil, fmt.Errorf("data not in required format %w", err)
	}
	var emp *Employee = nil
	for _, val := range empDetails {
		if val.ID == getRequest.ID {
			emp = &val
			break
		}
	}

	if emp == nil {
		return nil, fmt.Errorf("Details not found %w", err)
	}

	by, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("Marshaling error..")
		return nil, fmt.Errorf("Marshaling error %w", err)
	}

	return by, err
}

func updateEmployee(req *http.Request) error {
	return nil
}

func deleteEmployee(req *http.Request) error {
	return nil
}
func handleEmployee(res http.ResponseWriter, req *http.Request) {
	fmt.Println("handleEmployee", req.Method)
	switch req.Method {
	case "POST":
		createEmployee(req)
	case "GET":
		getEmployee(req)
	case "PUT":
		updateEmployee(req)
	case "DELETE":
		deleteEmployee(req)
	default:
		res.Write([]byte("Bad request"))
	}

}
func handleEmployees(res http.ResponseWriter, req *http.Request) {

	if req.Method != "GET" {
		res.Write([]byte("Invalid request.."))
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	by, err := json.Marshal(empDetails)
	if err != nil {
		fmt.Println("Failed to marshal")
		res.Write([]byte("Failed to convert to json"))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Length", strconv.Itoa(len(by)))
	res.Write(by)
	res.WriteHeader(http.StatusOK)

}

func handleRoot(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to http server"))
	res.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/employee", handleEmployee)
	http.HandleFunc("/employees", handleEmployees)

	err := http.ListenAndServe(":3456", nil)

	fmt.Println(err)
}
