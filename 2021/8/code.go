package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	util "../utils"
)

func main() {
	input := readFile()
	ans := 0

	for _, line := range input {
		split := strings.Split(line, " | ")
		all := strings.Split(split[0], " ")

		digToPos := make(map[int][][]string)
		digs := make(map[int][]string)

		for _, dig := range all {
			digSl := strings.Split(dig, "")
			sort.Strings(digSl)
			if len(dig) == 2 { // 1
				digs[1] = digSl
			}
			if len(dig) == 3 { // 7
				digs[7] = digSl
			}
			if len(dig) == 4 { // 4
				digs[4] = digSl
			}
			if len(dig) == 5 { // 2 3 5
				digToPos[2] = append(digToPos[2], digSl)
				digToPos[3] = append(digToPos[3], digSl)
				digToPos[5] = append(digToPos[5], digSl)
			}
			if len(dig) == 6 { // 6 9 0
				digToPos[6] = append(digToPos[6], digSl)
				digToPos[9] = append(digToPos[9], digSl)
				digToPos[0] = append(digToPos[0], digSl)
			}
			if len(dig) == 7 { // 8
				digs[8] = digSl
			}
		}

		for _, dig := range digToPos[9] { // find 9
			if len(util.Intersect(dig, digs[4])) == len(digs[4]) {
				digs[9] = dig
				break
			}
		}

		for _, dig := range digToPos[0] { // find 0
			if len(util.Intersect(dig, digs[1])) == len(digs[1]) && !util.Equal(dig, digs[9]) {
				digs[0] = dig
				break
			}
		}

		for _, dig := range digToPos[6] { // find 6
			if !util.Equal(dig, digs[9]) && !util.Equal(dig, digs[0]) {
				digs[6] = dig
				break
			}
		}

		for _, dig := range digToPos[3] { // find 3
			if len(util.Intersect(dig, digs[1])) == len(digs[1]) {
				digs[3] = dig
				break
			}
		}

		for _, dig := range digToPos[5] { // find 5
			if !util.Equal(dig, digs[3]) {
				theSeg := util.Exclude(digs[1], digs[6])
				if len(util.Intersect(theSeg, dig)) == 0 {
					digs[5] = dig
					break
				}
			}
		}

		for _, dig := range digToPos[2] { // find 2
			if !util.Equal(dig, digs[3]) && !util.Equal(dig, digs[5]) {
				digs[2] = dig
				break
			}
		}

		lineNumSl := []int{}
		displayed := strings.Split(split[1], " ")
		for _, dig := range displayed {
			digSl := strings.Split(dig, "")
			sort.Strings(digSl)

			for n, kDig := range digs {
				if util.Equal(digSl, kDig) {
					lineNumSl = append(lineNumSl, n)
					break
				}
			}

		}

		x, _ := strconv.Atoi(strings.Trim(strings.Replace(fmt.Sprint(lineNumSl), " ", "", -1), "[]"))
		fmt.Println(x)
		ans += x
	}

	// 1013894 - too low
	fmt.Println(ans)
}

/*


















 */
func readFile() (input []string) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
