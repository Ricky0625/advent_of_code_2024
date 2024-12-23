package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Ricky0625/advent_of_code_2024/utils"
)

func GetNums(str string) ([]int, error) {
	nums := make([]int, 2)
	tokens := strings.Split(str, ",")

	if len(tokens) != 2 {
		return nil, fmt.Errorf("Invalid size of tokens")
	}

	for i, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert token to num: %v\n", token)
		}
		nums[i] = num
	}

	return nums, nil
}

func main() {
	lines, err := utils.ReadLineByLine("input.txt")
	if err != nil {
		fmt.Printf("Failed to read input file: %v\n", err)
		return
	}

	var total int

	// define regex
	// \d means number, {} means range
	// \d{1,3} means at least 1 digit, at most 3 digit
	// do note that the x,y is wrapped with a pair of parentheses (captured group)
	reg, err := regexp.Compile(`mul\((\d{1,3},\d{1,3})\)`)
	if err != nil {
		fmt.Println("Invalid regex expression.")
		return
	}

	for _, line := range lines {
		// returns all matches; each includes full match and captured groups
		matches := reg.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if len(match) < 2 {
				fmt.Println("No captured groups!")
				break
			}
			nums, err := GetNums(match[1])
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			total += nums[0] * nums[1]
		}
	}

	fmt.Printf("Day3 Part1: %v\n", total)
}
