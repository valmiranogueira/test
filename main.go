package main

import (
    "fmt"

    "golang.org/x/text/unicode/norm"
)

func main() {
    input := "Percona"
    fmt.Println(norm.NFC.String(input))
}
