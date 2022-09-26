package main

import (
	"fmt"
	"regexp"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int

type Command struct {
	Action bool
	X      [2]int
	Y      [2]int
	Z      [2]int
}

func main() {
	lines := ReadFile("./xsinput.txt")

	input := make([]Command, len(lines))
	reg := regexp.MustCompile(`^(\w+) x\=(\d+)..(\d+),y\=(\d+)..(\d+),z\=(\d+)..(\d+)$`)

	for idx, line := range lines {
		res := reg.FindStringSubmatch(line)
		fmt.Println("res", res)
		XLower, _ := strconv.Atoi(res[2])
		XUpper, _ := strconv.Atoi(res[3])
		YLower, _ := strconv.Atoi(res[4])
		YUpper, _ := strconv.Atoi(res[5])
		ZLower, _ := strconv.Atoi(res[6])
		ZUpper, _ := strconv.Atoi(res[7])

		input[idx] = Command{
			Action: res[1] == "on",
			X:      [2]int{XLower, XUpper},
			Y:      [2]int{YLower, YUpper},
			Z:      [2]int{ZLower, ZUpper},
		}
	}
	for _, r := range input {
		fmt.Println(r)
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}
