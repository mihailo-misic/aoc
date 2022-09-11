package main

import (
	"fmt"
	"strconv"

	. "github.com/mihailo-misic/aoc/util"
)

const STEPS = 50

var algString string
var firstChar rune

func main() {
	lines := ReadFile("./input.txt")

	img := [][]rune{}

	for i, line := range lines {
		if i == 0 {
			algString = line
			firstChar = rune(algString[0])
			continue
		}

		if i == 1 {
			continue
		}

		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		img = append(img, row)
	}

	imgStr := ""
	for _, row := range img {
		for _, ch := range row {
			imgStr += string(ch)
		}
		imgStr += "\n"
	}
	fmt.Printf(">> Input <<\n%v\n", imgStr)

	for step := 0; step < STEPS; step++ {
		img = enhance(img, step)
		enhStr := ""
		for _, row := range img {
			for _, ch := range row {
				enhStr += string(ch)
			}
			enhStr += "\n"
		}
		fmt.Printf(">> Enh %v <<\n%v\n", step+1, enhStr)
	}

	litPixelsCount := 0
	for _, row := range img {
		for _, px := range row {
			if px == '#' {
				litPixelsCount++
			}
		}
	}

	CopyToClipboard(strconv.Itoa(litPixelsCount))
	fmt.Println("\nAnswer:", litPixelsCount)
}

func enhance(img [][]rune, step int) (enhancedImg [][]rune) {
	readImg := getReadImage(img, step)

	lastRow := len(readImg) - 1
	lastCol := len(readImg[0]) - 1

	enhancedImg = getWriteImage(img)
	for rIdx := 1; rIdx < lastRow; rIdx++ {
		for cIdx := 1; cIdx < lastCol; cIdx++ {
			pixels := getPixels(readImg, rIdx, cIdx)
			algIdx, _ := BinaryToDecimal(pixels)
			enhancedPixel := algString[algIdx]

			if enhancedPixel == '#' {
				enhancedImg[rIdx-1][cIdx-1] = '#'
			}
		}
	}

	return
}

func getPixels(readImg [][]rune, row, col int) (pixels string) {
	for rIdx := row - 1; rIdx <= row+1; rIdx++ {
		for cIdx := col - 1; cIdx <= col+1; cIdx++ {
			px := readImg[rIdx][cIdx]
			if px == '.' {
				pixels += "0"
			}
			if px == '#' {
				pixels += "1"
			}
		}
	}

	return
}

func getWriteImage(img [][]rune) (writeImg [][]rune) {
	for i := 0; i < len(img)+2; i++ {
		emptyRow := make([]rune, len(img[0])+2)
		for i := range emptyRow {
			emptyRow[i] = '.'
		}
		writeImg = append(writeImg, emptyRow)
	}

	return
}

func getReadImage(img [][]rune, step int) (readImg [][]rune) {
	infChar := '.'
	if step%2 == 1 && firstChar == '#' {
		infChar = '#'
	}

	for i := 0; i < 2; i++ {
		emptyRow := make([]rune, len(img[0])+4)
		for i := range emptyRow {
			emptyRow[i] = infChar
		}
		readImg = append(readImg, emptyRow)
	}

	for _, row := range img {
		readImgRow := []rune{infChar, infChar}

		for _, char := range row {
			readImgRow = append(readImgRow, char)
		}

		readImgRow = append(readImgRow, infChar, infChar)
		readImg = append(readImg, readImgRow)
	}

	for i := 0; i < 2; i++ {
		emptyRow := make([]rune, len(img[0])+4)
		for i := range emptyRow {
			emptyRow[i] = infChar
		}
		readImg = append(readImg, emptyRow)
	}

	return
}
