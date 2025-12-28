package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
)

const part int = 2
const inputPath string = "./input.txt"

var answer int

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile(inputPath)

	rgx := regexp.MustCompile(" +")

	ops := []string{}
	splits := []int{}
	for idx, char := range lines[len(lines)-1] {
		if char != ' ' {
			ops = append(ops, string(char))
			if idx-1 > 0 {
				splits = append(splits, idx-1)
			}
		}
	}

	nums := make([][]int, len(ops))
	snums := [][]string{}

	for lIdx, line := range lines {
		if part == 1 {
			line = rgx.ReplaceAllString(line, "|")
			line = strings.Trim(line, "|")
			items := strings.Split(line, "|")

			for itemIdx, item := range items {
				num, err := strconv.Atoi(item)

				if err == nil {
					nums[itemIdx] = append(nums[itemIdx], num)
				}
			}
		}

		if part == 2 {
			if lIdx == len(lines)-1 {
				break
			}

			lineNums := []string{}
			lineNum := ""

			for idx, char := range line {

				if util.Includes(splits, idx) {
					lineNums = append(lineNums, lineNum)
					lineNum = ""
					continue
				}

				lineNum += string(char)
			}
			lineNums = append(lineNums, lineNum)
			snums = append(snums, lineNums)
		}
	}

	if part == 2 {
		for numIdx, snum := range snums[0] {
			for charIdx := len(snum) - 1; charIdx >= 0; charIdx-- {
				snum := ""
				for _, row := range snums {
					snum += string(row[numIdx][charIdx])
				}
				cleanNum := strings.Trim(snum, " ")
				num, _ := strconv.Atoi(cleanNum)
				nums[numIdx] = append(nums[numIdx], num)
			}
		}
	}

	for idx, op := range ops {
		colResult := 0

		for i, num := range nums[idx] {

			if i == 0 {
				colResult = num
				continue
			}

			switch op {
			case "*":
				colResult *= num
			case "+":
				colResult += num
			}
		}

		answer += colResult
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
