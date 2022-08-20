package main

import (
	"fmt"

	. "github.com/mihailo-misic/aoc/util"
)

func main() {
	input := ReadFile("./input.txt")

	matrix := [][]int{}
	for _, line := range input {
		row := []int{}
		for _, sq := range line {
			row = append(row, int(sq-'0'))
		}
		matrix = append(matrix, row)
		fmt.Println("row", row)
	}

	total := len(matrix) * len(matrix[0])

	for s := 0; s < 1000; s++ {
		flashes := 0
		fmt.Println("---------------------------------")
		// Grow
		for y, row := range matrix {
			for x := range row {
				matrix[y][x]++
			}
		}

		// Flash (maybe)
		for y, row := range matrix {
			for x, sq := range row {
				if sq == 10 {
					flash(&matrix, y, x, &flashes)
				}
			}
		}
		for _, row := range matrix {
			fmt.Println("row", row)
		}

		fmt.Println("flashes", flashes)
		if flashes == total {
			fmt.Println("STEP", s+1)
			break
		}
	}
}

func flash(mx *[][]int, y, x int, flashes *int) {
	*flashes++
	(*mx)[y][x] = 0

	hasUp := y-1 >= 0
	hasDown := y+1 < len(*mx)
	hasLeft := x-1 >= 0
	hasRight := x+1 < len((*mx)[0])

	if hasUp {
		if hasLeft {
			// ul
			if (*mx)[y-1][x-1] != 0 {
				(*mx)[y-1][x-1]++
				if (*mx)[y-1][x-1] > 9 {
					flash(mx, y-1, x-1, flashes)
				}
			}
		}

		// inc u
		if (*mx)[y-1][x] != 0 {
			(*mx)[y-1][x]++
			if (*mx)[y-1][x] > 9 {
				flash(mx, y-1, x, flashes)
			}
		}

		if hasRight {
			// inc ur
			if (*mx)[y-1][x+1] != 0 {
				(*mx)[y-1][x+1]++
				if (*mx)[y-1][x+1] > 9 {
					flash(mx, y-1, x+1, flashes)
				}
			}
		}
	}

	if hasRight {
		// inc r
		if (*mx)[y][x+1] != 0 {
			(*mx)[y][x+1]++
			if (*mx)[y][x+1] > 9 {
				flash(mx, y, x+1, flashes)
			}
		}
	}

	if hasDown {
		if hasRight {
			// inc dr
			if (*mx)[y+1][x+1] != 0 {
				(*mx)[y+1][x+1]++
				if (*mx)[y+1][x+1] > 9 {
					flash(mx, y+1, x+1, flashes)
				}
			}
		}

		// inc d
		if (*mx)[y+1][x] != 0 {
			(*mx)[y+1][x]++
			if (*mx)[y+1][x] > 9 {
				flash(mx, y+1, x, flashes)
			}
		}

		if hasLeft {
			// inc dl
			if (*mx)[y+1][x-1] != 0 {
				(*mx)[y+1][x-1]++
				if (*mx)[y+1][x-1] > 9 {
					flash(mx, y+1, x-1, flashes)
				}
			}
		}
	}

	if hasLeft {
		// inc l
		if (*mx)[y][x-1] != 0 {
			(*mx)[y][x-1]++
			if (*mx)[y][x-1] > 9 {
				flash(mx, y, x-1, flashes)
			}
		}
	}
}
