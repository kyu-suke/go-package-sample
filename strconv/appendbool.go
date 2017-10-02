package main

import (
	"strconv"
	"fmt"
)

func main () {
	str := []byte("Hello, strconv, ")
	r := strconv.AppendBool(str, true)
	fmt.Println(r)
	fmt.Println(string(r))

	r2 := strconv.AppendBool(str, false)
	fmt.Println(r2)
	fmt.Println(string(r2))
}
