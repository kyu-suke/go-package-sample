package main

import (
	"strconv"
	"fmt"
)

func main() {
	b := []byte("rune (ascii):")
	b = strconv.AppendQuoteRuneToGraphic(b, '☺')
	fmt.Println(string(b))

	b2 := []byte("rune (ascii):")
	b2 = strconv.AppendQuoteRuneToGraphic(b2, '爆')
	fmt.Println(string(b2))
}

