package util

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

// Starts the timer
func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Prints duration since provided time
func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

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

func GetIndexOf[V comparable](slice []V, val V) int {
	for idx := 0; idx < len(slice); idx++ {
		if val == slice[idx] {
			return idx
		}
	}

	return -1
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

func CopyToClipboard(text string) error {
	var command *exec.Cmd

	switch runtime.GOOS {
	case "darwin": // macOS
		command = exec.Command("pbcopy")
	case "linux":
		command = exec.Command("xclip", "-in", "-selection", "clipboard")
	default:
		return errors.New("unsupported platform")
	}

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

// Print text in the same terminal line
func Printiln(text ...any) {
	fmt.Print("\033[G\033[K")
	fmt.Println(text...)
	fmt.Print("\033[A")
}

func PrintSlice[V constraints.Ordered](slice []V) {
	for _, row := range slice {
		fmt.Println(row)
	}
}

func RemoveIndex[V constraints.Ordered](slice []V, index int) []V {
	newSlice := []V{}
	newSlice = append(newSlice, slice[:index]...)

	return append(newSlice, slice[index+1:]...)
}
