package main

import (
	"fmt"
	"regexp"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int

type Command struct {
	Action bool
	X      [2]int
	Y      [2]int
	Z      [2]int
}

const PART = 2

func main() {
	lines := ReadFile("./input.txt")

	commands := make([]Command, len(lines))
	reg := regexp.MustCompile(`^(\w+) x\=(-?\d+)..(-?\d+),y\=(-?\d+)..(-?\d+),z\=(-?\d+)..(-?\d+)$`)

	for idx, line := range lines {
		res := reg.FindStringSubmatch(line)
		XLower, _ := strconv.Atoi(res[2])
		XUpper, _ := strconv.Atoi(res[3])
		YLower, _ := strconv.Atoi(res[4])
		YUpper, _ := strconv.Atoi(res[5])
		ZLower, _ := strconv.Atoi(res[6])
		ZUpper, _ := strconv.Atoi(res[7])

		commands[idx] = Command{
			Action: res[1] == "on",
			X:      [2]int{XLower, XUpper},
			Y:      [2]int{YLower, YUpper},
			Z:      [2]int{ZLower, ZUpper},
		}
	}
	for _, r := range commands {
		fmt.Println(r)
	}

	if PART == 1 {
		part1(commands)
	} else {
		part2(commands)
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}

func part2(commands []Command) {
}

func part1(commands []Command) {
	cubeMap := map[[3]int]bool{}

	for i, cmd := range commands {
		fmt.Println("Doing", i, "/", len(commands))

		if isOut(cmd.X) || isOut(cmd.Y) || isOut(cmd.Z) {
			continue
		}

		for x := cmd.X[0]; x <= cmd.X[1]; x++ {
			for y := cmd.Y[0]; y <= cmd.Y[1]; y++ {
				for z := cmd.Z[0]; z <= cmd.Z[1]; z++ {
					cubeMap[[3]int{x, y, z}] = cmd.Action
				}
			}
		}
	}

	for _, on := range cubeMap {
		if on {
			answer++
		}
	}
}

const MIN = -50
const MAX = 50

func isOut(in [2]int) bool {
	if (MIN <= in[0] && in[0] <= MAX) || (MIN <= in[1] && in[1] <= MAX) {
		return false
	}

	return true
}
