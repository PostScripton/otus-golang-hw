package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	const input = "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(input))
}
