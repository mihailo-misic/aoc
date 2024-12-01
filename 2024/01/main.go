package main

import (
	"fmt"
	"math"
	"sort"
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

	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for idx, line := range lines {
		split := strings.Split(line, "   ")
		left[idx], _ = strconv.Atoi(split[0])
		right[idx], _ = strconv.Atoi(split[1])
	}
	sort.Ints(left)
	sort.Ints(right)

	for idx := 0; idx < len(left); idx++ {
		l := left[idx]
		r := right[idx]

		if part == 1 {
			answer += int(math.Abs(float64(l) - float64(r)))
		}

		if part == 2 {
			count := countInstances(l, right)
			answer += l * count
		}
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func countInstances(num int, nums []int) int {
	count := 0

	for _, n := range nums {
		if n == num {
			count++
		}
	}

	return count
}
