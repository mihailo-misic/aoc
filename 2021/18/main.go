package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	. "github.com/mihailo-misic/aoc/util"
)

type Node struct {
	L int
	R int
	H int
}

func NewNode(Left, Right, Height int) *Node {
	return &Node{Left, Right, Height}
}

func (n *Node) IsPair() bool {
	return n.L != -1 && n.R != -1
}

func main() {
	input := ReadFile("./input.txt")

	maxMagni := 0

	for a := range input {
		for b := range input {
			if a == b {
				continue
			}
			magni := run([]string{input[a], input[b]})
			if magni > maxMagni {
				maxMagni = magni
			}
		}
	}

	fmt.Println("magni", maxMagni)
}

func run(input []string) int {
	nodes := []*Node{}

	for i, line := range input {
		line = strings.Replace(line, ",", "", -1)

		nodes = append(nodes, lineToNodes(line)...)
		if i > 0 {
			for _, n := range nodes {
				n.H++
			}
		}

		nodes = reduce(nodes)
	}

	return getMagni(nodes)
}

func getMagni(nodes []*Node) (magni int) {
	maxH := 0
	for _, n := range nodes {
		if maxH < n.H {
			maxH = n.H
		}
	}

	for h := maxH; h >= 0; h-- {
		for i := 0; i < len(nodes); i++ {
			n := nodes[i]
			if n.H != h {
				continue
			}
			if n.IsPair() {
				n.H--
				n.L = 3*n.L + 2*n.R
				n.R = -1
			} else {
				if i == len(nodes)-1 {
					break
				}
				if n.R != -1 {
					n.L = n.R
					n.R = -1
				}
				n.H--
				n.L = 3*n.L + 2*nodes[i+1].L
				nodes = append(nodes[:i+1], nodes[i+2:]...)
			}
		}
	}

	return nodes[0].L
}

func reduce(nodes []*Node) []*Node {
	// look for explosions
	nodes = explode(nodes)

	// look for split
	nodes, split := split(nodes)
	if split {
		nodes = reduce(nodes)
	}

	// done? return
	return nodes
}

func split(nodes []*Node) ([]*Node, bool) {
	var node *Node
	splitIdx := -1
	for i, n := range nodes {
		if n.L > 9 || n.R > 9 {
			splitIdx = i
			node = n
			break
		}
	}
	if node == nil {
		return nodes, false
	}
	splitNum := node.R
	leftSplit := false
	if node.L > 9 {
		splitNum = node.L
		leftSplit = true
	}

	splitNode := NewNode(
		int(math.Floor(float64(splitNum)/2.0)),
		int(math.Ceil(float64(splitNum)/2.0)),
		node.H+1,
	)
	fmt.Println("\nSplit:", splitIdx)
	for _, n := range nodes {
		fmt.Printf("%+v ", *n)
	}
	fmt.Println()

	if node.IsPair() {
		if leftSplit {
			node.L = -1
			nodes = append(nodes[:splitIdx+1], nodes[splitIdx:]...)
			nodes[splitIdx] = splitNode
		} else {
			node.R = -1
			nodes = append(nodes[:splitIdx+1], nodes[splitIdx:]...)
			nodes[splitIdx+1] = splitNode
		}
	} else {
		nodes[splitIdx] = splitNode
	}

	for _, n := range nodes {
		fmt.Printf("%+v ", *n)
	}
	fmt.Println("")

	return nodes, true
}

func explode(nodes []*Node) []*Node {
	expIdx := -1
	for i, n := range nodes {
		if n.IsPair() && n.H > 3 {
			expIdx = i
			break
		}
	}
	if expIdx == -1 {
		return nodes
	}
	fmt.Println("\nExplode:", expIdx)
	for _, n := range nodes {
		fmt.Printf("%+v ", *n)
	}
	fmt.Println()

	expNode := nodes[expIdx]
	// Send L left
	merged := false
	var leftNode *Node
	if expIdx-1 >= 0 {
		leftNode = nodes[expIdx-1]
	}
	if leftNode != nil {
		if leftNode.R != -1 {
			leftNode.R += expNode.L
		} else {
			leftNode.L += expNode.L
			if leftNode.H == expNode.H-1 {
				leftNode.R = 0
				merged = true
			}
		}
	}

	// Send R right
	var rightNode *Node
	if expIdx+1 < len(nodes) {
		rightNode = nodes[expIdx+1]
	}
	if rightNode != nil {
		if rightNode.L != -1 {
			rightNode.L += expNode.R
		} else {
			rightNode.R += expNode.R
			if rightNode.H == expNode.H-1 {
				rightNode.L = 0
				merged = true
			}
		}
	}

	if !merged && leftNode != nil && leftNode.H == expNode.H {
		expNode.H--
		expNode.L = -1
		expNode.R = 0
	} else if !merged && rightNode != nil && rightNode.H == expNode.H {
		expNode.H--
		expNode.L = 0
		expNode.R = -1
	} else {
		nodes = append(nodes[:expIdx], nodes[expIdx+1:]...)
	}

	for _, n := range nodes {
		fmt.Printf("%+v ", *n)
	}
	fmt.Println()

	return explode(nodes)
}

func lineToNodes(line string) (nodes []*Node) {
	height := -1

	for i, r := range line {
		char := string(r)
		if char == "[" {
			height++
			continue
		}
		if char == "]" {
			height--
			continue
		}
		if unicode.IsDigit(r) {
			dig := int(r - '0')
			prevR := rune(line[i-1])
			if unicode.IsDigit(prevR) {
				nodes[len(nodes)-1].R = dig
				continue
			}

			prevChar := string(prevR)
			if prevChar == "[" {
				nodes = append(nodes, NewNode(dig, -1, height))
				continue
			}
			if prevChar == "]" {
				nodes = append(nodes, NewNode(-1, dig, height))
				continue
			}

		}
	}

	return
}
