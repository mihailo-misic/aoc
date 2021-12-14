package main

import (
	"fmt"
	"strings"

	. "../utils"
)

func main() {
	input := ReadFile("./sinput.txt")

	chain := ""
	reactions := map[string]string{}

	for i, line := range input {
		if i == 0 {
			chain = line
			continue
		}
		if i == 1 {
			continue
		}

		lineInfo := strings.Split(line, " -> ")
		fmt.Println(lineInfo)
		reactions[lineInfo[0]] = lineInfo[1]
	}

	fmt.Println("reactions", reactions)

	fmt.Println("chain", chain)
	for step := 0; step < 10; step++ {
		newChain := ""
		for i := 0; i < len(chain)-1; i++ {
			key := string(chain[i]) + string(chain[i+1])
			newChain += string(chain[i]) + reactions[key]
			if i == len(chain)-2 {
				newChain += string(chain[i+1])
			}
		}
		chain = newChain
		fmt.Println("chain", chain)
	}

	fmt.Println("Answer", chain)
}
