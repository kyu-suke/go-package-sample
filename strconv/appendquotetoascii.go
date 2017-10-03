package main

import (
	"strconv"
	"fmt"
)

func main() {
	b := []byte("quote (ascii):")
	b = strconv.AppendQuoteToASCII(b, `"Fran & Freddi爆e's Diner"`)
	fmt.Println(string(b))

}
