package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Function to check if there's an X pattern centered at (i, j)
func isXPattern(matrix [][]rune, i, j int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	// Top-left to bottom-right
	topLeftToBottomRight := i-1 >= 0 && j-1 >= 0 && i+1 < rows && j+1 < cols &&
		matrix[i-1][j-1] == 'M' && matrix[i][j] == 'A' && matrix[i+1][j+1] == 'S'

	// Bottom-left to top-right
	bottomLeftToTopRight := i+1 < rows && j-1 >= 0 && i-1 >= 0 && j+1 < cols &&
		matrix[i+1][j-1] == 'M' && matrix[i][j] == 'A' && matrix[i-1][j+1] == 'S'

	// Bottom-right to top-left
	bottomRightToTopLeft := i+1 < rows && j+1 < cols && i-1 >= 0 && j-1 >= 0 &&
		matrix[i+1][j+1] == 'M' && matrix[i][j] == 'A' && matrix[i-1][j-1] == 'S'

	// Top-right to bottom-left
	topRightToBottomLeft := i-1 >= 0 && j+1 < cols && i+1 < rows && j-1 >= 0 &&
		matrix[i-1][j+1] == 'M' && matrix[i][j] == 'A' && matrix[i+1][j-1] == 'S'

	// Check if any two intersecting diagonals form an X
	return (topLeftToBottomRight && bottomLeftToTopRight) ||
		(topLeftToBottomRight && topRightToBottomLeft) ||
		(bottomLeftToTopRight && bottomRightToTopLeft) ||
		(topRightToBottomLeft && bottomRightToTopLeft)
}

func main() {
	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	check(err)
	sum1 := 0
	sum2 := 0
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var charMatrix [][]rune
	for fileScanner.Scan() {
		line := fileScanner.Text()
		charRow := []rune(line)
		charMatrix = append(charMatrix, charRow)
	}
	for i, cRow := range charMatrix {
		for j := range cRow {
			// Boundary checks are crucial to prevent index out of range errors

			// Top
			if i >= 3 && j < len(cRow) {
				combined := string([]rune{charMatrix[i][j], charMatrix[i-1][j], charMatrix[i-2][j], charMatrix[i-3][j]})
				if combined == "XMAS" {
					sum1++
				}
			}

			// Top-left
			if i >= 3 && j >= 3 {
				combined := string([]rune{charMatrix[i][j], charMatrix[i-1][j-1], charMatrix[i-2][j-2], charMatrix[i-3][j-3]})
				if combined == "XMAS" {
					sum1++
				}
			}

			// Top-right
			if i >= 3 && j+3 < len(cRow) {
				combined := string([]rune{charMatrix[i][j], charMatrix[i-1][j+1], charMatrix[i-2][j+2], charMatrix[i-3][j+3]})
				if combined == "XMAS" {
					sum1++
				}
			}

			// Left
			if j >= 3 {
				combined := string([]rune{charMatrix[i][j], charMatrix[i][j-1], charMatrix[i][j-2], charMatrix[i][j-3]})
				if combined == "XMAS" {
					sum1++
				}
			}

			// Bottom-left
			if i+3 < len(charMatrix) && j >= 3 {
				combined := string([]rune{charMatrix[i][j], charMatrix[i+1][j-1], charMatrix[i+2][j-2], charMatrix[i+3][j-3]})
				if combined == "XMAS" {
					sum1++
				}
			}

			// Bottom
			if i+3 < len(charMatrix) {
				combined := string([]rune{charMatrix[i][j], charMatrix[i+1][j], charMatrix[i+2][j], charMatrix[i+3][j]})
				if combined == "XMAS" {
					sum1++
				}
			}

			// Bottom-right
			if i+3 < len(charMatrix) && j+3 < len(cRow) {
				combined := string([]rune{charMatrix[i][j], charMatrix[i+1][j+1], charMatrix[i+2][j+2], charMatrix[i+3][j+3]})
				if combined == "XMAS" {
					sum1++
				}
			}

			// Right
			if j+3 < len(cRow) {
				combined := string([]rune{charMatrix[i][j], charMatrix[i][j+1], charMatrix[i][j+2], charMatrix[i][j+3]})
				if combined == "XMAS" {
					sum1++
				}
			}

		}

	}

	for i, cRow := range charMatrix {
		for j := range cRow {
			if isXPattern(charMatrix, i, j) {
				sum2++
			}

		}

	}
	fmt.Printf("Sum part1: %d\n", sum1)
	fmt.Printf("Sum part2: %d\n", sum2)
}
