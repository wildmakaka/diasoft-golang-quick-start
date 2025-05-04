package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {

	const text string = "Hello, DIASOFT!"
	fmt.Println(reverse.String(text))
}
