package main

import (
	"fmt"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
	"github.com/yourbasic/graph"
)

var answer int

type Pod struct {
	Id       int
	Type     string
	Position int
}

func (p *Pod) GetMoveCost() int {
	switch p.Type {
	case "A":
		return 1
	case "B":
		return 10
	case "C":
		return 100
	default:
		return 1000
	}
}

func (p *Pod) GetTarget(pods []Pod) int {
	var sibling Pod
	for _, pod := range pods {
		if pod.Id == p.Id {
			continue // Skip yourself
		}

		if pod.Type == p.Type && pod.Id != p.Id {
			sibling = pod
			break
		}
	}

	rooms := TypeToRooms[p.Type]
	if sibling.Position == rooms[0] {
		return rooms[1]
	}

	return rooms[0]
}

func (p *Pod) GetPathCost(path []int, pods []Pod) int {
	moveCost := p.GetMoveCost()

	podsInTheWay := []Pod{}
	for i := 1; i < len(path); i++ {
		position := path[i]
		for _, pod := range pods {
			if pod.Type != p.Type && pod.Position == position {
				podsInTheWay = append(podsInTheWay, pod)
			}
		}
	}

	movementCost := moveCost * (len(path) - 1)
	dodgeCost := moveCost * len(podsInTheWay) * 2

	return movementCost + dodgeCost
}

var TypeToRooms = map[string][2]int{
	"A": {12, 11},
	"B": {14, 13},
	"C": {16, 15},
	"D": {18, 17},
}

func main() {
	lines := ReadFile("./sinput.txt")

	Graph := createGraph()
	fmt.Println("Graph", Graph)

	pods := []Pod{}

	for lineIdx, line := range lines {
		if lineIdx == 2 || lineIdx == 3 {
			for charIdx, char := range line {
				charStr := string(char)
				validChars := []string{"A", "B", "C", "D"}

				if Includes(validChars, charStr) {
					offset := 2
					if lineIdx == 3 {
						offset = 1
					}
					pod := Pod{
						Id:       len(pods),
						Type:     charStr,
						Position: 10 + charIdx - offset,
					}
					pods = append(pods, pod)
				}
			}
		}
	}

	situationCost := 0
	for _, pod := range pods {
		target := pod.GetTarget(pods)
		path, _ := graph.ShortestPath(Graph, pod.Position, target)
		cost := pod.GetPathCost(path, pods)
		situationCost += cost
		fmt.Println(pod.Type, path, cost)
	}
	fmt.Println("situationCost", situationCost)

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}

/*
		00 01 02 03 04 05 06 07 08 09 10
	         B11   C13   B15   D17
			 A12   D14   C16   A18
*/

func createGraph() *graph.Mutable {
	Graph := graph.New(11 + 4*2)

	// Add edges for hallways
	for i := 1; i <= 10; i++ {
		Graph.AddBothCost(i-1, i, 1)
	}

	// Add edges for rooms
	for i := 2; i <= 8; i += 2 {
		Graph.AddBothCost(i, i+10-1, 1)
		Graph.AddBothCost(i+10-1, i+10, 1)
	}

	return Graph
}
