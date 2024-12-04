package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func removeDontSections(input string) string {
	var result strings.Builder // Efficient string construction
	inDont := false            // Track whether we're inside a "don't()" block

	for i := 0; i < len(input); {
		// Check for "don't()"
		if strings.HasPrefix(input[i:], "don't()") {
			inDont = true
			i += len("don't()") // Move past "don't()"
			continue
		}

		// Check for "do()" and close the "don't()" block if inside
		if inDont && strings.HasPrefix(input[i:], "do()") {
			inDont = false
			i += len("do()") // Move past "do()"
			continue
		}

		// If not inside "don't()" block, append to result
		if !inDont {
			result.WriteByte(input[i])
		}
		i++
	}

	return result.String()
}

func mul_sum(lines string) int {
	sum := 0
	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(lines, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}
	return sum
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	check(err)

	sum1 := 0
	sum2 := 0
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var lines string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines += line

	}
	sum1 = mul_sum(lines)
	lines = removeDontSections(lines)
	sum2 = mul_sum(lines)

	fmt.Printf("Sum part1: %d\n", sum1)
	fmt.Printf("Sum part2: %d\n", sum2)
}
