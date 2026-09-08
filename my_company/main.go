package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	fmt.Println(norm.NFC.String("Hello from test"))
}
