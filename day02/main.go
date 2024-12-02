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

func check_all_increasing(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		diff := list[i+1] - list[i]
		if !(diff >= 1 && diff <= 3) {
			return false
		}
	}
	return true
}

func check_all_decreasing(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		diff := list[i] - list[i+1]
		if !(diff >= 1 && diff <= 3) {
			return false
		}
	}
	return true
}

func check_allow_one_fault(list []int) bool {
	for i := range list {
		sub_list := make([]int, 0, len(list)-1)
		sub_list = append(sub_list, list[:i]...)
		sub_list = append(sub_list, list[i+1:]...)
		if check_all_increasing(sub_list) || check_all_decreasing(sub_list) {
			return true
		}
	}
	return false
}

func main() {

	readFile, err := os.Open("input.txt")
	defer readFile.Close()
	check(err)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var lists [][]int
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		nums := strings.Fields(line)
		lists = append(lists, []int{})
		for _, num := range nums {
			num_int, _ := strconv.Atoi(num)
			lists[i] = append(lists[i], num_int)
		}
		i++
	}
	sum1 := 0
	for _, list := range lists {
		if check_all_increasing(list) || check_all_decreasing(list) {
			sum1++
		}

	}

	sum2 := 0
	for _, list := range lists {
		if check_all_increasing(list) || check_all_decreasing(list) || check_allow_one_fault(list) {
			sum2++
		}

	}
	fmt.Printf("Sum part1: %d\n", sum1)
	fmt.Printf("Sum part2: %d\n", sum2)
}
