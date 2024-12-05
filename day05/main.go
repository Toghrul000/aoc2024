package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Topological sort
func reorderUpdate(update []int, rules map[int][]int) []int {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, page := range update {
		graph[page] = []int{}
		inDegree[page] = 0
	}

	for page1, dependencies := range rules {
		for _, page2 := range dependencies {
			if contains(update, page1) && contains(update, page2) {
				graph[page1] = append(graph[page1], page2)
				inDegree[page2]++
			}
		}
	}

	// Topological sorting using Kahn's algorithm
	var sorted []int
	queue := []int{}

	// Start with nodes that have 0 in-degree
	for page, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, page)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return sorted
}

func contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	check(err)

	sum1 := 0
	sum2 := 0
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	rules := make(map[int][]int)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.Contains(line, "|") {
			nums := strings.Split(line, "|")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])

			if _, ok := rules[num1]; !ok {
				rules[num1] = []int{}
			}
			rules[num1] = append(rules[num1], num2)

		} else {
			string_nums := strings.Split(line, ",")
			var nums []int
			for _, v := range string_nums {
				num, _ := strconv.Atoi(v)
				nums = append(nums, num)
			}

			positions := make(map[int]int)

			for i, v := range nums {
				positions[v] = i
			}
			correct := true
			for page1, should_come_after_nums := range rules {

				if _, ok := positions[page1]; !ok {
					continue
				}

				for _, page2 := range should_come_after_nums {

					if _, ok := positions[page2]; !ok {
						continue
					}

					if positions[page1] > positions[page2] {
						correct = false
					}
				}
			}
			if correct {
				sum1 += nums[len(nums)/2]
			} else {
				nums = reorderUpdate(nums, rules)
				sum2 += nums[len(nums)/2]
			}

		}

	}

	fmt.Printf("Sum part1: %d\n", sum1)
	fmt.Printf("Sum part2: %d\n", sum2)
}
