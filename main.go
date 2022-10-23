package main

import (
	"fmt"
)

func main() {
	input := scanline()

	convertbin2dec(input)
}

func scanline() string {
	fmt.Print("input: ")
	input := ""
	fmt.Scanln(&input)
	return input
}
