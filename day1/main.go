package main

import (
	"fmt"
	"math"
	"sort"
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
	var right []int

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
		right = append(right, rightNum)
	}

	// sort numberssss
	sort.Ints(left)
	sort.Ints(right)

	if len(left) != len(right) {
		fmt.Printf("Length of both left and right not the same!")
		return
	}

	for i := range len(left) {
		// ok, wtf? Go does not have Abs for Int?
		subtotal := math.Abs(float64(left[i] - right[i]))
		total += int(subtotal)
	}

	fmt.Printf("Day 1 Part 1: %v\n", total)
}
