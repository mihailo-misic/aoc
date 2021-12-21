package main

import (
	"fmt"
	"math"
	"sort"

	. "../utils"
)

type Coord struct {
	X int
	Y int
}

var risk int
var terrain [][]int
var openNodes []*Node
var closedNodes []*Node

func main() {
	input := ReadFile("./input.txt")

	for _, line := range input {
		row := []int{}
		for _, cr := range line {
			row = append(row, int(cr-'0'))

		}
		terrain = append(terrain, row)
	}

	expandTerrain()

	startNode := NewNode(Coord{0, 0}, nil)
	endNode := NewNode(Coord{len(terrain[0]) - 1, len(terrain) - 1}, nil)

	openNodes = append(openNodes, startNode)

	for len(openNodes) > 0 {
		fmt.Println("openNodes", len(openNodes))
		cNode := getCurrent()

		if cNode.Pos == endNode.Pos {
			endNode = cNode
			break // DONE
		}

		children := cNode.GetChildren()

		for _, child := range children {
			idxInOpen := child.IndexIn(openNodes)
			if idxInOpen == -1 {
				openNodes = append(openNodes, child)
				continue
			}

			if child.GetG() < openNodes[idxInOpen].GetG() {
				openNodes[idxInOpen].Parent = cNode
			}
		}
	}

	nextNode := endNode
	for nextNode != nil {
		risk += nextNode.V
		nextNode = nextNode.Parent
	}
	risk -= startNode.V

	fmt.Println("Answer", risk)
}

func expandTerrain() {
	/*
		evs := [][][]int{}
		evs = append(evs, terrain)

		for i := 1; i < 9; i++ {
			ev := [][]int{}

			for _, row := range terrain {
				evr := []int{}
				for _, num := range row {
					en := num + i
					if en > 9 {
						en = en - 9
					}
					evr = append(evr, en)
				}
				ev = append(ev, evr)
			}

			evs = append(evs, ev)
		}
		for i, ev := range evs {
			fmt.Println(i)
			for _, evr := range ev {
				fmt.Println(evr)
			}
		}
	*/

	newTer := [][]int{}
	for i := 0; i < len(terrain)*5; i++ {
		newTer = append(newTer, make([]int, len(terrain[0])*5))
	}

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			mut := r + c
			for tY, row := range terrain {
				nY := r*(len(terrain)) + tY
				for tX, num := range row {
					nX := c*(len(terrain[0])) + tX
					en := num + mut
					if en > 9 {
						en = en - 9
					}
					newTer[nY][nX] = en
				}
			}

		}
	}

	terrain = newTer
}

func (n *Node) GetRisk() (r int) {
	r = n.V

	p := n.Parent
	for p != nil {
		r += p.V
		p = p.Parent
	}

	return
}

func (n *Node) GetChildren() (children []*Node) {
	pos := n.Pos

	coords := []Coord{
		Coord{pos.X, pos.Y - 1}, // up
		Coord{pos.X, pos.Y + 1}, // down
		Coord{pos.X - 1, pos.Y}, // left
		Coord{pos.X + 1, pos.Y}, // right
	}

	for _, coord := range coords {
		child := NewNode(coord, n)

		if child == nil {
			continue
		}

		if child.IndexIn(closedNodes) != -1 {
			continue
		}

		children = append(children, child)
	}

	return
}

func (n *Node) IndexIn(list []*Node) int {
	for i, ln := range list {
		if n.Pos == ln.Pos {
			return i
		}
	}

	return -1
}

func getCurrent() *Node {
	sort.Slice(openNodes, func(i, j int) bool {
		return openNodes[i].GetF() < openNodes[j].GetF()
	})

	currentNode := openNodes[0]
	openNodes = openNodes[1:]
	closedNodes = append(closedNodes, currentNode)

	return currentNode
}

type Node struct {
	Pos    Coord
	V      int
	H      int
	Parent *Node
}

func NewNode(pos Coord, parent *Node) *Node {
	if pos.X < 0 || pos.X > len(terrain[0])-1 {
		return nil
	}
	if pos.Y < 0 || pos.Y > len(terrain)-1 {
		return nil
	}

	n := Node{
		Pos:    pos,
		V:      terrain[pos.Y][pos.X],
		Parent: parent,
	}
	n.SetH()

	return &n
}

func getDist(n1, n2 *Node) int {
	hD := math.Abs(float64(n1.Pos.X - n2.Pos.X))
	vD := math.Abs(float64(n1.Pos.Y - n2.Pos.Y))

	return int(math.Pow(hD, 2) + math.Pow(vD, 2))
}

func (n *Node) GetG() int {
	parentG := 0
	if n.Parent != nil {
		parentG = n.Parent.GetG()
	}

	return parentG + 1
}

func (n *Node) SetH() {
	endNode := &Node{
		Pos: Coord{
			X: len(terrain[0]) - 1,
			Y: len(terrain) - 1,
		},
	}
	n.H = getDist(n, endNode)
}

func (n *Node) GetF() int {
	return n.GetRisk()*9999 + n.GetG() + n.H
}
