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

	if PART == 1 {
		part1(commands)
	} else {
		part2(commands)
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}

func part2(commands []Command) {
	finalCommands := []Command{}

	for i, cmd := range commands {
		fmt.Println("Doing", i, "/", len(commands))

		toAdd := []Command{}
		for _, cmd2 := range finalCommands {
			intersection, intersects := cmd.GetIntersectionWith(cmd2)
			if intersects {
				toAdd = append(toAdd, intersection)
			}
		}

		if cmd.Action {
			toAdd = append(toAdd, cmd)
		}

		finalCommands = append(finalCommands, toAdd...)
	}

	for _, cmd := range finalCommands {
		if cmd.Action {
			answer += cmd.GetVolume()
		} else {
			answer -= cmd.GetVolume()
		}
	}
}

func (c *Command) GetIntersectionWith(c2 Command) (intersection Command, intesect bool) {
	x1 := Max([]int{c.X[0], c2.X[0]})
	x2 := Min([]int{c.X[1], c2.X[1]})
	y1 := Max([]int{c.Y[0], c2.Y[0]})
	y2 := Min([]int{c.Y[1], c2.Y[1]})
	z1 := Max([]int{c.Z[0], c2.Z[0]})
	z2 := Min([]int{c.Z[1], c2.Z[1]})

	if x1 > x2 || y1 > y2 || z1 > z2 {
		return Command{}, false
	}

	action := c.Action
	if c.Action && c2.Action {
		action = false
	}
	if !c.Action && !c2.Action {
		action = true
	}

	intersection = Command{
		Action: action,
		X:      [2]int{x1, x2},
		Y:      [2]int{y1, y2},
		Z:      [2]int{z1, z2},
	}

	return intersection, true
}

func (c *Command) GetVolume() int {
	width := c.X[1] - c.X[0] + 1
	height := c.Y[1] - c.Y[0] + 1
	depth := c.Z[1] - c.Z[0] + 1

	return width * height * depth
}

func part1(commands []Command) {
	cubeMap := map[[3]int]bool{}

	for i, cmd := range commands {
		fmt.Println("Doing", i, "/", len(commands))

		if isOutOfBounds(cmd.X) || isOutOfBounds(cmd.Y) || isOutOfBounds(cmd.Z) {
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

func isOutOfBounds(in [2]int) bool {
	MIN := -50
	MAX := 50

	if (MIN <= in[0] && in[0] <= MAX) || (MIN <= in[1] && in[1] <= MAX) {
		return false
	}

	return true
}
