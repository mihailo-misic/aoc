package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
)

var answer int

const part int = 2
const inputPath = "./input.txt"

type Span struct {
	From int
	To   int
}

func (aSpan *Span) Overlaps(bSpan Span) bool {
	return aSpan.From <= bSpan.To && bSpan.From <= aSpan.To
}

func (aSpan *Span) Merge(bSpan Span) Span {
	merged := Span{}

	if aSpan.From < bSpan.From {
		merged.From = aSpan.From
	} else {
		merged.From = bSpan.From
	}

	if aSpan.To > bSpan.To {
		merged.To = aSpan.To
	} else {
		merged.To = bSpan.To
	}

	return merged
}

func main() {
	defer util.Duration(util.Track("main"))

	lines := util.ReadFile(inputPath)

	mode := "spans"
	spans := []Span{}
	ids := []int{}

	for _, line := range lines {
		if line == "" {
			mode = "ids"
		}

		if mode == "spans" {
			spanSlice := strings.Split(line, "-")
			From, _ := strconv.Atoi(spanSlice[0])
			To, _ := strconv.Atoi(spanSlice[1])
			spans = append(spans, Span{From, To})
		}

		if mode == "ids" {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}

	slices.SortFunc(spans, func(x, y Span) int {
		return cmp.Compare(x.From, y.From)
	})

	if part == 1 {
		for _, id := range ids {
			for _, span := range spans {
				if id < span.From {
					break
				}

				if id <= span.To {
					answer++
					break
				}
			}
		}
	}

	if part == 2 {
		mergedSpans := []Span{}

		for _, aSpan := range spans {
			overlapped := false

			for bIdx, bSpan := range mergedSpans {

				if aSpan.Overlaps(bSpan) {
					mergedSpan := aSpan.Merge(bSpan)
					mergedSpans[bIdx] = mergedSpan
					overlapped = true

					continue
				}
			}

			if !overlapped {
				mergedSpans = append(mergedSpans, aSpan)
			}
		}

		for _, s := range mergedSpans {
			answer += s.To - s.From + 1
		}
	}

	util.CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}
