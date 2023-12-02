package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

func main() {
	lines := ReadFile("./input.txt")

	gameRgx := regexp.MustCompile(`Game (\d+): `)
	handRgx := regexp.MustCompile(`(\d+) (\w+)`)
	for _, line := range lines {
		match := gameRgx.FindStringSubmatch(line)

		gameId, err := strconv.Atoi(match[1])
		if err != nil {
			panic("gameId is not a number")
		}

		gameStr := gameRgx.ReplaceAllString(line, "")
		possible := true

		colorToMin := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		turns := strings.Split(gameStr, ";")
		for _, turn := range turns {
			if !possible && part == 1 {
				break
			}

			hands := strings.Split(turn, ", ")
			for _, hand := range hands {
				colorToAmount := map[string]int{
					"red":   12,
					"green": 13,
					"blue":  14,
				}

				handMatch := handRgx.FindStringSubmatch(hand)
				amount, color := handMatch[1], handMatch[2]

				cAmount, ok := colorToAmount[color]
				if ok != true && part == 1 {
					possible = false
					panic(fmt.Sprintf("Unknown color: %v", color))
				}

				if cAmount > 0 {
					amountNum, err := strconv.Atoi(amount)
					if err != nil {
						panic("amount is not a number")
					}

					colorToAmount[color] -= amountNum

					if colorToMin[color] < amountNum {
						colorToMin[color] = amountNum
					}
				}

				if colorToAmount[color] < 0 && part == 1 {
					possible = false
					break
				}
			}
		}

		if part == 1 {
			if possible {
				answer += gameId
			}
		}

		if part == 2 {
			prod := 1
			for _, min := range colorToMin {
				prod *= min
			}
			answer += prod
		}
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}
