package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	X int
	Y int
}

func main() {
	input := readFile()
	ans := 0

	matrix := [][]int{}
	for _, line := range input {
		lineS := strings.Split(line, "")
		lineI := []int{}
		for _, n := range lineS {
			nI, _ := strconv.Atoi(n)
			lineI = append(lineI, nI)
		}
		matrix = append(matrix, lineI)
	}

	lowPoints := []Coord{}
	for y, row := range matrix {
		for x, num := range row {
			if isLowest(matrix, num, x, y) {
				lowPoints = append(lowPoints, Coord{x, y})
			}
		}
	}

	basins := [][]Coord{}
	for _, lowPoint := range lowPoints {
		basins = append(basins, spread(matrix, lowPoint))
	}

	lens := []int{}
	for _, b := range basins {
		fmt.Println(b)
		lens = append(lens, len(b))
	}
	sort.Slice(lens, func(i, j int) bool {
		return lens[j] < lens[i]
	})

	fmt.Println(ans + lens[0]*lens[1]*lens[2])
}

func spread(matrix [][]int, lowPoint Coord) (points []Coord) {
	open := []Coord{lowPoint}

	findOpen(matrix, open, &points)

	return
}

// [{6 4} {6 3} {7 4} {7 3} {7 2} {8 4} {8 3} {5 4} {9 4}]
// [{6 4} {6 3} {7 4}             {8 4} {8 3} {5 4} {9 4} {9 4}]

func findOpen(matrix [][]int, open []Coord, points *[]Coord) {
	if len(open) == 0 {
		return
	}

	nopen := []Coord{}
	for _, p := range open {
		if p.Y-1 >= 0 {
			up := matrix[p.Y-1][p.X]
			uCrd := Coord{p.X, p.Y - 1}
			if up != 9 && getIndex(nopen, uCrd) == -1 && getIndex(*points, uCrd) == -1 {
				nopen = append(nopen, uCrd)
			}
		}
		if p.Y+1 < len(matrix) {
			down := matrix[p.Y+1][p.X]
			dCrd := Coord{p.X, p.Y + 1}
			if down != 9 && getIndex(nopen, dCrd) == -1 && getIndex(*points, dCrd) == -1 {
				nopen = append(nopen, dCrd)
			}
		}

		if p.X-1 >= 0 {
			left := matrix[p.Y][p.X-1]
			lCrd := Coord{p.X - 1, p.Y}
			if left != 9 && getIndex(nopen, lCrd) == -1 && getIndex(*points, lCrd) == -1 {
				nopen = append(nopen, lCrd)
			}
		}
		if p.X+1 < len(matrix[0]) {
			right := matrix[p.Y][p.X+1]
			rCrd := Coord{p.X + 1, p.Y}
			if right != 9 && getIndex(nopen, rCrd) == -1 && getIndex(*points, rCrd) == -1 {
				nopen = append(nopen, rCrd)
			}
		}

		*points = append(*points, p)
		//nopen = remove(nopen, p)
	}

	findOpen(matrix, nopen, points)
}

func getIndex(crds []Coord, crd Coord) int {
	for i, c := range crds {
		if c.X == crd.X && c.Y == crd.Y {
			return i
		}
	}

	return -1
}

func remove(crds []Coord, crd Coord) []Coord {
	if len(crds) <= 1 {
		return []Coord{}
	}
	idx := getIndex(crds, crd)

	return append(crds[:idx], crds[idx+1:]...)
}

func isLowest(matrix [][]int, num, x, y int) bool {
	up := 10
	if y-1 >= 0 && y-1 < len(matrix) {
		up = matrix[y-1][x]
	}

	down := 10
	if y+1 >= 0 && y+1 < len(matrix) {
		down = matrix[y+1][x]
	}

	left := 10
	if x-1 >= 0 && x-1 < len(matrix[0]) {
		left = matrix[y][x-1]
	}

	right := 10
	if x+1 >= 0 && x+1 < len(matrix[0]) {
		right = matrix[y][x+1]
	}

	return num < up && num < down && num < left && num < right
}

/*


















 */
func readFile() (input []string) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
