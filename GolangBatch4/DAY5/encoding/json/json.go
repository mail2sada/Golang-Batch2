package main

import (
	"encoding/json"
	"fmt"
)

var jsonMessage = `[{
	"Name": "Pradeep",
	"Age": 30,
	"Address": "Banglore"
},
{
	"Name": "Anil",
	"Age": 35,
	"Address": "Banglore"
},
{
	"Name": "Ashok",
	"Age": 40,
	"Address": "Banglore"
},
{
	"Name": "Arun",
	"Age": 23,
	"Address": "Banglore"
}]`

type PersonalDetails struct {
	N    string `json:"Name"`
	A    int    `json:"Age"`
	Addr string `json:"Address"`
}

func main() {

	fmt.Println("JSON Encoding")

	pDetail := []PersonalDetails{}

	err := json.Unmarshal([]byte(jsonMessage), &pDetail)
	if err != nil {
		fmt.Println("Issue with json", err)
	} else {
		fmt.Println("Json string is", pDetail)
	}

	data, err := json.MarshalIndent(pDetail, "", "\t")

	if err != nil {
		fmt.Println("Got error")
	} else {
		fmt.Println(string(data))
	}

}
