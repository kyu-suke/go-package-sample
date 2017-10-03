package main

import (
	"strconv"
	"fmt"
)

func main() {
	b := []byte("rune (ascii):")
	b = strconv.AppendQuoteRuneToASCII(b, '☺')
	fmt.Println(string(b))

	b2 := []byte("rune (ascii):")
	b2 = strconv.AppendQuoteRuneToASCII(b2, '爆')
	fmt.Println(string(b2))
}

