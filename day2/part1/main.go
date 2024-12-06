package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Ricky0625/advent_of_code_2024/utils"
)

// convert strings into ints
func StringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("Error parsing number: %s\n", str)
		}
		ints[i] = num
	}
	return ints, nil
}

func IsValidReport(arr []int) (bool, error) {
	if len(arr) < 2 {
		return false, fmt.Errorf("array size too small!")
	}

	// determine whether the list should be increasing or decreasing
	isIncreasing := (arr[1] - arr[0]) > 0

	for i := 0; i < len(arr)-1; i++ {
		// 1 3 5 7 9. diff between them should be negative
		// 9 7 5 3 2. diff between them should be positive
		difference := arr[i] - arr[i+1]

		if (isIncreasing && difference > 0) || (!isIncreasing && difference < 0) {
			return false, nil
		}

		// difference should be >=1 and <=3
		if absDiff := int(math.Abs(float64(difference))); !(absDiff >= 1 && absDiff <= 3) {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	lines, err := utils.ReadLineByLine("input.txt")
	if err != nil {
		fmt.Printf("Failed to read input file: %v\n", err)
		return
	}

	var total int

	for _, line := range lines {
		strings := strings.Fields(line)
		ints, err := StringsToInts(strings)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		valid, err := IsValidReport(ints)
		if err != nil {
			fmt.Printf("Failed to check: %v\n", ints)
			break
		}
		if valid {
			total += 1
		}
	}

	fmt.Println("Day 2 Part 1: ", total)
}
