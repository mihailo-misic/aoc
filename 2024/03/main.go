package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./input.txt")

	rgx := regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)
	do := true
	for _, line := range lines {
		matches := rgx.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			fmt.Println(">>> match", match)
		}

		for _, m := range matches {
			if m[1] == "" {
				do = m[0] == "do()"
				fmt.Println(m[0], do)
				continue
			}

			if !do {
				continue
			}

			if part == 1 || do {
				a, _ := strconv.Atoi(m[1])
				b, _ := strconv.Atoi(m[2])
				fmt.Println(">>> ", a, b)
				answer += a * b
			}
		}

	}

	// 78965138 - too high
	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
