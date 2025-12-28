package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

type Range struct {
	From int
	To   int
}

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")

	ranges := []Range{}

	for _, line := range lines {
		rawRanges := strings.Split(line, ",")

		for _, rawRange := range rawRanges {
			segments := strings.Split(rawRange, "-")

			From, _ := strconv.Atoi(segments[0])
			To, _ := strconv.Atoi(segments[1])

			ranges = append(ranges, Range{From, To})
		}
	}

	for _, r := range ranges {

		for num := r.From; num <= r.To; num++ {
			sNum := strconv.Itoa(num)
			sNumLen := len(sNum)
			mid := sNumLen / 2

			if sNum[0:mid] == sNum[mid:] {
				answer += num
				continue
			}

			if part == 2 {

				for add := 1; add <= mid; add++ {

					if sNumLen%add != 0 {
						continue
					}

					pair := sNum[0:add]
					miss := false

					for j := add; j < sNumLen-add+1; j += add {
						pair2 := sNum[j : j+add]

						if pair != pair2 {
							miss = true
							break
						}
					}

					if miss {
						continue
					}

					answer += num
					break
				}
			}
		}
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
