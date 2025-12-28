package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")

	banks := [][]int{}
	for _, line := range lines {
		bankStr := strings.Split(line, "")
		bank := []int{}

		for _, batteryStr := range bankStr {
			battery, _ := strconv.Atoi(batteryStr)
			bank = append(bank, battery)
		}

		banks = append(banks, bank)
	}

	batteryCount := 2
	if part == 2 {
		batteryCount = 12
	}

	for _, bank := range banks {
		answer += findHighestJolts(bank, batteryCount, "")
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func findHighestJolts(bank []int, leftToFind int, joltsStr string) int {
	if leftToFind == 0 {
		jolts, _ := strconv.Atoi(joltsStr)

		return jolts
	}

	leftToFind--
	largest := 0
	largestIdx := 0

	for i := 0; i < len(bank)-leftToFind; i++ {
		bat := bank[i]

		if bat == 9 {
			largest = bat
			largestIdx = i
			break
		}

		if bat > largest {
			largest = bat
			largestIdx = i
		}
	}

	joltsStr += strconv.Itoa(largest)

	return findHighestJolts(bank[largestIdx+1:], leftToFind, joltsStr)
}
