/*

 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID         int    `json:"EmpID"`
	Name       string `json:"EmpName"`
	Department string `json:"Dep"`
	Desig      string `json:"Designation"`
}

func main() {

	var b bytes.Buffer
	var emp = Employee{}

	b.WriteString(`
	{
		"EmpID": 1,
		"EmpName": "Anil Kumar",
		"Dep": "NGN",
		"Designation":"SE"
	}
	`)

	err := json.Unmarshal(b.Bytes(), &emp)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(emp)

	emp1 := Employee{ID: 2, Name: "Nishanth", Department: "NGN", Desig: "SE"}

	by, err := json.Marshal(&emp1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(by))

}
