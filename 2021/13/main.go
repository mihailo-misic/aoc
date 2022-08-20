package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	. "github.com/mihailo-misic/aoc/util"
)

type Dot struct {
	X int
	Y int
}

type Fold struct {
	Axis   string
	Offset int
}

func main() {
	input := ReadFile("./input.txt")

	folds := []Fold{}
	dots := []*Dot{}
	yMax := 0
	xMax := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold") {
			fold := strings.Split(line, " ")[2]
			foldInfo := strings.Split(fold, "=")
			offset, _ := strconv.Atoi(foldInfo[1])
			folds = append(folds, Fold{foldInfo[0], offset})
			continue
		}

		dotInfo := strings.Split(line, ",")
		x, _ := strconv.Atoi(dotInfo[0])
		y, _ := strconv.Atoi(dotInfo[1])
		dots = append(dots, &Dot{x, y})

		if xMax < x {
			xMax = x
		}

		if yMax < y {
			yMax = y
		}
	}

	for _, fold := range folds {
		pap := [][]string{}
		for r := 0; r <= yMax; r++ {
			row := []string{}
			for c := 0; c <= xMax; c++ {
				row = append(row, " ")
			}
			pap = append(pap, row)
		}
		for _, dot := range dots {
			if fold.Axis == "x" && dot.X > fold.Offset {
				dot.X = fold.Offset*2 - dot.X
			}
			if fold.Axis == "y" && dot.Y > fold.Offset {
				dot.Y = fold.Offset*2 - dot.Y
			}

			pap[(*dot).Y][(*dot).X] = "@"
		}

		printLines("/tmp/13.txt", pap)
	}

	uniqDots := map[string]bool{}
	for _, dot := range dots {
		uniqDots[fmt.Sprintf("%d-%d", dot.X, dot.Y)] = true
	}

	fmt.Println("uniqDots", uniqDots)

	fmt.Println("Answer", len(uniqDots))
}

func printLines(filePath string, values [][]string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	v := []string{}
	for _, vr := range values {
		v = append(v, strings.Join(vr, ""))
	}

	for _, value := range v {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}
