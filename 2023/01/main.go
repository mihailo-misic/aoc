package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int

var wordToNum = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
}

func main() {
	lines := ReadFile("./input.txt")

	// 56001 - too small

	for _, line := range lines {
		pLine := ""
		word := ""
		for _, r := range line {
			sr := string(r)
			if unicode.IsDigit(r) {
				word = ""
				pLine += sr
				continue
			}

			word += sr
			if len(word) < 3 {
				continue
			}

			for w, num := range wordToNum {
				if strings.Contains(word, w) {
					pLine += num
					word = string(word[len(word)-1:])
				}
			}
		}

		first := 'X'
		last := 'X'
		for _, r := range pLine {
			if unicode.IsDigit(r) {
				if first == 'X' {
					first = r
					last = r
				} else {
					last = r
				}
			}
		}
		num, _ := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
		fmt.Println(pLine)

		answer += num
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}
