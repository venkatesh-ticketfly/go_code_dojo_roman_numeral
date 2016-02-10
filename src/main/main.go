package main

import (
	"fmt"
	"roman"
)

func main() {
	val, err := roman.NewNumerals("I")
	fmt.Println(val, err)
}
