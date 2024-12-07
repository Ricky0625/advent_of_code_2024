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
	// apparently i need to read all as a string
	input, err := utils.ReadAll("input.txt")
	if err != nil {
		fmt.Printf("Failed to read input file: %v\n", err)
		return
	}

	var total int

	// same regex as part1
	// added regex rules to get do() and don't()
	reg, err := regexp.Compile(`mul\((\d{1,3},\d{1,3})\)|do\(\)|don't\(\)`)
	if err != nil {
		fmt.Println("Invalid regex expression.")
		return
	}

	// returns all matches; each includes full match and captured groups
	matches := reg.FindAllStringSubmatch(input, -1)

	// the idea here is to iterate through the slice of slices
	// each slices will have two value, one full match, one captured group
	// the regex above will get the do(), don't() and mul()
	// iterate through each slice at index 0, if toggle flag on if do(), off if don't()

	do := true
	for _, match := range matches {
		if match[0] == "do()" {
			do = true
			continue
		} else if match[0] == "don't()" {
			do = false
			continue
		} else if !do {
			continue
		}

		nums, err := GetNums(match[1])
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		total += nums[0] * nums[1]
	}

	fmt.Printf("Day3 Part2: %v\n", total)
}
