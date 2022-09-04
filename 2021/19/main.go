package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/mihailo-misic/aoc/util"
)

type Beacon struct {
	X int
	Y int
	Z int
}

type Vector struct {
	X int
	Y int
	Z int
}

var scanners = map[int][]Beacon{}

var globalBeacons = []Beacon{}

const REQUIRED_MATCH_COUNT = 12

func main() {
	lines := ReadFile("./input.txt")

	scanNum := 0
	for _, line := range lines {
		scanNum = readLines(line, scanNum)
	}

	scannersToGo := map[int][]Beacon{}

	for scannerNum, scan := range scanners {
		if scannerNum == 0 {
			addToGlobalBeacons(scan)
			continue
		}

		scannersToGo[scannerNum] = scan
	}

	ops := []func(Beacon) Beacon{
		turnNothing,
		turnRight, turnRight, turnRight, turnDown,
		turnRight, turnRight, turnRight, turnDown,
		turnRight, turnRight, turnRight, turnUp,
		turnRight, turnRight, turnRight, turnUp,
		turnRight, turnRight, turnRight, turnDown,
		turnRight, turnRight, turnRight,
	}

	for len(scannersToGo) > 0 {
		for scannerNum, scan := range scannersToGo {
			rotatedScan := scan
			for _, op := range ops {
				rotatedScan = rotate(rotatedScan, op)

				matchedScan := checkForOverlap(globalBeacons, rotatedScan)
				if matchedScan != nil {
					fmt.Println("Matched", scannerNum)
					addToGlobalBeacons(matchedScan)
					delete(scannersToGo, scannerNum)
					break
				}
			}
		}
	}

	fmt.Println("\nAnswer:", len(globalBeacons))
}

func rotate(scan []Beacon, rotation func(Beacon) Beacon) (rotatedScan []Beacon) {
	for _, beacon := range scan {
		rotatedScan = append(rotatedScan, rotation(beacon))
	}

	return
}

func turnNothing(b Beacon) Beacon {
	return b
}

func turnRight(b Beacon) Beacon {
	return Beacon{b.Z * -1, b.Y, b.X}
}

func turnDown(b Beacon) Beacon {
	return Beacon{b.X, b.Z, b.Y * -1}
}

func turnUp(b Beacon) Beacon {
	return Beacon{b.X, b.Z * -1, b.Y}
}

func addToGlobalBeacons(scan []Beacon) {
	for _, beacon := range scan {
		add := true
		for _, gBeacon := range globalBeacons {
			if gBeacon.isSameAs(beacon) {
				add = false
				break
			}
		}
		if add {
			globalBeacons = append(globalBeacons, beacon)
		}
	}
}

func checkForOverlap(scan1, scan2 []Beacon) []Beacon {
	for _, beacon1 := range scan1 {
		for _, beacon2 := range scan2 {
			moveVector := beacon2.getMoveVector(beacon1)
			movedScan := moveBeacons(scan2, moveVector)
			matchCount := getMatchCount(scan1, movedScan)
			if matchCount >= REQUIRED_MATCH_COUNT {
				return movedScan
			}
		}
	}

	return nil
}

func moveBeacons(scan1 []Beacon, moveVector Vector) []Beacon {
	movedScan := []Beacon{}
	for _, beacon := range scan1 {
		movedScan = append(movedScan, beacon.move(moveVector))
	}
	return movedScan
}

func (beacon Beacon) move(moveVector Vector) Beacon {
	return Beacon{
		X: beacon.X + moveVector.X,
		Y: beacon.Y + moveVector.Y,
		Z: beacon.Z + moveVector.Z,
	}
}

func getMatchCount(scan1, scan2 []Beacon) (matchCount int) {
	for _, beacon1 := range scan1 {
		for _, beacon2 := range scan2 {
			if beacon1.isSameAs(beacon2) {
				matchCount++
				break
			}
		}
	}

	return matchCount
}

func (beacon1 Beacon) isSameAs(beacon2 Beacon) bool {
	return beacon1.X == beacon2.X && beacon1.Y == beacon2.Y && beacon1.Z == beacon2.Z
}

func (beacon1 Beacon) getMoveVector(beacon2 Beacon) Vector {
	return Vector{
		X: delta(beacon1.X, beacon2.X),
		Y: delta(beacon1.Y, beacon2.Y),
		Z: delta(beacon1.Z, beacon2.Z),
	}
}

func delta(a, b int) int {
	if a > b {
		return (a - b) * -1
	}

	return b - a
}

func readLines(line string, scanNum int) int {
	if line == "" {
		return scanNum
	}

	scanNumRgx := regexp.MustCompile(`--- scanner (\d+) ---`)
	matches := scanNumRgx.FindStringSubmatch(line)
	if len(matches) > 1 {
		scanNum, _ = strconv.Atoi(matches[1])
		return scanNum
	}

	coords := strings.Split(line, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])

	scanners[scanNum] = append(scanners[scanNum], Beacon{x, y, z})

	return scanNum
}
