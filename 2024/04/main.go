package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./input.txt")

	board := make([][]string, len(lines))
	for lidx, line := range lines {
		board[lidx] = strings.Split(line, "")
	}

	for y, row := range board {
		for x, l := range row {
			if part == 1 {
				words := []string{
					getRight(board, y, x),
					getDown(board, y, x),
					getDiagRight(board, y, x),
					getDiagLeft(board, y, x),
				}

				for _, word := range words {
					if word == "XMAS" || word == "SAMX" {
						answer++
					}
				}
			}

			if part == 2 {
				// Skip if on edge
				if y == 0 || y == len(board)-1 || x == 0 || x == len(board[0])-1 {
					continue
				}

				// Not A? I don't cAre
				if l != "A" {
					continue
				}

				dr := board[y-1][x-1] + board[y+1][x+1]
				if dr != "SM" && dr != "MS" {
					continue
				}

				dl := board[y-1][x+1] + board[y+1][x-1]
				if dl != "SM" && dl != "MS" {
					continue
				}

				answer++
			}
		}
	}

	/*

			S  S  S
		     A A A
		      MMM
			SAMXMAS
			  MMM
			 A A A
			S  S  S

	*/

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func getRight(b [][]string, y, x int) (str string) {
	if x+3 >= len(b[0]) {
		return ""
	}

	for i := 0; i < 4; i++ {
		str += b[y][x+i]
	}

	return str
}

func getDown(b [][]string, y, x int) (str string) {
	if y+3 >= len(b) {
		return ""
	}

	for i := 0; i < 4; i++ {
		str += b[y+i][x]
	}

	return str
}

func getDiagRight(b [][]string, y, x int) (str string) {
	if y+3 >= len(b) || x+3 >= len(b) {
		return ""
	}

	for i := 0; i < 4; i++ {
		str += b[y+i][x+i]
	}

	return str
}

func getDiagLeft(b [][]string, y, x int) (str string) {
	if y+3 >= len(b) || x-3 < 0 {
		return ""
	}

	for i := 0; i < 4; i++ {
		str += b[y+i][x-i]
	}

	return str
}
