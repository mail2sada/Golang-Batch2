package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("Demo: Bytes builder")

	var builder bytes.Buffer

	builder.WriteString("Hello and welcome")
	builder.WriteRune('は')
	builder.WriteString("こんにちは")
	data := builder.Bytes()
	fmt.Println(string(data))
}
