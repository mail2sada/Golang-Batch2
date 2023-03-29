package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	var builder strings.Builder = strings.Builder{}

	data := []byte{'X', 'Y', 'Z'}
	builder.WriteString("Hello")
	fmt.Println(builder.String())

	builder.WriteByte('A')

	fmt.Println(builder.String())

	builder.WriteRune('¶')
	fmt.Println(builder.String())
	builder.Write(data)
	fmt.Println(builder.String())

	var buffer bytes.Buffer

	buffer.

}
