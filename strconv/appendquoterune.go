package main

import (
	"strconv"
	"fmt"
)

func main() {
	b := []byte("rune:")
	b = strconv.AppendQuoteRune(b, '☺')
	fmt.Println(string(b))

	b2 := []byte("rune:")
	b2 = strconv.AppendQuoteRune(b2, '❗')
	fmt.Println(string(b2))

	b3 := []byte("rune:")
	b3 = strconv.AppendQuoteRune(b3, 'r')
	fmt.Println(string(b3))
}
