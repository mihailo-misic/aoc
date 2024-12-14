package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

type Equation struct {
	Goal int
	Nums []int
}

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")
	equations := []Equation{}

	for _, line := range lines {
		split := strings.Split(line, ": ")

		goal, _ := strconv.Atoi(split[0])

		numsStr := strings.Split(split[1], " ")
		nums := make([]int, len(numsStr))
		for i := 0; i < len(numsStr); i++ {
			nums[i], _ = strconv.Atoi(numsStr[i])
		}

		equations = append(equations, Equation{goal, nums})
	}

	for _, eq := range equations {
		if solve(eq, 0, 0) {
			answer += eq.Goal
		}
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func solve(eq Equation, cur int, idx int) bool {
	if idx == 0 {
		return solve(eq, eq.Nums[idx], idx+1)
	}

	if idx < len(eq.Nums) {
		n := eq.Nums[idx]
		if part == 1 {
			return solve(eq, cur+n, idx+1) || solve(eq, cur*n, idx+1)
		}

		if part == 2 {
			conc := fmt.Sprintf("%v%v", cur, n)
			concInt, _ := strconv.Atoi(conc)

			return solve(eq, cur+n, idx+1) || solve(eq, cur*n, idx+1) || solve(eq, concInt, idx+1)
		}
	}

	return cur == eq.Goal
}
