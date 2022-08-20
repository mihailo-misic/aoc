package main

import (
	"fmt"
	"regexp"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

var reg = regexp.MustCompile(".+x=([-\\d]+)\\.\\.([-\\d]+).+y=([-\\d]+)\\.\\.([-\\d]+)")

type Probe struct {
	X  int
	Y  int
	XV int
	YV int
}

func NewProbe(xVelocity, yVelocity int) *Probe {
	return &Probe{
		X:  0,
		Y:  0,
		XV: xVelocity,
		YV: yVelocity,
	}
}

var highestY int
var x, x1, y, y1 int

func main() {
	input := ReadFile("./input.txt")

	for _, line := range input {
		match := reg.FindStringSubmatch(line)
		x, _ = strconv.Atoi(match[1])
		x1, _ = strconv.Atoi(match[2])
		y1, _ = strconv.Atoi(match[3])
		y, _ = strconv.Atoi(match[4])
	}

	hitCount := 0
	for yV := 241; yV >= y1; yV-- {
		for xV := x1; xV >= -10; xV-- {
			hit := launch(xV, yV)
			if hit {
				hitCount++
			}
		}
	}

	fmt.Println("Answer:", hitCount, highestY)
}

func launch(xVelocity, yVelocity int) bool {
	p := NewProbe(xVelocity, yVelocity)
	fmt.Printf("%+v\n", *p)
	maxY := 0

	for {
		p.X += p.XV
		p.Y += p.YV
		if p.X > x1 || p.Y < y1 {
			return false
		}

		if p.Y > maxY { // bump maxY
			maxY = p.Y
		}

		if p.X >= x && p.X <= x1 && p.Y >= y1 && p.Y <= y { // is hit?
			if maxY > highestY {
				highestY = maxY
			}

			return true
		}

		if p.XV > 0 {
			p.XV--
		}
		if p.XV < 0 {
			p.XV++
		}
		p.YV--
	}
}
