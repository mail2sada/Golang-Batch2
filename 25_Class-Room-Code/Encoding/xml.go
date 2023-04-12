/*
`<Employee><EmpID>1</EmpID><EmpName>Anil Kumar</EmpName><Dep>NGN</Dep><Designation>SE</Designation></Employee>`
*/

package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type Employee struct {
	ID         int    `xml:"EmpID"`
	Name       string `xml:"EmpName"`
	Department string `xml:"Dep"`
	Desig      string `xml:"Designation"`
}

func main() {

	var b bytes.Buffer
	var emp = Employee{ID: 1, Name: "Anil Kumar", Department: "NGN", Desig: "SE"}

	by, err := xml.Marshal(&emp)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(by))
	b.WriteString(`<Employee><EmpID>1</EmpID><EmpName>Anil Kumar</EmpName><Dep>NGN</Dep><Designation>SE</Designation></Employee>`)

	var emp1 = Employee{}
	xml.Unmarshal(b.Bytes(), &emp1)

	fmt.Println(emp1)
}
