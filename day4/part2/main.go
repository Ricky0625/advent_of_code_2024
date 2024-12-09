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

func lookForMasInDiagonalDirection(grid [][]rune, rowIndex, colIndex int, diagSet [2][2]int) bool {
	offsetX1 := rowIndex + diagSet[0][0]
	offsetY1 := colIndex + diagSet[0][1]

	offsetX2 := rowIndex + diagSet[1][0]
	offsetY2 := colIndex + diagSet[1][1]

	if (grid[offsetX1][offsetY1] == 'M' && grid[offsetX2][offsetY2] == 'S') || (grid[offsetX1][offsetY1] == 'S' && grid[offsetX2][offsetY2] == 'M') {
		return true
	}

	return false
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

/*
For part 2, i probably can just skip the "outer layer"
So row 0, col 0, row len(grid)-1, col len(grid[row])-1 can skip
*/
func main() {
	lines, err := utils.ReadLineByLine("input.txt")
	if err != nil {
		fmt.Println("Error opening file!")
	}

	count := 0
	grid := stringsToGrid(lines)

	/*
		(TL)   (TR)
			 A
		(BL)   (BR)
		1. (TL)A(BR)
		2. (BR)A(TL)
		3. (TR)A(BL)
		4. (BL)A(TR)
	*/
	directions := [][2][2]int{
		{
			{-1, -1}, // (TL)
			{1, 1},   // (BR)
		},
		{
			{1, -1}, // (TR)
			{-1, 1}, // (BL)
		},
	}

	// if A appears at first row and last row, no X-MAS will be formed
	// iterate through each row but skip first and last row
	for rowIndex := 1; rowIndex < len(grid)-1; rowIndex++ {

		// if A appears at first col and last col, no X-MAS will be formed
		// iterate through each col in each row but skip first and last col
		for colIndex := 1; colIndex < len(grid[rowIndex])-1; colIndex++ {

			// use 'A' as starting point to locate X shape
			if grid[rowIndex][colIndex] != rune('A') {
				continue
			}

			isXmas := false

			// two diagonal direction should be true
			for _, dir := range directions {

				if isXmas = lookForMasInDiagonalDirection(grid, rowIndex, colIndex, dir); !isXmas {
					break
				}

			}

			if isXmas {
				count++
			}
		}
	}

	fmt.Printf("Day4 Part2: %v\n", count)
}
