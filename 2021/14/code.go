package main

import (
	"fmt"

	. "../utils"
)

func main() {
	input := ReadFile("./sinput.txt")
	ans := 0

	for _, line := range input {
		fmt.Println(line)
	}

	fmt.Println("Answer", ans)
}
