package main

import (
	"fmt"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 1

func main() {
	lines := ReadFile("./sinput.txt")

	for _, line := range lines {
		// Code
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
