package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

type StepInfo struct {
	GhostId int
	StepNum int
}

var instructions = []string{}
var network = map[string]map[string]string{}

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./input.txt")

	nodeRgx := regexp.MustCompile(`(\w+) = \((\w+), (\w+)`)

	for lIdx, line := range lines {
		if lIdx == 0 {
			instructions = strings.Split(line, "")
			continue
		}

		if lIdx == 1 {
			continue
		}

		match := nodeRgx.FindStringSubmatch(line)

		network[match[1]] = map[string]string{}
		network[match[1]]["L"] = match[2]
		network[match[1]]["R"] = match[3]
	}

	if part == 1 {
		at := "AAA"

		for steps := 0; ; steps++ {
			if at == "ZZZ" {
				answer = steps
				break
			}

			inst := instructions[steps%len(instructions)]
			at = network[at][inst]
		}
	}

	if part == 2 {
		ghosts := []string{}

		for pos := range network {
			if strings.HasSuffix(pos, "A") {
				ghosts = append(ghosts, pos)
			}
		}

		ghostsStepCount := []int{}

		for _, ghost := range ghosts {
			at := ghost

			for steps := 1; ; steps++ {
				at = network[at][instructions[(steps-1)%len(instructions)]]

				if strings.HasSuffix(at, "Z") {
					ghostsStepCount = append(ghostsStepCount, steps)
					break
				}
			}
		}

		answer = lcm(ghostsStepCount[0], ghostsStepCount[1], ghostsStepCount[2:]...)
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func gcd(a, b int) int {
	for b != 0 {
		bBefore := b
		b = a % b
		a = bBefore
	}

	return a
}

func lcm(a, b int, numbers ...int) int {
	result := a * b / gcd(a, b)

	for _, num := range numbers {
		result = lcm(result, num)
	}

	return result
}
