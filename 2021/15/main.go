package main

import (
	"container/heap"
	"fmt"
	"math"

	. "../utils"
)

type Coord struct {
	X int
	Y int
}

type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].GetF() < h[j].GetF() }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(n interface{}) {
	*h = append(*h, n.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[0 : n-1]

	return node
}

var terrain [][]int

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

	openNodes := &NodeHeap{}
	heap.Init(openNodes)
	heap.Push(openNodes, startNode)

	closedNodes := &NodeHeap{}
	heap.Init(closedNodes)

	fmt.Println(startNode.G)

	for _, r := range terrain {
		fmt.Println(r)
	}

	for openNodes.Len() > 0 {
		cNode := heap.Pop(openNodes).(*Node)
		heap.Push(closedNodes, cNode)
		fmt.Println(getDist(cNode, endNode))

		if cNode.Pos == endNode.Pos {
			endNode = cNode
			break // DONE
		}

		children := cNode.GetChildren(closedNodes)

		for _, child := range children {
			idxInOpen := child.IndexIn(openNodes)
			if idxInOpen == -1 {
				heap.Push(openNodes, child)
				continue
			}

			if child.G < (*openNodes)[idxInOpen].G {
				(*openNodes)[idxInOpen] = child
				heap.Fix(openNodes, idxInOpen)
			}
		}
	}

	pRisk := 0
	c := endNode
	for c.Parent != nil {
		pRisk += c.V
		c = c.Parent
	}
	fmt.Println("pRisk", pRisk)

	// 2977 -- too high
	// 3483 -- too high
	fmt.Println("Answer", endNode.Risk-startNode.V)
}

type Node struct {
	Pos    Coord
	V      int
	H      int
	G      int
	Risk   int
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
	n.SetG()    // Can change
	n.SetRisk() // Can change

	return &n
}

func (n *Node) GetChildren(closedNodes *NodeHeap) (children []*Node) {
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

func (n *Node) IndexIn(list *NodeHeap) int {
	for i, ln := range *list {
		if n.Pos == ln.Pos {
			return i
		}
	}

	return -1
}

func (n *Node) SetRisk() {
	if n.Parent == nil {
		n.Risk = n.V
	} else {
		n.Risk = n.Parent.Risk + n.V
	}
}

func getDist(n1, n2 *Node) int {
	hD := math.Abs(float64(n1.Pos.X - n2.Pos.X))
	vD := math.Abs(float64(n1.Pos.Y - n2.Pos.Y))

	return int(hD + vD)
}

func (n *Node) SetG() {
	if n.Parent == nil {
		n.G = 1
	} else {
		n.G = 1 + n.Parent.G
	}
	n.G += n.V * 1000
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
	return n.G + n.H
}

func expandTerrain() {
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
