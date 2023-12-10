package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
	"github.com/samber/lo"
)

var answer int
var part int = 2

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./input.txt")

	sequences := [][]int{}

	for _, line := range lines {
		sLine := strings.Split(line, " ")
		sequence := lo.Map(sLine, func(num string, idx int) int {
			numI, _ := strconv.Atoi(num)
			return numI
		})
		sequences = append(sequences, sequence)
	}

	for _, seq := range sequences {
		answer += findNextNum(seq)
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func findNextNum(seq []int) int {
	seqs := [][]int{
		seq,
	}

	for sIdx := 1; ; sIdx++ {
		prevSeq := seqs[sIdx-1]
		newSeq := []int{}

		allDelta0 := true

		for psIdx := 0; psIdx < len(prevSeq)-1; psIdx++ {
			delta := prevSeq[psIdx+1] - prevSeq[psIdx]
			newSeq = append(newSeq, delta)
			if delta != 0 {
				allDelta0 = false
			}
		}

		seqs = append(seqs, newSeq)

		if allDelta0 {
			break
		}
	}

	if part == 1 {
		for sIdx := len(seqs) - 1; sIdx > -1; sIdx-- {
			if sIdx == len(seqs)-1 {
				seqs[sIdx] = append(seqs[sIdx], 0)
				continue
			}

			a := seqs[sIdx][len(seqs[sIdx])-1]
			b := seqs[sIdx+1][len(seqs[sIdx+1])-1]
			seqs[sIdx] = append(seqs[sIdx], a+b)
		}

		return seqs[0][len(seqs[0])-1]
	}

	for sIdx := len(seqs) - 1; sIdx > -1; sIdx-- {
		if sIdx == len(seqs)-1 {
			seqs[sIdx] = append(seqs[sIdx], 0)
			continue
		}

		b := seqs[sIdx][0]
		a := seqs[sIdx+1][0]
		seqs[sIdx] = append([]int{b - a}, seqs[sIdx]...)
	}

	return seqs[0][0]
}
