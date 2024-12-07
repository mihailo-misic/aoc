package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
	"golang.org/x/exp/slices"
)

var answer int
var part int = 2

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./input.txt")

	rules := [][]int{}
	sequences := [][]int{}

	rulesTime := true
	for _, line := range lines {
		if line == "" {
			rulesTime = false
			continue
		}

		if rulesTime {
			rStrings := strings.Split(line, "|")
			r1, _ := strconv.Atoi(rStrings[0])
			r2, _ := strconv.Atoi(rStrings[1])
			rules = append(rules, []int{r1, r2})
		}

		if !rulesTime {
			sStrings := strings.Split(line, ",")
			seq := []int{}
			for _, sString := range sStrings {
				s, _ := strconv.Atoi(sString)
				seq = append(seq, s)
			}
			sequences = append(sequences, seq)
		}
	}

	for _, seq := range sequences {
		invalidPair := getInvalidPair(seq, rules)

		if invalidPair == nil {
			if part == 1 {
				answer += seq[len(seq)/2]
			}

			if part == 2 {
				continue
			}
		}

		if part == 2 {
			if invalidPair != nil {
				swapUntilValid(seq, invalidPair[0], invalidPair[1], rules)
				answer += seq[len(seq)/2]
			}
		}
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func swapUntilValid(seq []int, l, r int, rules [][]int) {
	seq[l], seq[r] = seq[r], seq[l]

	invalidPair := getInvalidPair(seq, rules)
	if invalidPair == nil {
		return
	}

	swapUntilValid(seq, invalidPair[0], invalidPair[1], rules)
}

func getInvalidPair(seq []int, rules [][]int) []int {
	relevantRules := [][]int{}

	for _, r := range rules {
		if slices.Contains(seq, r[0]) && slices.Contains(seq, r[1]) {
			relevantRules = append(relevantRules, r)
		}
	}

	for ridx, s := range seq {
		for _, r := range relevantRules {
			if r[0] == s {
				for lidx, n := range seq[:ridx] {
					if n == r[1] {
						return []int{lidx, ridx}
					}
				}
			}
		}
	}

	return nil
}
