package main

import (
	"fmt"
	"strconv"
	"unicode"

	. "github.com/mihailo-misic/aoc/util"
	"github.com/samber/lo"
)

var answer int
var part int = 2

var NOT_INTERESTING = []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '.'}

type Point struct {
	Val    rune
	Coords [2]int
}

func main() {
	lines := ReadFile("./input.txt")

	symbolCoords := [][2]int{}

	board := [][]rune{}

	for rIdx, line := range lines {
		row := []rune{}

		for cIdx, char := range line {
			row = append(row, char)

			if !lo.Contains(NOT_INTERESTING, char) {
				symbolCoords = append(symbolCoords, [2]int{rIdx, cIdx})
			}
		}

		board = append(board, row)
	}

	for _, sc := range symbolCoords {
		numbers := getSurroundingNumbers(board, sc)

		if part == 1 {
			for _, number := range numbers {
				answer += number
			}
		}

		if part == 2 {
			if board[sc[0]][sc[1]] == '*' && len(numbers) > 1 {
				product := 1

				for _, number := range numbers {
					product *= number
				}

				answer += product
			}
		}
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func getSurroundingNumbers(board [][]rune, coords [2]int) []int {
	neighbours := getNeighbours(board, coords)
	neighbourDigiPoints := lo.Filter(neighbours, func(np Point, idx int) bool {
		if !unicode.IsDigit(np.Val) {
			return false
		}

		return true
	})

	numbers := []int{}

	for _, ndp := range neighbourDigiPoints {
		num := expandNumber(board, ndp)
		if num > 0 {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func getNeighbours(board [][]rune, coords [2]int) []Point {
	y, x := coords[0], coords[1]
	points := []Point{}

	for rIdx := y - 1; rIdx <= y+1; rIdx++ {
		for cIdx := x - 1; cIdx <= x+1; cIdx++ {
			// OOB
			if rIdx < 0 || cIdx < 0 || rIdx >= len(board) || cIdx >= len(board[0]) {
				continue
			}

			// Same coord
			if rIdx == y && cIdx == x {
				continue
			}

			points = append(points, Point{
				Val:    board[rIdx][cIdx],
				Coords: [2]int{rIdx, cIdx},
			})
		}
	}

	return points
}

var coordsDone = map[string]bool{}

func expandNumber(board [][]rune, p Point) int {
	startIdx := p.Coords[1]

	for idx := startIdx - 1; idx >= 0; idx-- {
		charAtIdx := board[p.Coords[0]][idx]

		if !unicode.IsDigit(charAtIdx) {
			break
		}

		startIdx = idx
	}

	coordKey := fmt.Sprintf("%v:%v", startIdx, p.Coords[0])
	if coordsDone[coordKey] {
		return 0
	}
	coordsDone[coordKey] = true

	numStr := ""
	for idx := startIdx; idx < len(board[0]); idx++ {
		charAtIdx := board[p.Coords[0]][idx]

		if !unicode.IsDigit(charAtIdx) {
			break
		}

		numStr += string(charAtIdx)
	}

	number, err := strconv.Atoi(numStr)
	if err != nil {
		panic(fmt.Sprintf("Numstr is NaN: %v", err))
	}

	return number
}

func printBoard(board [][]rune) {
	for _, row := range board {
		fmt.Println(lo.Map(row, func(r rune, idx int) string { return string(r) }))
	}
	fmt.Println()
}
