package main

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/mihailo-misic/aoc/util"
	"github.com/yourbasic/graph"
)

var answer int

/*
		00 01 02 03 04 05 06 07 08 09 10
	          11    15    19    23
	          12    16    20    24
	          13    17    21    25
			  14    18    22    26
*/

/*
		00 01 02 03 04 05 06 07 08 09 10
	          11    13    15    17
			  12    14    16    18
*/

const PART = 2 // 1 or 2

var TypeToRooms = map[string][]int{
	"A": {12, 11},
	"B": {14, 13},
	"C": {16, 15},
	"D": {18, 17},
}
var roomSize = 2

var typeToMoveCost = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

func main() {
	start := time.Now()
	lines := ReadFile(fmt.Sprintf("./input%v.txt", PART))

	pods := []Pod{}

	if PART == 2 {
		roomSize = 4
		TypeToRooms = map[string][]int{
			"A": {14, 13, 12, 11},
			"B": {18, 17, 16, 15},
			"C": {22, 21, 20, 19},
			"D": {26, 25, 24, 23},
		}

	}
	lineOffset := 0
	for lineIdx, line := range lines {
		addedOnLine := 0
		if lineIdx > 2 {
			lineOffset++
		}

		for _, char := range line {
			charStr := string(char)
			validChars := []string{"A", "B", "C", "D"}
			position := 11 + lineOffset + (addedOnLine * roomSize)

			if Includes(validChars, charStr) {
				pod := Pod{
					Id:       len(pods),
					Type:     charStr,
					Position: position,
					MoveCost: typeToMoveCost[charStr],
				}
				addedOnLine++
				pods = append(pods, pod)
			}
		}
	}

	situations.Push(pods)
	solve()

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Println("\nAnswer:", answer)

	fmt.Println("Duration", time.Since(start))
}

func compareFunc(a, b []Pod) bool {
	return getSituationCost(a) > getSituationCost(b)
}

// Min heap
var situations = GetNewHeap(compareFunc)
var highestSpent = 0

var seenSituations = map[string]bool{}

func GetSituationKey(pods *[]Pod) (key string) {
	for _, pod := range *pods {
		key += fmt.Sprint(pod.Id, pod.Position)
	}

	return
}

func solve() {
	// pop pods from situations heap
	pods, popped := situations.Pop()
	if !popped {
		return
	}

	// get movable pods
	movablePods := getMovablePods(pods)

	// for each movable pod get possible positions
	for _, pod := range movablePods {
		possiblePositions := pod.GetPossiblePositions(&pods)

		for _, pos := range possiblePositions {
			clonedPods := make([]Pod, len(pods))
			copy(clonedPods, pods)

			for idx, cp := range clonedPods {
				if cp.Id == pod.Id {
					// move the cloned pod in cloned pods
					clonedPods[idx].GoTo(pos, &clonedPods)

					situationKey := GetSituationKey(&clonedPods)
					if _, ok := seenSituations[situationKey]; ok {
						continue
					}
					seenSituations[situationKey] = true

					energySpent := getEnergySpent(&clonedPods)
					if energySpent > highestSpent {
						highestSpent = energySpent
					}

					Printiln("situations", situations.Size(), highestSpent)

					if isEverybodyHappy(&clonedPods) {
						answer = energySpent
						situations = &Heap[[]Pod]{}

						return
					}

					movPods := getMovablePods(clonedPods)
					if isSomebodyBlocked(&movPods) {
						continue
					}

					situations.Push(clonedPods)
				}
			}
		}
	}

	solve()
}

func getMovablePods(pods []Pod) (movablePods []Pod) {
	for _, pod := range pods {
		if pod.Position != pod.GetTarget(&pods) {
			movablePods = append(movablePods, pod)
		}
	}

	return
}

type Pod struct {
	Id          int
	Type        string
	Position    int
	MoveCost    int
	EnergySpent int
}

func (p *Pod) GetSiblings(pods *[]Pod) (siblings []Pod) {
	for _, pod := range *pods {
		if pod.Type == p.Type && pod.Id != p.Id {
			siblings = append(siblings, pod)
		}
	}

	return
}

func (p *Pod) GetAvailableRooms(pods *[]Pod) []int {
	siblings := p.GetSiblings(pods)

	rooms := TypeToRooms[p.Type]

	for roomIdx, room := range rooms {
		taken := false

		for _, sibling := range siblings {
			if sibling.Position == room {
				taken = true
				break
			}
		}

		if !taken {
			return rooms[roomIdx:]
		}
	}

	return []int{}
}

var targetMemo = map[string]int{}

func (p *Pod) GetTarget(pods *[]Pod) int {
	siblings := p.GetSiblings(pods)
	memoKey := p.GetMemoKey(&siblings)
	if res, ok := targetMemo[memoKey]; ok {
		return res
	}

	undetirmendPods := append(siblings, *p)
	rooms := TypeToRooms[p.Type]

	for _, room := range rooms {
		shortestPath := 1000
		closestPod := *p

		for _, pod := range undetirmendPods {
			pathLen := len(pod.GetPathTo(room))
			if pathLen < shortestPath {
				shortestPath = pathLen
				closestPod = pod
			}
		}

		if closestPod.Id == p.Id {
			targetMemo[memoKey] = room

			return room
		}

		newUndetirmendPods := []Pod{}
		for _, pod := range undetirmendPods {
			if pod.Id != closestPod.Id {
				newUndetirmendPods = append(newUndetirmendPods, pod)
			}
		}
		undetirmendPods = newUndetirmendPods
	}

	lastRoom := rooms[len(rooms)-1]
	targetMemo[memoKey] = lastRoom

	return lastRoom
}

