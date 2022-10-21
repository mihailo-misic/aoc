package util

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"golang.org/x/exp/constraints"
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

func Includes[V comparable](slice []V, val V) bool {
	hMap := make(map[V]bool)
	for _, item := range slice {
		hMap[item] = true
	}

	return hMap[val] == true
}

func Unique[V comparable](slice []V) (uniqueSlice []V) {
	hMap := make(map[V]bool)

	for _, item := range slice {
		if _, ok := hMap[item]; !ok {
			hMap[item] = true
			uniqueSlice = append(uniqueSlice, item)
		}
	}

	return
}

func Merge[V comparable](slices [][]V) []V {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}

	mergedSlice := make([]V, totalLen)
	var idx int
	for _, s := range slices {
		idx += copy(mergedSlice[idx:], s)
	}

	return mergedSlice
}

func Intersect[V comparable](s1, s2 []V) (inter []V) {
	hMap := make(map[V]bool)
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

func Exclude[V comparable](s1, s2 []V) (inter []V) {
	hMap := make(map[V]bool)
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

func Equal[V comparable](a, b []V) bool {
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

func Min[V constraints.Ordered](slice []V) (min V) {
	min = slice[0]

	for _, value := range slice {
		if value < min {
			min = value
		}
	}

	return
}

func Max[V constraints.Ordered](slice []V) (max V) {
	max = slice[0]

	for _, value := range slice {
		if value > max {
			max = value
		}
	}

	return
}

func BinaryToDecimal(binary string) (decimal int64, err error) {
	decimal, err = strconv.ParseInt(binary, 2, 64)

	return
}

// Linux only via xclip
func CopyToClipboard(text string) error {
	command := exec.Command("xclip", "-in", "-selection", "clipboard")
	command.Stdin = bytes.NewReader([]byte(text))

	if err := command.Start(); err != nil {
		return fmt.Errorf("error starting xclip command: %w", err)
	}

	err := command.Wait()
	if err != nil {
		return fmt.Errorf("error running xclip %w", err)
	}

	return nil
}
