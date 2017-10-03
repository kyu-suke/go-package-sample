package main

import (
	"strconv"
	"fmt"
)

func main() {
	b := []byte("string :")
	b = strconv.AppendQuoteToGraphic(b, "append quote")
	fmt.Println(string(b))

	b2 := []byte("string :")
	b2 = strconv.AppendQuoteToGraphic(b2, `append "quote" ☺ hoge`)
	fmt.Println(string(b2))
}


