package main

import (
	"fmt"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int

func main() {
	lines := ReadFile("./sinput.txt")

	for _, line := range lines {
		// Code
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}
