package main

import (
	"fmt"
	"strconv"

	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 1

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./sinput.txt")

	for _, line := range lines {
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
