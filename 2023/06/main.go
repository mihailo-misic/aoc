package main

import (
	"fmt"
	"regexp"
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

	var times []int
	var time string
	var distances []int
	var distance string

	numRgx := regexp.MustCompile(`\d+`)

	for lIdx, line := range lines {
		numsStr := numRgx.FindAllString(line, -1)

		if lIdx == 0 {
			if part == 1 {
				times = lo.Map(numsStr, func(numStr string, idx int) int {
					num, _ := strconv.Atoi(numStr)
					return num
				})
			}

			if part == 2 {
				time = strings.Join(numsStr, "")
			}
		}

		if lIdx == 1 {
			if part == 1 {
				distances = lo.Map(numsStr, func(numStr string, idx int) int {
					num, _ := strconv.Atoi(numStr)
					return num
				})
			}

			if part == 2 {
				distance = strings.Join(numsStr, "")
			}
		}
	}

	if part == 2 {
		timeAsInt, _ := strconv.Atoi(time)
		times = []int{timeAsInt}
		distanceAsInt, _ := strconv.Atoi(distance)
		distances = []int{distanceAsInt}
	}

	fmt.Println("times", times)
	fmt.Println("distances", distances)

	for rIdx := 0; rIdx < len(times); rIdx++ {
		maxTime := times[rIdx]
		goal := distances[rIdx]
		wonTimes := 0

		for holdTime := 1; holdTime < maxTime; holdTime++ {
			if race(holdTime, goal, maxTime) {
				wonTimes++
			}
		}

		if answer == 0 {
			answer = 1
		}

		answer *= wonTimes
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func race(holdTime, goal, maxTime int) bool {
	remainingTime := maxTime - holdTime

	return holdTime*remainingTime > goal
}
