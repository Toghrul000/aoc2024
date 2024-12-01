package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var list1, list2 []int

	for fileScanner.Scan() {
		line := fileScanner.Text()

		nums := strings.Fields(line)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	var sum_part1 float64
	for i := 0; i < len(list1); i++ {
		sum_part1 += math.Abs(float64(list1[i] - list2[i]))
	}

	fmt.Printf("Sum part1: %f\n", sum_part1)
	occurences := make(map[int]int)
	for _, num1 := range list1 {
		for _, num2 := range list2 {
			if num1 == num2 {
				occurences[num1]++
			}
		}
	}
	var sum_part2 int
	for k, v := range occurences {
		sum_part2 += k * v
	}

	fmt.Printf("Sum part2: %d\n", sum_part2)
	readFile.Close()
}
