package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	. "github.com/mihailo-misic/aoc/util"
)

type Player struct {
	Id       int
	Position int
	Score    int
	Throws   int
	History  []int
}

var answer int

var p1 = Player{Id: 1}
var p2 = Player{Id: 2}

var p1Wins int64 = 0
var p2Wins int64 = 0

const GOAL = 11

// 10 = 3.926645636s
// 11 = 5.623761788s
// 14 = 1h26m39.998895069s

func main() {
	start := time.Now()
	lines := ReadFile("../input.txt")

	for i, line := range lines {
		s := strings.Split(line, ": ")

		if i == 0 {
			p1.Position, _ = strconv.Atoi(s[1])
			p1.History = append(p1.History, p1.Position)
		}
		if i == 1 {
			p2.Position, _ = strconv.Atoi(s[1])
			p2.History = append(p2.History, p2.Position)
		}
	}

	// add 1-0 to open
	// add 2-0 to open
	// add 3-0 to open
	// while open != empty
	//   each open
	//     add x+1-0 to open
	//     add x+2-0 to open
	//     add x+3-0 to open
	//     remove that open (x-0)

	// Game
	//  p1 pos
	//  p1 sco
	//  p2 pos
	//  p2 sco
	//  turn

	rollDice(p1, p2, 1, 0)
	rollDice(p1, p2, 2, 0)
	rollDice(p1, p2, 3, 0)

	winner := p1
	winnerWins := p1Wins
	if p1Wins < p2Wins {
		winner = p2
		winnerWins = p2Wins
	}
	fmt.Printf("P%v wins with", winner.Id)
	answer = int(winnerWins)

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)

	fmt.Println("Time:", time.Since(start))
}

func rollDice(p1, p2 Player, diceNum, throwCount int) {
	throwerIdx := math.Mod(math.Floor(float64(throwCount)/3), 2)

	if throwerIdx == 0 {
		p1 = goTo(p1, diceNum)

		if p1.Score >= GOAL {
			p1Wins++
			//fmt.Printf("[1] Wins: %v, P: %+v\n", p1Wins, p1)
			return
		}
	}
	if throwerIdx == 1 {
		p2 = goTo(p2, diceNum)

		if p2.Score >= GOAL {
			p2Wins++
			//fmt.Printf("[2] Wins: %v, P: %+v\n", p2Wins, p2)
			return
		}
	}

	rollDice(p1, p2, 1, throwCount+1)
	rollDice(p1, p2, 2, throwCount+1)
	rollDice(p1, p2, 3, throwCount+1)
}

func goTo(p Player, diceNum int) (mP Player) {
	mP = Player{
		Id:       p.Id,
		Position: (p.Position+diceNum-1)%10 + 1,
		Score:    p.Score,
		Throws:   p.Throws + 1,
	}
	if mP.Throws == 3 {
		mP.Score += mP.Position
		mP.Throws = 0
	}
	mP.History = append(p.History, mP.Position)

	return mP
}
