package main

import (
	"fmt"
	"regexp"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int

type Vars map[string]int

type Cmd struct {
	Op    string
	Left  string
	Right string
}

/*

0  > 5
1  > 5
2  > 1
3  > 15
4  > 2
5  < -1
6  > 5
7  < -8
8  < -7
9  < -8
10 > 7
11 < -2
12 < -2
13 < -13

CORRECT
i5  = i4  +1
i7  = i6  -3
i8  = i3  +8
i9  = i2  -7
i11 = i10 +5
i12 = i1  +3
i13 = i0  -8

*/

/*

LOWEST
91811241914941 - too high
91811241911641


HIGHEST:
96916989924992 - too low
96918996924991 - just right
96919896924991 - too high

*/

var commands = []Cmd{}

func main() {
	lines := ReadFile("./input.txt")

	reg := regexp.MustCompile(`^(\w+) (\w) ?(.*)$`)
	for _, line := range lines {
		if line == "" {
			break
		}

		res := reg.FindStringSubmatch(line)
		commands = append(commands, Cmd{
			Op:    res[1],
			Left:  res[2],
			Right: res[3],
		})
	}

	grpCmds := splitCmds(commands)

	groupsMemo := make(map[int][]int)
	groupsMemo[14] = []int{0}

	for groupIdx := len(grpCmds) - 1; groupIdx >= 0; groupIdx-- {
		cmds := grpCmds[groupIdx]

		groupsMemo[groupIdx] = []int{}
		nextGroupMemo := groupsMemo[groupIdx+1]
		fmt.Println("\nnextGroupMemo", nextGroupMemo)

		resChan := make(chan []int, len(nextGroupMemo))

		for _, gNum := range nextGroupMemo {
			go func(goalNum int) {
				res := []int{}
				for i := 9; i >= 1; i-- {
					for z := 0; z < 100000; z++ {
						v := Vars{
							"w": 0,
							"x": 0,
							"y": 0,
							"z": z,
						}

						ok, _ := solve(i, cmds, v, goalNum)
						if ok {
							Printiln("Group:", groupIdx, "|", "Goal:", goalNum, "|", "Number:", i, "|", "Required Z:", z)
							if !Includes(res, z) {
								res = append(res, z)
							}
						}
					}
				}
				resChan <- res
			}(gNum)
		}

		for i := 0; i < len(nextGroupMemo); i++ {
			groupsMemo[groupIdx] = Merge([][]int{groupsMemo[groupIdx], <-resChan})
		}

		groupsMemo[groupIdx] = Unique(groupsMemo[groupIdx])
	}

	if answer != 0 {
		CopyToClipboard(strconv.Itoa(answer))
		fmt.Println("answer", answer)
	}
}

func solve(modelNum int, cmds []Cmd, vars Vars, goal int) (bool, Vars) {
	nums := IntToSlice(modelNum, []int{})
	for _, num := range nums {
		if num == 0 {
			return false, vars
		}
	}

	for _, cmd := range cmds {
		newNums, err := cmd.Exec(nums, &vars)
		if err != nil {
			vars["z"] = 1
			break
		}
		nums = newNums
	}

	return vars["z"] == goal, vars
}

func (cmd Cmd) Exec(nums []int, vars *Vars) ([]int, error) {
	left := (*vars)[cmd.Left]
	right := cmd.ParseRight(vars)

	switch cmd.Op {
	case "inp":
		num, nums := nums[0], nums[1:]
		(*vars)[cmd.Left] = num

		return nums, nil

	case "add":
		(*vars)[cmd.Left] += right

	case "mul":
		(*vars)[cmd.Left] *= right

	case "div":
		/*
			if right == 0 {
				return nums, errors.New("Invalid division")
			}

			val := math.Round(float64(left) / float64(right))
			(*vars)[cmd.Left] = int(val)
		*/
		(*vars)[cmd.Left] /= right

	case "mod":
		/*
			if left < 0 || right <= 0 {
				return nums, errors.New("Invalid modulus")
			}

			val := left % right
			(*vars)[cmd.Left] = val
		*/
		(*vars)[cmd.Left] %= right

	case "eql":
		if left == right {
			(*vars)[cmd.Left] = 1
		} else {
			(*vars)[cmd.Left] = 0
		}
	}

	return nums, nil
}

func (cmd Cmd) ParseRight(vars *Vars) int {
	num, err := strconv.Atoi(cmd.Right)
	if err != nil {
		return (*vars)[cmd.Right]
	}

	return num
}

func splitCmds(cmds []Cmd) (groupedCmds [][]Cmd) {
	for i := 0; i < 14; i++ {
		idx := i * 18
		groupedCmds = append(groupedCmds, cmds[idx:idx+18])
	}

	return
}

/*
Helpers



*/

func IntToSlice(n int, seq []int) []int {
	if n != 0 {
		seq = append([]int{n % 10}, seq...)

		return IntToSlice(n/10, seq)
	}

	return seq
}

func concatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}
