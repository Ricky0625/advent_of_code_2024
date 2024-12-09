package main

import (
	"fmt"

	"github.com/Ricky0625/advent_of_code_2024/utils"
)

func stringsToGrid(strings []string) [][]rune {
	grid := make([][]rune, len(strings))

	for i, str := range strings {
		grid[i] = []rune(str)
	}

	return grid
}

func lookForXmaxInDirection(grid [][]rune, rowIndex, colIndex int, dir [2]int) bool {
	word := "XMAS"
	dx := dir[0]
	dy := dir[1]

	// iterate through the word "XMAS"
	for i := 0; i < len(word); i++ {
		offsetRow, offsetCol := rowIndex+(dx*i), colIndex+(dy*i)

		// out of bound
		if offsetRow < 0 || offsetRow >= len(grid) || offsetCol < 0 || offsetCol >= len(grid[rowIndex]) {
			return false
		}

		// cannot form "XMAS"
		if grid[offsetRow][offsetCol] != rune(word[i]) {
			return false
		}

	}

	return true
}

/*
The grid is something like this

	0 1 2 3 4 5 6    x

0   M M M S X X M
1   M S A M X M S
2   A M X S X M A
3   M S A M A S M
4   X M A S A M X
5   X X A M M X X
6   S M S M S A S

y
*/
func main() {
	lines, err := utils.ReadLineByLine("input.txt")
	if err != nil {
		fmt.Println("Error opening file!")
	}

	count := 0
	grid := stringsToGrid(lines)
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{-1, 1},
		{-1, -1},
		{1, 1},
		{1, -1},
	}

	// iterate through each row
	for rowIndex, row := range grid {

		// iterate through each col in each row
		for colIndex := 0; colIndex < len(row); colIndex++ {

			// find xmas at the current [row,col] in 8 direction
			for _, dir := range directions {

				if isXmas := lookForXmaxInDirection(grid, rowIndex, colIndex, dir); isXmas {
					count++
				}

			}
		}
	}

	fmt.Printf("Day4 Part1: %v\n", count)
}
