package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kr/pretty"
	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
	"github.com/samber/lo"
)

var answer int
var part int = 2

var grid = [][]string{}

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./s2input.txt")
	startPos := [2]int{}

	for lIdx, line := range lines {
		row := strings.Split(line, "")
		grid = append(grid, row)

		if lo.Contains(row, "S") {
			for rIdx, node := range row {
				if node == "S" {
					startPos = [2]int{lIdx, rIdx}
				}
			}
		}
	}

	grid[startPos[0]][startPos[1]] = recognizeS(startPos)

	nodeLoop := getNodeLoopFromS(startPos)

	if part == 1 {
		answer = len(nodeLoop) / 2
	}
	if part == 2 {
		answer = getTilesInLoopCount(nodeLoop)
		pretty.Println("grid", grid)
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

var openerToCompatible = map[string][]string{
	"L": {"L", "7", "J", "-"},
	"F": {"F", "7", "J", "-"},
	"|": {"L", "F", "-"},
}

func getTilesInLoopCount(nodeLoop [][2]int) int {
	removeTrash(nodeLoop)

	tilesInCount := 0

	for rIdx, row := range grid {
		opener := ""

		for cIdx, c := range row {
			if opener == "" && c != "." && c != "-" {
				opener = c
				continue
			}

			if opener != "" {
				if c == "." {
					tilesInCount++
					grid[rIdx][cIdx] = "X"
					continue
				}

				if !lo.Contains(openerToCompatible[c], c) {
					opener = ""
				}
			}
		}
	}

	return tilesInCount
}

func removeTrash(nodeLoop [][2]int) {
	for rIdx, row := range grid {
		for cIdx, gChar := range row {
			if gChar != "." {
				coords := [2]int{rIdx, cIdx}
				if !lo.Contains(nodeLoop, coords) {
					grid[rIdx][cIdx] = "."
				}
			}
		}
	}
}

func recognizeS(s [2]int) string {
	if hasOpeningTo(s, "up") {
		if hasOpeningTo(s, "down") {
			return "|"
		}
		if hasOpeningTo(s, "right") {
			return "L"
		}
		if hasOpeningTo(s, "left") {
			return "J"
		}
	}

	if hasOpeningTo(s, "right") {
		if hasOpeningTo(s, "left") {
			return "-"
		}
	}

	if hasOpeningTo(s, "down") {
		if hasOpeningTo(s, "left") {
			return "7"
		}
		if hasOpeningTo(s, "right") {
			return "F"
		}
	}

	return "X"
}

func getNodeLoopFromS(startPos [2]int) [][2]int {
	loop := [][2]int{startPos}

	for {
		pos := getNext(loop)

		if pos[0] == -1 {
			return loop
		}

		loop = append(loop, pos)
	}

}

func canGo(p [2]int, dir string) bool {
	pc := grid[p[0]][p[1]]

	if dir == "up" {
		return p[0]-1 >= 0 && lo.Contains([]string{"|", "L", "J"}, pc)
	}
	if dir == "right" {
		return p[1]+1 < len(grid[0]) && lo.Contains([]string{"-", "L", "F"}, pc)
	}
	if dir == "down" {
		return p[0]+1 < len(grid) && lo.Contains([]string{"|", "7", "F"}, pc)
	}
	if dir == "left" {
		return p[1]-1 >= 0 && lo.Contains([]string{"-", "J", "7"}, pc)
	}

	return false
}

func hasOpeningTo(p [2]int, dir string) bool {
	if dir == "up" && p[0]-1 >= 0 {
		nc := grid[p[0]-1][p[1]]

		return lo.Contains([]string{"|", "7", "F"}, nc)
	}
	if dir == "right" && p[1]+1 < len(grid[0]) {
		nc := grid[p[0]][p[1]+1]

		return lo.Contains([]string{"-", "J", "7"}, nc)
	}
	if dir == "down" && p[0]+1 < len(grid) {
		nc := grid[p[0]+1][p[1]]

		return lo.Contains([]string{"|", "L", "J"}, nc)
	}
	if dir == "left" && p[1]-1 >= 0 {
		nc := grid[p[0]][p[1]-1]

		return lo.Contains([]string{"-", "L", "F"}, nc)
	}

	return false
}

func getNext(loop [][2]int) [2]int {
	p := loop[len(loop)-1]

	up := [2]int{p[0] - 1, p[1]}
	if canGo(p, "up") && hasOpeningTo(p, "up") && !lo.Contains(loop, up) {
		return up
	}

	rp := [2]int{p[0], p[1] + 1}
	if canGo(p, "right") && hasOpeningTo(p, "right") && !lo.Contains(loop, rp) {
		return rp
	}

	dp := [2]int{p[0] + 1, p[1]}
	if canGo(p, "down") && hasOpeningTo(p, "down") && !lo.Contains(loop, dp) {
		return dp
	}

	lp := [2]int{p[0], p[1] - 1}
	if canGo(p, "left") && hasOpeningTo(p, "left") && !lo.Contains(loop, lp) {
		return lp
	}

	return [2]int{-1, -1}
}
