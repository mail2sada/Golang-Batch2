package main

import (
	"fmt"
	"strings"
)

// こんにちは
func main() {
	fmt.Println("Strings builder")

	var builder strings.Builder

	builder.Write([]byte("hello"))
	builder.WriteString("12345")
	builder.WriteRune('は')
	str := builder.String()

	fmt.Println(str)
}
