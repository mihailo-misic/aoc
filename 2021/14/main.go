package main

import (
	"fmt"
	"strings"
	"time"

	. "github.com/mihailo-misic/aoc/util"
)

var reactions = map[string]string{}

func main() {
	input := ReadFile("./input.txt")

	chain := ""

	for i, line := range input {
		if i == 0 {
			chain = line
			continue
		}
		if i == 1 {
			continue
		}

		lineInfo := strings.Split(line, " -> ")
		reactions[lineInfo[0]] = lineInfo[1]
	}

	minChar := "H"
	maxChar := "K"
	chainToNum := make(map[string]map[string]int)
	for in := range reactions {
		chainToNum[in] = make(map[string]int)

		react := getReactions(in, 20)

		for i, char := range react {
			if i == 0 || i == len(react)-1 {
				continue
			}
			cs := string(char)
			if cs == minChar || cs == maxChar {
				chainToNum[in][cs]++
			}
		}

		fmt.Println(chainToNum)
	}

	chainTo20 := getReactions(chain, 20)

	min := 0
	max := 0
	for i := 0; i < len(chainTo20)-1; i++ {
		key := string(chainTo20[i]) + string(chainTo20[i+1])
		nums := chainToNum[key]
		min += nums[minChar]
		max += nums[maxChar]
		if string(chainTo20[i]) == minChar {
			min++
		}
		if string(chainTo20[i]) == maxChar {
			max++
		}
	}

	fmt.Println("Answer", max-min)
}

func getReactions(chain string, steps int) string {
	for step := 0; step < steps; step++ {
		start := time.Now()
		newChain := ""

		chanMap := make(map[int](chan string))
		chunkedChain := chunkStr(chain, 1000)

		for i, chunk := range chunkedChain {
			chanMap[i] = make(chan string, 1)
			go buildChain(chunk, chanMap[i])
		}

		for i := 0; i < len(chanMap); i++ {
			s := <-chanMap[i]
			if i > 0 {
				s = s[1:]
			}
			newChain += s
		}

		chain = newChain
		fmt.Println("step", step, len(chain), time.Since(start))
	}

	return chain
}

func buildChain(chain string, c chan<- string) {
	newChain := ""
	for i := 0; i < len(chain)-1; i++ {
		key := string(chain[i]) + string(chain[i+1])
		newChain += string(chain[i]) + reactions[key]
		if i == len(chain)-2 {
			newChain += string(chain[i+1])
		}
	}

	c <- newChain
}

func chunkStr(str string, chunkSize int) (chunks []string) {
	a := []rune(str)

	for i := 0; i < len(a); i += chunkSize {
		end := i + chunkSize
		if end > len(a) {
			end = len(a)
		}
		start := i - 1
		if start < 0 {
			start = 0
		}

		chunks = append(chunks, string(str[start:end]))
	}

	return
}
