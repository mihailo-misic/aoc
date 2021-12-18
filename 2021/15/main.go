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
	input := ReadFile("./sinput.txt")

	for _, line := range input {
		row := []int{}
		for _, cr := range line {
			row = append(row, int(cr-'0'))

		}
		terrain = append(terrain, row)
	}
	for _, row := range terrain {
		fmt.Println(row)
	}

	startNode := NewNode(Coord{0, 0}, nil)
	endNode := NewNode(Coord{len(terrain[0]) - 1, len(terrain) - 1}, nil)

	openNodes = append(openNodes, startNode)

	for len(openNodes) > 0 {
		cNode := getCurrent()
		fmt.Println(cNode)

		if cNode.Pos == endNode.Pos {
			fmt.Println("DONE")
			break // DONE
		}

		children := cNode.GetChildren()

		for _, child := range children {
			idxInOpen := child.IndexIn(openNodes)
			if idxInOpen != -1 && child.GetG() > openNodes[idxInOpen].GetG() {
				continue
			}

			openNodes = append(openNodes, child)
		}
	}

	// TODO backtrack from endNode
	nextNode := endNode
	for nextNode != nil {
		risk += nextNode.V
		nextNode = nextNode.Parent
	}

	fmt.Println(endNode)
	fmt.Println("Answer", risk)
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
	return n.V + n.GetG() + n.H
}