func (p *Pod) GetOtherPods(pods *[]Pod) (otherPods []Pod) {
	for _, pod := range *pods {
		if pod.Id != p.Id {
			otherPods = append(otherPods, pod)
		}
	}

	return
}

func (p *Pod) GetWalkablePathTo(destination int, pods *[]Pod) (walkablePath []int) {
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

func (p *Pod) GoTo(destination int, pods *[]Pod) {
	path := p.GetWalkablePathTo(destination, pods)
	p.EnergySpent += p.MoveCost * len(path)
	p.Position = destination
}

var entrances = []int{2, 4, 6, 8}

const HALLWAY_LEFT = 0
const HALLWAY_RIGHT = 10

func (p *Pod) GetMemoKey(pods *[]Pod) (memoKey string) {
	memoKey = fmt.Sprint(p.Id, p.Position, p.EnergySpent)
	for _, pod := range *pods {
		if pod.Id != p.Id {
			memoKey += fmt.Sprint(pod.Id, pod.Position, pod.EnergySpent)
		}
	}

	return
}

var posMemo = map[string][]int{}

func (p *Pod) GetPossiblePositions(pods *[]Pod) (positions []int) {
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

var Graph = createGraph()

func getEnergySpent(pods *[]Pod) (energySpent int) {
	for _, pod := range *pods {
		energySpent += pod.EnergySpent
	}

	return
}

func isEverybodyHappy(pods *[]Pod) bool {
	for _, pod := range *pods {
		if pod.Position != pod.GetTarget(pods) {
			return false
		}
	}

	return true
}

var blockedMemo = map[string]bool{}

func isSomebodyBlocked(pods *[]Pod) bool {
	memoKey := getSituationMemoKey(*pods)
	if res, ok := blockedMemo[memoKey]; ok {
		return res
	}

	for _, pod := range *pods {
		if pod.EnergySpent > 0 {
			target := pod.GetTarget(pods)

			pathToTarget := pod.GetPathTo(target)
			otherPods := pod.GetOtherPods(pods)

			for _, otherPod := range otherPods {
				if otherPod.EnergySpent > 0 && Includes(pathToTarget, otherPod.Position) {
					otherPodTarget := otherPod.GetTarget(pods)

					if pod.GetDirection(target) != otherPod.GetDirection(otherPodTarget) {
						blockedMemo[memoKey] = true
						return true
					}
				}
			}
		}
	}
	blockedMemo[memoKey] = false

	return false
}

func (p *Pod) GetDirection(target int) string {
	pathToTarget := p.GetPathTo(target)
	nextNodeIdx := pathToTarget[0]

	if nextNodeIdx < p.Position {
		return "Left"
	}
	return "Right"
}

func getSituationMemoKey(pods []Pod) (memoKey string) {
	for _, pod := range pods {
		memoKey += fmt.Sprint(pod.Id, pod.Position, pod.EnergySpent)
	}

	return
}

var sitMemo = map[string]int{}

func getSituationCost(pods []Pod) (situationCost int) {
	memoKey := getSituationMemoKey(pods)
	if res, ok := sitMemo[memoKey]; ok {
		return res
	}

	for _, pod := range pods {
		situationCost += pod.EnergySpent

		target := pod.GetTarget(&pods)
		if pod.Position == target {
			continue
		}

		rooms := TypeToRooms[pod.Type]
		if Includes(rooms, pod.Position) {
			movementCount := roomSize - GetIndexOf(rooms, pod.Position) + 1
			situationCost += movementCount * pod.MoveCost
			continue
		}

		path := pod.GetPathTo(target)
		pathCost := pod.GetPathCost(path)
		situationCost += pathCost
	}

	sitMemo[memoKey] = situationCost

	return
}

var pathMemo = map[string][]int{}

func (p *Pod) GetPathTo(nodeId int) []int {
	memoKey := fmt.Sprint(p.Position, nodeId)
	if res, ok := pathMemo[memoKey]; ok {
		return res
	}
	path, _ := graph.ShortestPath(Graph, p.Position, nodeId)

	pathTo := path[1:]
	pathMemo[memoKey] = pathTo

	return pathTo
}

func (p *Pod) GetPathCost(path []int) (cost int) {
	moveCost := p.MoveCost

	cost += moveCost * len(path)

	return
}

func createGraph() *graph.Mutable {
	roomCount := 2
	if PART == 2 {
		roomCount = 4
	}

	Graph := graph.New(11 + 4*roomCount)

	// Add edges for hallways
	for i := 1; i <= 10; i++ {
		Graph.AddBothCost(i-1, i, 1)
	}

	// Add edges for rooms
	for i := 2; i <= 8; i += 2 {
		offset := 0
		if PART == 2 {
			offset = i - 2
		}

		Graph.AddBothCost(i, i+offset+9, 1)
		Graph.AddBothCost(i+offset+9, i+offset+10, 1)
		if PART == 2 {
			Graph.AddBothCost(i+offset+10, i+offset+11, 1)
			Graph.AddBothCost(i+offset+11, i+offset+12, 1)
		}
	}

	return Graph
}
