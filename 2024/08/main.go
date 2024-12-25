package main

import (
	"fmt"
	"strconv"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")

	maxY := len(lines)
	maxX := len(lines[0])
	charToCoords := map[rune][][2]int{}

	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				if _, exists := charToCoords[char]; !exists {
					charToCoords[char] = [][2]int{}
				}

				charToCoords[char] = append(charToCoords[char], [2]int{y, x})
			}
		}
	}

	antiNodeCoords := map[string]bool{}

	for _, coords := range charToCoords {
		if len(coords) < 2 {
			continue
		}

		for a, coord := range coords {
			if part == 2 {
				antiNodeCoords[fmt.Sprint(coord[0], coord[1])] = true
			}

			for b, otherCoord := range coords {
				if a == b {
					continue
				}

				deltaY := coord[0] - otherCoord[0]
				deltaX := coord[1] - otherCoord[1]

				if part == 1 {
					ancY := coord[0] + deltaY
					ancX := coord[1] + deltaX

					if ancY >= 0 && ancY < maxY && ancX >= 0 && ancX < maxX {
						antiNodeCoords[fmt.Sprint(ancY, ancX)] = true
					}
				}

				if part == 2 {
					if deltaY == 0 {
						if deltaX < 0 {
							deltaX = -1
						} else {
							deltaX = 1
						}
					}
					if deltaX == 0 {
						if deltaY < 0 {
							deltaY = -1
						} else {
							deltaY = 1
						}
					}

					ancY := coord[0]
					ancX := coord[1]
					for {
						ancY += deltaY
						ancX += deltaX

						if ancY < 0 || ancY >= maxY || ancX < 0 || ancX >= maxX {
							break
						}

						antiNodeCoords[fmt.Sprint(ancY, ancX)] = true
					}
				}
			}
		}
	}

	answer = len(antiNodeCoords)

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
