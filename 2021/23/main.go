package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	. "github.com/mihailo-misic/aoc/util"
	"github.com/yourbasic/graph"
)

var answer int

type Pod struct {
	Id          int
	Type        string
	Position    int
	EnergySpent int
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
	p.EnergySpent += p.GetMoveCost() * len(path)
	p.Position = destination
}

var entrances = []int{2, 4, 6, 8}

const HALLWAY_LEFT = 0
const HALLWAY_RIGHT = 10

func (p *Pod) GetMemoKey(pods []Pod) (memoKey string) {
	memoKey = fmt.Sprint(p.Type, p.Id, p.Position, p.EnergySpent)
	for _, pod := range pods {
		if pod.Id != p.Id {
			memoKey += fmt.Sprint(pod.Type, pod.Id, pod.Position, pod.EnergySpent)
		}
	}

	return
}

var posMemo = map[string][]int{}

func (p *Pod) GetPossiblePositions(pods []Pod) (positions []int) {
	memoKey := p.GetMemoKey(pods)
	if res, ok := posMemo[memoKey]; ok {
		return res
	}
	paths := p.GetWalkablePathTo(p.GetTarget(pods), pods)

	if p.EnergySpent == 0 {
		pathToLeft := p.GetWalkablePathTo(HALLWAY_LEFT, pods)
		paths = append(paths, pathToLeft...)
		pathToRight := p.GetWalkablePathTo(HALLWAY_RIGHT, pods)
		paths = append(paths, pathToRight...)
	}

	uniquePositions := Unique(paths)

	for _, pos := range uniquePositions {
		if p.EnergySpent > 0 && p.GetTarget(pods) != pos {
			continue
		}
		if Includes(entrances, pos) { // Remove entrances
			continue
		}

		positions = append(positions, pos)
	}

	posMemo[memoKey] = positions

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
	start := time.Now()
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

	// TODO switch to breadth first with priority
	answer = solve(pods)

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)

	fmt.Println("Duration", time.Since(start))
}

var _lowest = math.MaxInt
var optionsOpen = 0

func solve(pods []Pod) int {
	energySpent := getEnergySpent(pods)

	// if all are in their room set answer and STOP
	if energySpent > _lowest || isEverybodyHappy(pods) {
		optionsOpen--
		return energySpent
	}

	var movablePods []Pod
	for _, pod := range pods {
		if pod.Position != pod.GetTarget(pods) {
			movablePods = append(movablePods, pod)
		}
	}

	lowestCost := math.MaxInt

	for _, pod := range movablePods {
		possiblePositions := pod.GetPossiblePositions(pods)

		for _, pos := range possiblePositions {
			optionsOpen++
			clonedPods := make([]Pod, len(pods))
			copy(clonedPods, pods)

			for idx, cp := range clonedPods {
				if cp.Id == pod.Id {
					clonedPods[idx].GoTo(pos, clonedPods)
					cost := solve(clonedPods)
					if cost < lowestCost {
						lowestCost = cost
						if lowestCost < _lowest {
							_lowest = lowestCost
							fmt.Println("new Lowest", _lowest)
						}
					}
					break
				}
			}
		}
	}

	optionsOpen--
	Printiln("optionsOpen", optionsOpen)
	return lowestCost
}

/*
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
*/

func getEnergySpent(pods []Pod) (energySpent int) {
	for _, pod := range pods {
		energySpent += pod.EnergySpent
	}

	return
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
		situationCost += pod.EnergySpent
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
