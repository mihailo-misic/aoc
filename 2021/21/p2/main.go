package main

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/mihailo-misic/aoc/util"
)

type Player struct {
	Id         int
	ThrowsLeft int
	Position   int
	Score      int
}

type Dice struct {
	TimesRolled int
	Number      int
}

var answer int

var p1 = Player{Id: 1}
var p2 = Player{Id: 2}

const GOAL = 1000

func main() {
	lines := ReadFile("../input.txt")

	d := Dice{0, 0}
	for i, line := range lines {
		s := strings.Split(line, ": ")

		if i == 0 {
			p1.Position, _ = strconv.Atoi(s[1])
		}
		if i == 1 {
			p2.Position, _ = strconv.Atoi(s[1])
		}
	}

	for {
		p1.ThrowsLeft = 3
		p2.ThrowsLeft = 3

		for p1.ThrowsLeft > 0 {
			p1.throwDice(&d)
		}
		if p1.Score >= GOAL {
			break
		}

		for p2.ThrowsLeft > 0 {
			p2.throwDice(&d)
		}
		if p2.Score >= GOAL {
			break
		}
	}

	loser := p1
	if p2.Score < p1.Score {
		loser = p2
	}
	answer = loser.Score * d.TimesRolled

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}

func (p *Player) throwDice(d *Dice) {
	d.roll()
	p.ThrowsLeft--

	p.Position += d.Number
	p.Position = p.Position % 10
	if p.Position == 0 {
		p.Position = 10
	}

	if p.ThrowsLeft == 0 {
		p.Score += p.Position
		fmt.Println(p.Id, "is at", p.Score)
	}
}

func (d *Dice) roll() {
	d.TimesRolled++
	d.Number++

	if d.Number > 100 {
		d.Number = 1
	}
}
