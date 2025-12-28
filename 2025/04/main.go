package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")

	grid := [][]string{}

	for _, line := range lines {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	answer = clearGrid(grid)

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func clearGrid(grid [][]string) (cleared int) {
	for y := 0; y < len(grid); y++ {

		for x, char := range grid[y] {

			if char == "@" {
				neighbours := 0

				// Top left
				if safeGet(grid, y-1, x-1) == "@" {
					neighbours++
				}
				// Top
				if safeGet(grid, y-1, x) == "@" {
					neighbours++
				}
				// Top right
				if safeGet(grid, y-1, x+1) == "@" {
					neighbours++
				}

				// Left
				if safeGet(grid, y, x-1) == "@" {
					neighbours++
					if neighbours > 3 {
						continue
					}
				}
				// Right
				if safeGet(grid, y, x+1) == "@" {
					neighbours++
					if neighbours > 3 {
						continue
					}
				}

				// Bottom Left
				if safeGet(grid, y+1, x-1) == "@" {
					neighbours++
					if neighbours > 3 {
						continue
					}
				}
				// Bottom
				if safeGet(grid, y+1, x) == "@" {
					neighbours++
					if neighbours > 3 {
						continue
					}
				}
				// Bottom right
				if safeGet(grid, y+1, x+1) == "@" {
					neighbours++
					if neighbours > 3 {
						continue
					}
				}

				cleared++
				grid[y][x] = "."
			}
		}
	}

	if part == 2 && cleared > 0 {
		cleared += clearGrid(grid)
	}

	return cleared
}

func safeGet[V comparable](grid [][]V, y int, x int) (char V) {
	isOutOfBounds := y < 0 || x < 0 || y >= len(grid) || x >= len(grid[0])

	if isOutOfBounds {
		return
	}

	return grid[y][x]
}
