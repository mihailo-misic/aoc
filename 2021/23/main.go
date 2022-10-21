package main

import (
	"fmt"
	"math"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
	"github.com/yourbasic/graph"
)

var answer int

type Pod struct {
	Id          int
	Type        string
	Position    int
	CurrentCost int
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

func (p *Pod) GetSibling(pods []Pod) Pod {
	for _, pod := range pods {
		if pod.Type == p.Type && pod.Id != p.Id {
			return pod
		}
	}

	return Pod{}
}

func (p *Pod) GetTarget(pods []Pod) int {
	sibling := p.GetSibling(pods)

	rooms := TypeToRooms[p.Type]
	if sibling.Position == rooms[0] {
		return rooms[1]
	}

	return rooms[0]
}

func (p *Pod) GetOtherPods(pods []Pod) (otherPods []Pod) {
	for _, pod := range pods {
		if pod.Id != p.Id {
			otherPods = append(otherPods, pod)
		}
	}

	return
}

func (p *Pod) GetPathTo(nodeId int) []int {
	path, _ := graph.ShortestPath(Graph, p.Position, nodeId)

	return path[1:]
}

func (p *Pod) GetWalkablePathTo(destination int, pods []Pod) (walkablePath []int) {
	fullPath := p.GetPathTo(destination)
	otherPods := p.GetOtherPods(pods)

	for _, pos := range fullPath {
		for _, otherPod := range otherPods {
			if pos == otherPod.Position {
				return
			}
		}

		walkablePath = append(walkablePath, pos)
	}

	return
}

func (p *Pod) GetPathCost(path []int, pods []Pod) (cost int) {
	moveCost := p.GetMoveCost()

	cost += moveCost * len(path)

	for i := 1; i < len(path); i++ {
		for _, pod := range pods {
			if pod.Position == path[i] {
				cost += moveCost * 2
			}
		}
	}

	return
}

func (p *Pod) GoTo(destination int, pods []Pod) {
	path := p.GetWalkablePathTo(destination, pods)
	p.CurrentCost += p.GetMoveCost() * len(path)
	p.Position = destination
}

var entrances = []int{2, 4, 6, 8}

const HALLWAY_LEFT = 0
const HALLWAY_RIGHT = 10

func (p *Pod) GetPossiblePositions(pods []Pod) (positions []int) {
	pathToTarget := p.GetWalkablePathTo(p.GetTarget(pods), pods)
	pathToLeft := p.GetWalkablePathTo(HALLWAY_LEFT, pods)
	pathToRight := p.GetWalkablePathTo(HALLWAY_RIGHT, pods)

	mergedPath := Merge([][]int{pathToTarget, pathToLeft, pathToRight})
	uniquePositions := Unique(mergedPath)

	// TODO Maybe prevent hallway to hallway movement
	// Remove entances
	for _, pos := range uniquePositions {
		if !Includes(entrances, pos) {
			positions = append(positions, pos)
		}
	}

	return
}

/*
		00 01 02 03 04 05 06 07 08 09 10
	         B11   C13   B15   D17
			 A12   D14   C16   A18
*/

var TypeToRooms = map[string][2]int{
	"A": {12, 11},
	"B": {14, 13},
	"C": {16, 15},
	"D": {18, 17},
}

var Graph = createGraph()

func main() {
	lines := ReadFile("./sinput.txt")

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

	solve(pods, -1)

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)
}

var _maxLvl = 6
var _curLvl = 0

var bestMoves = map[int]string{
	0: "B: 15 > 3",
	1: "C: 13 > 15",
	2: "D: 14 > 5",
	3: "B: 3 > 14",
	4: "B: 11 > 13",
	5: "D: 17 > 7",
	6: "A: 18 > 9",
	7: "D: 7 > 18",
	8: "D: 5 > 17",
	9: "A: 18 > 11",
}

func solve(pods []Pod, lastMovedPodId int) {
	if _curLvl > _maxLvl {
		return
	}

	// if all are in their room set answer and STOP
	if isEverybodyHappy(pods) {
		answer = getSituationCost(pods)
		return
	}

	var unhappyPods []Pod
	for _, pod := range pods {
		if pod.Position != pod.GetTarget(pods) && pod.Id != lastMovedPodId {
			unhappyPods = append(unhappyPods, pod)
		}
	}
	fmt.Println("unhappyPods", unhappyPods)

	lowestCost := math.MaxInt
	var bestPods []Pod
	var bestMove string
	bestPodId := -1

	for _, pod := range unhappyPods {
		possiblePositions := pod.GetPossiblePositions(pods)

		for _, pos := range possiblePositions {
			clonedPods := make([]Pod, len(pods))
			copy(clonedPods, pods)

			for idx, cp := range clonedPods {
				if cp.Id == pod.Id {
					clonedPods[idx].GoTo(pos, clonedPods)
					break
				}
			}

			sc := getSituationCost(clonedPods)
			if sc < lowestCost {
				bestMove = fmt.Sprintf("%v: %v > %v", pod.Type, pod.Position, pos)
				lowestCost = sc
				bestPods = clonedPods
				bestPodId = pod.Id
			}
		}
	}
	fmt.Println(bestMove == bestMoves[_curLvl], bestMove)

	_curLvl++
	solve(bestPods, bestPodId)
}

func isEverybodyHappy(pods []Pod) bool {
	for _, pod := range pods {
		if pod.Position != pod.GetTarget(pods) {
			return false
		}
	}

	return true
}

func getSituationCost(pods []Pod) (situationCost int) {
	for _, pod := range pods {
		path := pod.GetPathTo(pod.GetTarget(pods))
		cost := pod.GetPathCost(path, pods)
		situationCost += cost
		situationCost += pod.CurrentCost
		//fmt.Println(pod.Type, path, cost)
	}

	return situationCost
}

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
