package main

import (
	"fmt"

	. "../utils"
)

func main() {
	input := ReadFile("./sinput.txt")

	for _, line := range input {
		fmt.Println(line)
	}

	fmt.Println("Answer")
}
