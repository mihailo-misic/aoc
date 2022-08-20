package main

import "fmt"

var memo = make(map[int]int)

var DAYS = 256

func main() {
	train()

	fmt.Println(memo)
	grownPool := grow(input, DAYS/2)

	total := 0
	for _, f := range grownPool {
		done := memo[f]
		total += done
	}

	fmt.Println(total)
}

func train() {
	for i := 0; i <= 8; i++ {
		pool := []int{i}

		grownPool := grow(pool, DAYS/2)

		memo[i] = len(grownPool)
	}
}

func grow(pool []int, days int) []int {
	for day := 1; day <= days; day++ {
		poolLen := len(pool)

		for i := 0; i < poolLen; i++ {
			pool[i] = pool[i] - 1

			if pool[i] < 0 {
				pool[i] = 6
				pool = append(pool, 8)
			}
		}
	}

	return pool
}

var input = []int{1, 2, 1, 1, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 3, 1, 1, 1, 5, 1, 1, 1, 4, 5, 1, 1, 1, 3, 4, 1, 1, 1, 1, 1, 1, 1, 5, 1, 4, 1, 1, 1, 1, 1, 1, 1, 5, 1, 3, 1, 3, 1, 1, 1, 5, 1, 1, 1, 1, 1, 5, 4, 1, 2, 4, 4, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 5, 4, 3, 1, 1, 1, 1, 1, 1, 1, 5, 1, 3, 1, 4, 1, 1, 3, 1, 1, 1, 1, 1, 1, 2, 1, 4, 1, 3, 1, 1, 1, 1, 1, 5, 1, 1, 1, 2, 1, 1, 1, 1, 2, 1, 1, 1, 1, 4, 1, 3, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 4, 1, 1, 1, 1, 1, 3, 1, 3, 3, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 5, 1, 1, 1, 1, 2, 1, 1, 1, 4, 1, 1, 1, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 2, 3, 1, 2, 1, 1, 5, 4, 1, 1, 2, 1, 1, 1, 3, 1, 4, 1, 1, 1, 1, 3, 1, 2, 5, 1, 1, 1, 5, 1, 1, 1, 1, 1, 4, 1, 1, 4, 1, 1, 1, 2, 2, 2, 2, 4, 3, 1, 1, 3, 1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 4, 2, 1, 4, 1, 1, 1, 1, 1, 5, 1, 1, 4, 2, 1, 1, 2, 5, 4, 2, 1, 1, 1, 1, 4, 2, 3, 5, 2, 1, 5, 1, 3, 1, 1, 5, 1, 1, 4, 5, 1, 1, 1, 1, 4}
