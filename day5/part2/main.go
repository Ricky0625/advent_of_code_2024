package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Ricky0625/advent_of_code_2024/utils"
)

func populateRulesAndUpdates(rules map[int][]int, updatesArr *[][]int, lines []string) {
	isRule := true

	for _, line := range lines {
		if len(line) == 0 {
			isRule = false
			continue
		}

		var nums []int

		if isRule {
			ruleSet := strings.Split(line, "|")
			nums, _ = utils.StringsToNums(ruleSet)
			rules[nums[0]] = append(rules[nums[0]], nums[1])
		} else {
			updates := strings.Split(line, ",")
			nums, _ = utils.StringsToNums(updates)
			*updatesArr = append(*updatesArr, nums)
		}
	}
}

func checkRules(rules map[int][]int, update []int) bool {
	for i := 1; i < len(update); i++ {
		target := update[i]
		subarr := append([]int{}, update[:i]...)

		// this list contains of the number that should appear after target
		list, exists := rules[target]
		if !exists {
			// ignore if not found
			continue
		}

		for i := range subarr {
			if slices.Contains(list, subarr[i]) {
				return false
			}
		}
	}

	return true
}

func main() {
	lines, _ := utils.ReadLineByLine("input.txt")
	rules := make(map[int][]int)
	updates := [][]int{}
	total := 0

	populateRulesAndUpdates(rules, &updates, lines)
	for _, update := range updates {
		isBad := false // flag to indicate if the current update needs fixing

		for j := range len(update) - 1 {

			for k := j + 1; k < len(update); k++ {
				list, exists := rules[update[j]]

				if !exists || !slices.Contains(list, update[k]) {
					isBad = true
					update[j], update[k] = update[k], update[j]

					k-- // recheck the same pair after swapping
					continue
				}
			}
		}

		if isBad {
			total += update[len(update)/2]
		}
	}

	fmt.Printf("Day5 Part1: %v\n", total)
}
