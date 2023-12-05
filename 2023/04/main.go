package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
	"github.com/samber/lo"
)

var answer int
var part int = 2

func main() {
	lines := ReadFile("./input.txt")

	columnRgx := regexp.MustCompile(`:\s*`)
	spaceRgx := regexp.MustCompile(`\s+`)
	pipeRgx := regexp.MustCompile(`\s+\|\s+`)
	cardRgx := regexp.MustCompile(`Card\s+`)

	cardToNums := map[int][2][]string{}

	for _, line := range lines {
		cardAndNums := columnRgx.Split(line, -1)
		cardNum, _ := strconv.Atoi(cardRgx.Split(cardAndNums[0], -1)[1])

		winingAndMyNums := pipeRgx.Split(cardAndNums[1], -1)
		winingNums := spaceRgx.Split(winingAndMyNums[0], -1)
		myNums := spaceRgx.Split(winingAndMyNums[1], -1)

		cardToNums[cardNum] = [2][]string{winingNums, myNums}
	}

	cardNums := make([]int, 0, len(cardToNums))
	for k := range cardToNums {
		cardNums = append(cardNums, k)
	}
	sort.Ints(cardNums)

	if part == 1 {
		for _, card := range cardNums {
			nums := cardToNums[card]
			matches := getMatches(card, nums)

			points := 0

			for i := 0; i < matches; i++ {
				if points == 0 {
					points++
					continue
				}

				points *= 2
			}

			answer += points
		}
	}

	if part == 2 {
		cardNumToCopies := map[int]int{}

		for _, card := range cardNums {
			matches := getMatches(card, cardToNums[card])

			for i := matches; i > 0; i-- {
				cardNumToCopies[card+i]++
			}

			answer++
		}

		copyCardNums := make([]int, 0, len(cardNumToCopies))
		for k := range cardNumToCopies {
			copyCardNums = append(copyCardNums, k)
		}
		sort.Ints(copyCardNums)

		for _, cardNum := range copyCardNums {
			copyCount := cardNumToCopies[cardNum]
			matches := getMatches(cardNum, cardToNums[cardNum])

			for i := matches; i > 0; i-- {
				cardNumToCopies[cardNum+i] += copyCount
			}
			answer += copyCount
		}

		fmt.Println("cardNumToCopies", cardNumToCopies)
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

var cardMatchesMemo = map[int]int{}

func getMatches(card int, nums [2][]string) int {
	if val, ok := cardMatchesMemo[card]; ok {
		return val
	}

	matches := 0

	for _, winningNum := range nums[0] {
		if lo.Contains(nums[1], winningNum) {
			matches++
		}
	}

	cardMatchesMemo[card] = matches

	return matches
}
