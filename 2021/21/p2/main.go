package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/mihailo-misic/aoc/util"
)

type Player struct {
	Position int
	Score    int
}

type Game struct {
	Players   [2]Player
	TurnsLeft int
	PlayerIdx int
}

func (g *Game) Str() string {
	return fmt.Sprint(
		g.Players[0].Position,
		g.Players[0].Score,
		g.Players[1].Position,
		g.Players[1].Score,
		g.TurnsLeft,
		g.PlayerIdx,
	)
}

func (g *Game) Copy() Game {
	return Game{
		Players: [2]Player{
			{Position: g.Players[0].Position, Score: g.Players[0].Score},
			{Position: g.Players[1].Position, Score: g.Players[1].Score},
		},
		TurnsLeft: g.TurnsLeft,
		PlayerIdx: g.PlayerIdx,
	}
}

func (g *Game) Play(memo map[string][2]int) (p1Wins, p2Wins int) {
	memoKey := g.Str()
	if res, ok := memo[memoKey]; ok {
		return res[0], res[1]
	}

	if g.TurnsLeft == 0 {
		g.Players[g.PlayerIdx].Score += g.Players[g.PlayerIdx].Position

		if g.Players[g.PlayerIdx].Score >= GOAL {
			if g.PlayerIdx == 0 {
				return 1, 0
			}

			return 0, 1
		}

		g.TurnsLeft = 3
		g.PlayerIdx = g.PlayerIdx ^ 1
	}

	for roll := 1; roll <= 3; roll++ {
		newGame := g.Copy()
		newGame.TurnsLeft--

		newGame.Players[newGame.PlayerIdx].Position += roll
		if newGame.Players[newGame.PlayerIdx].Position > 10 {
			newGame.Players[newGame.PlayerIdx].Position -= 10
		}

		r1, r2 := newGame.Play(memo)
		p1Wins += r1
		p2Wins += r2
	}

	memo[memoKey] = [2]int{p1Wins, p2Wins}
	return
}

var answer int

const GOAL = 21

func main() {
	start := time.Now()
	lines := ReadFile("../input.txt")

	game := Game{TurnsLeft: 3}
	for i, line := range lines {
		s := strings.Split(line, ": ")

		if i == 0 {
			pos1, _ := strconv.Atoi(s[1])
			game.Players[0].Position = pos1
		}
		if i == 1 {
			pos2, _ := strconv.Atoi(s[1])
			game.Players[1].Position = pos2
		}
	}

	p1Wins, p2Wins := game.Play(map[string][2]int{})

	answer = p1Wins
	if p1Wins < p2Wins {
		answer = p2Wins
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)

	fmt.Println("Time:", time.Since(start))
}
