package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) (input []string) {
	file, err := os.Open(path)
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

func Includes(strSlice []string, str string) bool {
	hMap := make(map[string]bool)
	for _, s := range strSlice {
		hMap[s] = true
	}

	return hMap[str] == true
}

func Intersect(s1, s2 []string) (inter []string) {
	hMap := make(map[string]bool)
	for _, s := range s1 {
		hMap[s] = true
	}
	for _, s := range s2 {
		if hMap[s] {
			inter = append(inter, s)
		}
	}

	return
}

func Exclude(s1, s2 []string) (inter []string) {
	hMap := make(map[string]bool)
	for _, s := range s2 {
		hMap[s] = true
	}
	for _, s := range s1 {
		if !hMap[s] {
			inter = append(inter, s)
		}
	}

	return
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}