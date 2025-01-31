package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

type Coord [2]int

type Hiker struct {
	Trail []Coord
	Pos   Coord
}

var world = [][]rune{}

var headToPeaks = map[string]map[string]bool{}

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")

	for _, line := range lines {
		row := []rune{}

		for _, r := range line {
			row = append(row, r)
		}

		world = append(world, row)
	}

	hikers := []Hiker{}

	for y, row := range world {
		for x, r := range row {
			if r == '0' {
				c := Coord{y, x}
				h := Hiker{
					Trail: []Coord{c},
					Pos:   c,
				}

				hikers = append(hikers, h)
			}
		}
	}

	Run(hikers)

	if part == 1 {
		for head, peaks := range headToPeaks {
			fmt.Println(head, len(peaks))
		}

		for _, peaks := range headToPeaks {
			for range peaks {
				answer++
			}
		}
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func Run(hikers []Hiker) {
	newHikers := []Hiker{}

	for _, h := range hikers {
		validMoves := h.GetValidMoves()

		for _, validMove := range validMoves {
			newHiker := h.GoTo(validMove)

			if world[newHiker.Pos[0]][newHiker.Pos[1]] == '9' {
				if part == 1 {
					head := newHiker.Trail[0]
					headKey := fmt.Sprintf("%v:%v", head[0], head[1])
					peakKey := fmt.Sprintf("%v:%v", newHiker.Pos[0], newHiker.Pos[1])

					_, ok := headToPeaks[headKey]
					if !ok {
						headToPeaks[headKey] = map[string]bool{}
					}
					headToPeaks[headKey][peakKey] = true
				}

				if part == 2 {
					answer++
				}
				continue
			}

			newHikers = append(newHikers, newHiker)
		}
	}

	if len(newHikers) > 0 {
		Run(newHikers)
	}
}

func (h *Hiker) GoTo(c Coord) Hiker {
	newTrail := make([]Coord, len(h.Trail))
	copy(newTrail, h.Trail)
	newHiker := Hiker{
		Trail: append(newTrail, c),
		Pos:   c,
	}

	return newHiker
}

func (h *Hiker) GetValidMoves() []Coord {
	validPositions := []Coord{}

	y := h.Pos[0]
	x := h.Pos[1]

	height := world[y][x]

	// Up
	ny := y - 1
	if ny >= 0 {
		np := Coord{ny, x}
		nh := world[np[0]][np[1]]
		if nh == height+1 && !slices.Contains(h.Trail, np) {
			validPositions = append(validPositions, np)
		}
	}

	// Down
	ny = y + 1
	if ny < len(world) {
		np := Coord{ny, x}
		nh := world[np[0]][np[1]]
		if nh == height+1 && !slices.Contains(h.Trail, np) {
			validPositions = append(validPositions, np)
		}
	}

	// Left
	nx := x - 1
	if nx >= 0 {
		np := Coord{y, nx}
		nh := world[np[0]][np[1]]
		if nh == height+1 && !slices.Contains(h.Trail, np) {
			validPositions = append(validPositions, np)
		}
	}

	// Right
	nx = x + 1
	if nx < len(world[0]) {
		np := Coord{y, nx}
		nh := world[np[0]][np[1]]
		if nh == height+1 && !slices.Contains(h.Trail, np) {
			validPositions = append(validPositions, np)
		}
	}

	return validPositions
}
