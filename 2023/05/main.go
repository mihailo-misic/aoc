package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
	"github.com/samber/lo"
)

var answer int
var part int = 2

type Range struct {
	SrcStart int
	SrcEnd   int
	DstStart int
	DstEnd   int
}

var mappings = [][]Range{}

func collide(s [2]int, r Range) ([2]int, [2][2]int) {
	toMove := [2]int{
		max(s[0], r.SrcStart),
		min(s[1], r.SrcEnd),
	}
	moved := [2]int{
		toMove[0] - r.SrcStart + r.DstStart,
		toMove[1] - r.SrcStart + r.DstStart,
	}

	remaining := [2][2]int{
		[2]int{-1, -1},
		[2]int{-1, -1},
	}

	if toMove[0] != s[0] {
		remaining[0] = [2]int{s[0], toMove[0] - 1}
	}

	if toMove[1] != s[1] {
		remaining[1] = [2]int{toMove[1] + 1, s[1]}
	}

	return moved, remaining
}

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./input.txt")

	seeds := []int{}
	seedPairs := [][2]int{}

	seedRgx := regexp.MustCompile(`seeds: (.*)`)
	mappingRgx := regexp.MustCompile(`(.*) map:`)

	ranges := []Range{}

	for idx, line := range lines {
		if idx == 0 {
			seedsStr := seedRgx.FindStringSubmatch(line)[1]
			seedsSplit := strings.Split(seedsStr, " ")

			pair := [2]int{}
			for idx, seed := range seedsSplit {
				seedNum, _ := strconv.Atoi(seed)

				if part == 1 {
					seeds = append(seeds, seedNum)
				}

				if part == 2 {
					pid := idx % 2
					pair[pid] = seedNum
					if pid == 1 {
						pair[1] = pair[0] + seedNum - 1
						seedPairs = append(seedPairs, pair)
						pair = [2]int{}
					}
				}
			}

			continue
		}

		if len(line) < 2 {
			continue
		}

		mappingMatches := mappingRgx.FindStringSubmatch(line)
		if len(mappingMatches) > 1 {
			if len(ranges) > 0 {
				mappings = append(mappings, ranges)
				ranges = []Range{}
			}

			continue
		}

		rangeNumsStr := strings.Split(line, " ")
		rangeNums := lo.Map(rangeNumsStr, func(rangeNumStr string, idx int) int {
			rangeNum, _ := strconv.Atoi(rangeNumStr)
			return rangeNum
		})
		dstStart, srcStart, length := rangeNums[0], rangeNums[1], rangeNums[2]

		rng := Range{
			DstStart: dstStart,
			DstEnd:   dstStart + length - 1,
			SrcStart: srcStart,
			SrcEnd:   srcStart + length - 1,
		}

		ranges = append(ranges, rng)

		if idx == len(lines)-1 {
			mappings = append(mappings, ranges)
		}
	}

	answer = math.MaxInt

	if part == 1 {
		for _, seed := range seeds {
			for _, ranges := range mappings {
				for _, rng := range ranges {
					if seed >= rng.SrcStart && seed <= rng.SrcEnd {
						seed = seed - rng.SrcStart + rng.DstStart
						break
					}
				}
			}
			if seed < answer {
				answer = seed
			}
		}
	}

	if part == 2 {
		toDo := seedPairs
		done := [][2]int{}

		for _, ranges := range mappings {
			for sIdx := 0; sIdx < len(toDo); sIdx++ {
				s := toDo[sIdx]

				for _, r := range ranges {
					if s[0] <= r.SrcEnd && s[1] >= r.SrcStart { // they collide
						moved, remaining := collide(s, r)

						done = append(done, moved)

						if remaining[0][0] != -1 {
							toDo = append(toDo, remaining[0])
						}
						if remaining[1][0] != -1 {
							toDo = append(toDo, remaining[1])
						}

						toDo[sIdx] = [2]int{-1, -1}

						if remaining[0][0] == -1 && remaining[1][0] == -1 {
							break
						}
					}
				}
			}

			toDo = lo.Reject(toDo, func(item [2]int, idx int) bool {
				return item[0] == -1
			})
			toDo = append(toDo, done...)
			done = [][2]int{}

			sort.Slice(toDo, func(i, j int) bool {
				return toDo[i][0] < toDo[j][0]
			})
		}

		answer = toDo[0][0]
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
