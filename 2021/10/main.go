package main

import (
	"fmt"
	"sort"

	. "github.com/mihailo-misic/aoc/util"
)

var openToClose = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var errToPoint = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func main() {
	input := ReadFile("./input.txt")
	lineScores := []int{}

	for _, line := range input {
		closers := []string{}
		for i, ch := range line {
			c := string(ch)
			if cls, ok := openToClose[c]; ok {
				closers = append(closers, cls)
			} else {
				lCloser := closers[len(closers)-1]
				if c == lCloser {
					closers = closers[:len(closers)-1]
				} else {
					break
				}
			}
			if i == len(line)-1 {
				lineScr := 0
				fmt.Println("allClsr", closers)
				closers := reverse(closers)
				for _, cls := range closers {
					cl := string(cls)
					lineScr *= 5
					lineScr += errToPoint[cl]
				}
				fmt.Println("lineScr", lineScr)
				lineScores = append(lineScores, lineScr)
			}
		}
	}

	fmt.Println("lineScores", lineScores)
	sort.Ints(lineScores)
	midScr := lineScores[len(lineScores)/2]
	fmt.Println("midScr", midScr)
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
