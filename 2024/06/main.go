package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
)

var answer int
var part int = 2

type Board [][]string

type Coord [2]int

type Guard struct {
	Position  Coord
	Direction string
	Memories  map[string]bool
}

func main() {
	defer util.Duration(util.Track("main"))

	g := Guard{
		Direction: "up",
		Memories:  map[string]bool{},
	}
	board := Board{}

	lines := util.ReadFile("./input.txt")

	for y, line := range lines {
		row := strings.Split(line, "")
		board = append(board, row)

		if g.Position[0] == 0 {
			for x, c := range row {
				if c == "^" {
					g.Position = [2]int{y, x}
					g.RememberPosition()
				}
			}
		}
	}
	startPos := g.Position

	if part == 1 {
		for {
			if err := g.Move(board); err != nil {
				break
			}
		}

		answer = len(g.GetUniquePositions())
	}

	if part == 2 {
		for {
			if err := g.Move(board); err != nil {
				break
			}
		}

		path := g.GetUniquePositions()

		loopChan := make(chan bool)

		for posStr := range path {
			go Simulate(board, startPos, posStr, loopChan)
		}

		for i := 0; i < len(path); i++ {
			if <-loopChan {
				answer++
			}
		}
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func Simulate(board Board, startPos Coord, posStr string, loopChan chan<- bool) {
	pos := strings.Split(posStr, "-")
	y, _ := strconv.Atoi(pos[0])
	x, _ := strconv.Atoi(pos[1])

	if y == startPos[0] && x == startPos[1] {
		loopChan <- false
		return
	}

	ghost := Guard{
		Position:  startPos,
		Direction: "up",
		Memories:  map[string]bool{},
	}
	ghostBoard := CopyBoard(board)
	ghostBoard[y][x] = "#"

	for {
		err := ghost.Move(ghostBoard)
		if err != nil {
			if err.Error() == "looping" {
				loopChan <- true
				return
			}

			break
		}
	}

	loopChan <- false
}

func (g *Guard) Move(board Board) error {
	next, err := g.GetNext(board)
	if err != nil {
		return errors.New("outside")
	}

	nextSpace := board[next[0]][next[1]]
	if nextSpace == "#" {
		g.ChangeDirection()

		return g.Move(board)
	}

	g.Position = next
	if g.IsLooping() {
		return errors.New("looping")
	}
	g.RememberPosition()

	return nil
}

func (g *Guard) GetNext(board Board) (next Coord, err error) {
	switch g.Direction {
	case "up":
		next = Coord{g.Position[0] - 1, g.Position[1]}
	case "down":
		next = Coord{g.Position[0] + 1, g.Position[1]}
	case "left":
		next = Coord{g.Position[0], g.Position[1] - 1}
	case "right":
		next = Coord{g.Position[0], g.Position[1] + 1}
	}

	if IsOutside(next, board) {
		return next, errors.New("Outside")
	}

	return next, nil
}

func (g *Guard) IsLooping() bool {
	return g.Memories[g.GetMemoKey()]
}

func (g *Guard) GetMemoKey() string {
	return fmt.Sprintf("%v-%v-%v", g.Position[0], g.Position[1], g.Direction)
}

func (g *Guard) RememberPosition() {
	g.Memories[g.GetMemoKey()] = true
}

func (g *Guard) GetUniquePositions() map[string]bool {
	positionMap := map[string]bool{}

	for key := range g.Memories {
		spl := strings.Split(key, "-")
		positionMap[spl[0]+"-"+spl[1]] = true
	}

	return positionMap
}

var dirToNextDir = map[string]string{
	"up":    "right",
	"right": "down",
	"down":  "left",
	"left":  "up",
}

func (g *Guard) ChangeDirection() {
	g.Direction = dirToNextDir[g.Direction]
}

func IsOutside(c Coord, b Board) bool {
	y, x := c[0], c[1]

	return y < 0 || y >= len(b) || x < 0 || x >= len(b[0])
}

func CopyBoard(b Board) Board {
	cb := make(Board, len(b))

	for idx, row := range b {
		cr := make([]string, len(row))
		copy(cr, row)
		cb[idx] = cr
	}

	return cb
}
