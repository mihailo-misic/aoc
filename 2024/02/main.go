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

	rows := make([][]int, len(lines))
	for idx, line := range lines {
		numsStr := strings.Split(line, " ")
		nums := []int{}
		for _, num := range numsStr {
			n, _ := strconv.Atoi(num)
			nums = append(nums, n)
		}
		rows[idx] = nums
	}

	for _, row := range rows {
		safe := isRowSafe(row)

		if part == 2 && !safe {
			for i := 0; i < len(row); i++ {
				mr := removeIndex(row, i)
				safe = isRowSafe(mr)

				if safe {
					break
				}
			}
		}

		if safe {
			answer++
		}
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func removeIndex(s []int, index int) []int {
	ret := []int{}
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func isRowSafe(row []int) bool {
	inc := false

	for idx := 0; idx < len(row)-1; idx++ {
		c := row[idx]
		n := row[idx+1]

		if idx == 0 {
			inc = n > c
		}

		valid := isValid(inc, c, n)
		if !valid {
			return false
		}
	}

	return true
}

func isValid(inc bool, c, n int) bool {
	valid := map[int]bool{}

	if inc {
		valid[c+1] = true
		valid[c+2] = true
		valid[c+3] = true
	} else {
		valid[c-1] = true
		valid[c-2] = true
		valid[c-3] = true
	}

	return valid[n]
}
