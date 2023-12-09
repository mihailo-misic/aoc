package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mihailo-misic/aoc/util"
	. "github.com/mihailo-misic/aoc/util"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

var answer int
var part int = 2

var cardToStrength = map[rune]string{
	'A': "z",
	'K': "y",
	'Q': "x",
	'J': "v",
	'T': "t",
	'9': "s",
	'8': "r",
	'7': "q",
	'6': "p",
	'5': "o",
	'4': "n",
	'3': "m",
	'2': "l",
}

type Player struct {
	Hand     string
	HandRank string
	Bid      int
}

func main() {
	defer util.Duration(util.Track("main"))

	lines := ReadFile("./input.txt")

	if part == 2 {
		cardToStrength['J'] = "k"
	}

	players := []Player{}

	for _, line := range lines {
		parsedLine := strings.Split(line, " ")

		hand := parsedLine[0]
		bid, _ := strconv.Atoi(parsedLine[1])

		players = append(players, Player{
			Hand:     hand,
			HandRank: rankHand(hand),
			Bid:      bid,
		})
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].HandRank < players[j].HandRank
	})

	for idx, player := range players {
		score := (idx + 1) * player.Bid
		answer += score
	}

	CopyToClipboard(strconv.Itoa(answer))
	fmt.Printf("\nAnswer (Part %v): %v\n", part, answer)
}

func getHandScore(hand string) string {
	hs := ""

	for _, r := range hand {
		hs += cardToStrength[r]
	}

	return hs
}

func rankHand(hand string) string {
	cardToCount := map[string]int{}

	hs := getHandScore(hand)

	jCount := 0

	for _, r := range hand {
		card := string(r)

		if card == "J" {
			jCount++
			continue
		}

		if _, has := cardToCount[card]; has {
			cardToCount[card]++
		} else {
			cardToCount[card] = 1
		}
	}

	maxCard := ""
	maxVal := 0

	for card, val := range cardToCount {
		if val >= maxVal {
			maxVal = val
			maxCard = card
		}
	}

	cardToCount[maxCard] += jCount

	// Five of a kind - 7
	if len(cardToCount) == 1 {
		return "7" + hs
	}
	if len(cardToCount) == 2 {
		// Full house - 5
		if lo.Min(maps.Values(cardToCount)) == 2 {
			return "5" + hs
		}
		// Four of a kind - 6
		return "6" + hs
	}

	if len(cardToCount) == 3 {
		// Two pair - 3
		if lo.Max(maps.Values(cardToCount)) == 2 {
			return "3" + hs
		}
		// Three of a kind - 4
		return "4" + hs
	}

	// One pair - 2
	if len(cardToCount) == 4 {
		return "2" + hs
	}
	// High card - 1
	if len(cardToCount) == 5 {
		return "1" + hs
	}

	return "0" + hs
}
