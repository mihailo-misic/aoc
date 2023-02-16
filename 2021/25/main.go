package main

import (
	"fmt"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

var answer int

var yLimit int
var xLimit int

type BusyMap map[string]bool

var busyMap = make(BusyMap)

type Cucumber struct {
	Coordinates [2]int
	Direction   string
}

func (c *Cucumber) GetNextCoordinates() (nextY, nextX int) {
	nextY = c.Coordinates[0]
	nextX = c.Coordinates[1]

	if c.Direction == "Down" {
		nextY++
		if nextY > yLimit {
			nextY = 0
		}
	}
	if c.Direction == "Right" {
		nextX++
		if nextX > xLimit {
			nextX = 0
		}
	}

	return
}

func (c *Cucumber) Move() (moved bool) {
	nextY, nextX := c.GetNextCoordinates()

	busyMap[fmt.Sprint(c.Coordinates[0], c.Coordinates[1])] = false

	c.Coordinates[0] = nextY
	c.Coordinates[1] = nextX

	busyMap[fmt.Sprint(c.Coordinates[0], c.Coordinates[1])] = true

	return true
}

type Cucumbers []*Cucumber

func (cucumbers Cucumbers) GetMovable(direction string) (movableCucuambers Cucumbers) {
	for _, cucumber := range cucumbers {
		if cucumber.Direction == direction {
			nextY, nextX := cucumber.GetNextCoordinates()

			isBusy := busyMap[fmt.Sprint(nextY, nextX)]
			if !isBusy {
				movableCucuambers = append(movableCucuambers, cucumber)
			}
		}
	}

	return
}

func (cucumbers Cucumbers) Move() (moveCount int) {
	for idx, cucumber := range cucumbers {
		if cucumber.Direction == "Right" && cucumbers[idx].Move() {
			moveCount++
		}
	}

	for idx, cucumber := range cucumbers {
		if cucumber.Direction == "Down" && cucumbers[idx].Move() {
			moveCount++
		}
	}

	return
}

func (cucumbers Cucumbers) Print(yLimit, xLimit int) {
	board := make([][]string, yLimit+1)

	for y := 0; y <= yLimit; y++ {
		board[y] = make([]string, xLimit+1)
		for x := 0; x <= xLimit; x++ {
			board[y][x] = "."
		}
	}

	for _, c := range cucumbers {
		char := "v"
		if c.Direction == "Right" {
			char = ">"
		}

		board[c.Coordinates[0]][c.Coordinates[1]] = char
	}

	fmt.Println("\nStep", answer+1)
	for _, row := range board {
		fmt.Println(row)
	}
}

func main() {
	lines := ReadFile("./input.txt")

	cucumbers := Cucumbers{}

	yLimit = len(lines) - 1
	xLimit = len(lines[0]) - 1

	for y, line := range lines {
		for x, char := range line {
			if char == 'v' {
				cucumber := Cucumber{[2]int{y, x}, "Down"}
				cucumbers = append(cucumbers, &cucumber)
				busyMap[fmt.Sprint(y, x)] = true
			}

			if char == '>' {
				cucumber := Cucumber{[2]int{y, x}, "Right"}
				cucumbers = append(cucumbers, &cucumber)
				busyMap[fmt.Sprint(y, x)] = true
			}
		}
	}

	for {
		movableRightCucumbers := cucumbers.GetMovable("Right")
		movedCount := movableRightCucumbers.Move()

		movableDownCucumbers := cucumbers.GetMovable("Down")
		movedCount += movableDownCucumbers.Move()

		cucumbers.Print(yLimit, xLimit)

		answer++

		if movedCount == 0 {
			break
		}
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}
