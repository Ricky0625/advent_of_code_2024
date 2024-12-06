package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ricky0625/advent_of_code_2024/utils"
)

func main() {
	lines, err := utils.ReadLineByLine("input.txt")
	if err != nil {
		fmt.Printf("Failed to read input file: %v\n", err)
		return
	}

	var total int
	var left []int
	right := make(map[int]int)

	for _, line := range lines {
		// spilt string into tokens, string.Fields is better than strings.Split
		tokens := strings.Fields(line)
		if len(tokens) != 2 {
			fmt.Printf("Invalid input line: %s\n", line)
			continue
		}

		leftNum, err := strconv.Atoi(tokens[0])
		if err != nil {
			fmt.Printf("Error parsing left number in line: %s\n", line)
			continue
		}

		rightNum, err := strconv.Atoi(tokens[1])
		if err != nil {
			fmt.Printf("Error parsing right number in line: %s\n", line)
			continue
		}

		left = append(left, leftNum)
		// add rightNumber to set, increment value (this is the occurence lah)
		right[rightNum] = right[rightNum] + 1
	}

	for i := range left {
		value, exists := right[left[i]]
		if exists {
			total += (left[i] * value)
		}
	}

	fmt.Printf("Day 1 Part 2: %v\n", total)
}
