package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int

func groupNumbers(cmds []int) (groupedNumbers [][]int) {
	for i := 0; i < 14; i++ {
		idx := i * 18
		groupedNumbers = append(groupedNumbers, cmds[idx:idx+18])
	}

	return
}

const PART = 1

func main() {
	lines := ReadFile("./input.txt")

	numbers := []int{}

	reg := regexp.MustCompile(`^\w+ \w ?(.*)$`)

	for _, line := range lines {
		if line == "" {
			break
		}

		matches := reg.FindStringSubmatch(line)
		number, _ := strconv.Atoi(matches[1])
		numbers = append(numbers, number)
	}

	modelNumber := make([]int, 14)
	stack := [][2]int{}
	groupedCmds := groupNumbers(numbers)

	for groupIdx, group := range groupedCmds {
		puller, pusher := group[5], group[15]

		if puller > 0 {
			stack = append(stack, [2]int{groupIdx, pusher})
			continue
		}

		var popped [2]int
		popped, stack = stack[len(stack)-1], stack[:len(stack)-1]
		poppedGroupIdx, poppedNumber := popped[0], popped[1]

		modifier := poppedNumber + puller

		if PART == 1 {
			if modifier > 0 {
				modelNumber[groupIdx] = 9
				modelNumber[poppedGroupIdx] = 9 - modifier
			}

			if modifier < 0 {
				modelNumber[groupIdx] = 9 + modifier
				modelNumber[poppedGroupIdx] = 9
			}
		}

		if PART == 2 {
			if modifier > 0 {
				modelNumber[groupIdx] = 1 + modifier
				modelNumber[poppedGroupIdx] = 1
			}

			if modifier < 0 {
				modelNumber[groupIdx] = 1
				modelNumber[poppedGroupIdx] = 1 - modifier
			}
		}
	}

	answer, _ = strconv.Atoi(strings.Trim(strings.Replace(fmt.Sprint(modelNumber), " ", "", -1), "[]"))

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("answer", answer)
}
