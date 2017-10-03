package main

import (
	"strconv"
	"fmt"
)

func main() {
	b := []byte("quote:")
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

	b2 := []byte("quote:")
	b2 = strconv.AppendQuote(b2, "he say 'hoge hg☺o爆e', ")
	fmt.Println(string(b2))

}
