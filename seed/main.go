package main

import (
	"fmt"
	"strconv"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 1

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")

	for _, line := range lines {
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
