package main

import (
	"fmt"
	"strconv"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile("./input.txt")

	disk := []int{}

	for _, line := range lines {
		id := 0

		for idx, r := range line {
			isFile := idx%2 == 0
			amount := int(r - '0')

			for i := 0; i < amount; i++ {
				block := id

				if !isFile {
					block = -1
				}

				disk = append(disk, block)
			}

			if isFile {
				id++
			}
		}
	}

	var defragDisk []int

	if part == 1 {
		defragDisk = solve1(disk)
	}

	if part == 2 {
		defragDisk = solve2(disk, len(disk)-1)
	}

	for idx, num := range defragDisk {
		if num != -1 {
			answer += idx * num
		}
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func solve2(disk []int, rIdx int) []int {
	if rIdx <= 0 {
		return disk
	}

	fileEndIdx := 0
	fileStartIdx := 0

	// find file from rtl
	fNum := -1
	for ; rIdx >= 0; rIdx-- {
		num := disk[rIdx]

		if fNum == -1 {
			if num != -1 {
				fNum = num
				fileEndIdx = rIdx
				fileStartIdx = rIdx
			}

			continue
		}

		if fNum != num {
			break
		}

		fileStartIdx = rIdx
	}

	fileLen := fileEndIdx - fileStartIdx + 1

	// find space for file ltr
	spaceStartIdx := 1
	spaceEndIdx := 0
	sNum := -2
	for lIdx := 0; lIdx <= rIdx; lIdx++ {
		num := disk[lIdx]

		if sNum != -1 {
			if num == -1 {
				sNum = num
				spaceStartIdx = lIdx
				spaceEndIdx = lIdx
			}
		}

		if sNum == -1 {
			if num == -1 {
				spaceEndIdx = lIdx
			}

			if num != -1 {
				spaceStartIdx = 1
				spaceEndIdx = 0
				sNum = -2

				continue
			}
		}

		spaceLen := spaceEndIdx - spaceStartIdx + 1

		// move file to space
		if fileLen == spaceLen {
			for idx := spaceStartIdx; idx <= spaceEndIdx; idx++ {
				disk[idx] = fNum
			}

			for idx := fileStartIdx; idx <= fileEndIdx; idx++ {
				disk[idx] = -1
			}

			return solve2(disk, rIdx)
		}
	}

	return solve2(disk, rIdx)
}

func solve1(disk []int) []int {
	defragDisk := make([]int, len(disk))
	rIdx := len(disk) - 1

	for lIdx := 0; lIdx <= rIdx; lIdx++ {
		num := disk[lIdx]

		if num != -1 {
			defragDisk[lIdx] = num
			continue
		}

		if num == -1 {
			for {
				num := disk[rIdx]
				rIdx--

				if num != -1 {
					defragDisk[lIdx] = num
					break
				}
			}
		}
	}

	return defragDisk
}
