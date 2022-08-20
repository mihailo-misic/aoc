package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/mihailo-misic/aoc/util"
)

type Vector struct {
	X int
	Y int
	Z int
}

var scanners = map[int][]Vector{}

func main() {
	lines := ReadFile("./sinput.txt")

	scanNum := 0
	for _, line := range lines {
		readLines(line, scanNum)
	}

	// for each not matched scanner
	//   for each beacon until match or less than 12 left
	//     for each permutation
	//       check if 12 align

	fmt.Println(scanners)

	fmt.Println("\nAnswer:")
}

func readLines(line string, scanNum int) {
	if line == "" {
		return
	}

	scanNumRgx := regexp.MustCompile(`--- scanner (\d+) ---`)
	matches := scanNumRgx.FindStringSubmatch(line)
	if len(matches) > 1 {
		scanNum, _ = strconv.Atoi(matches[1])
		return
	}

	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])

	scanners[scanNum] = append(scanners[scanNum], Vector{x, y, z})
}
