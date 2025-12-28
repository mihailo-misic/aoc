package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

type Move struct {
	Direction string
	Amount    int
}

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")

	moves := []Move{}
	for _, line := range lines {
		split := strings.SplitAfterN(line, "", 2)
		amount, _ := strconv.Atoi(split[1])
		move := Move{
			Direction: split[0],
			Amount:    amount,
		}
		moves = append(moves, move)
	}

	position := 50

	for _, move := range moves {
		if part == 2 {
			answer += int(math.Abs(float64(move.Amount / 100)))
		}
		was0 := position == 0

		if move.Direction == "L" {
			position -= move.Amount % 100
		}
		if move.Direction == "R" {
			position += move.Amount % 100
		}

		if position < 0 {
			if part == 2 && !was0 {
				answer++
			}

			position += 100
		}
		if position > 99 {
			position -= 100

			if part == 2 && position != 0 {
				answer++
			}
		}

		if position == 0 {
			answer++
		}
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
