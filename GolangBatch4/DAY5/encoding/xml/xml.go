package main

import (
	"encoding/xml"
	"fmt"
)

type PersonalDetails struct {
	Name    string `xml:"Name"`
	Age     int    `xml:"Age"`
	Address string `xml:"Address"`
}

var xmlData = `<PersonalDetails>
<Name>Anand</Name>
<Age>35</Age>
<Address>Bangalore...</Address>
</PersonalDetails>`

func main() {

	fmt.Println("Demo XML Encoding....")

	pDetails := PersonalDetails{Name: "Anand", Age: 35, Address: "Bangalore..."}

	data, err := xml.MarshalIndent(pDetails, "", "  ")

	if err != nil {
		fmt.Println("Got error while ecoding json", err)
	} else {
		fmt.Println("XML Content is ", string(data))
	}

	pDetails1 := new(PersonalDetails)

	err = xml.Unmarshal([]byte(xmlData), pDetails1)
	if err != nil {
		fmt.Println("Received error while decoding xml", err)
	} else {
		fmt.Println(*&pDetails1)
	}

}
