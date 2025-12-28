package main

import (
	"fmt"
	"strconv"

	"github.com/mihailo-misic/aoc/util"
)

const part int = 1
const inputPath string = "./sinput.txt"

var answer int

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile(inputPath)

	for _, line := range lines {
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
