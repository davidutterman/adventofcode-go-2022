package day8

import (
	"adventofcode-go-2022/util"
	"strings"
)

func Part1() int {
	grid := parseInput()

	visibleThrees := map[[2]int]string{}

	for column := 1; column < len(grid)-1; column++ {

		// from left
		highestFromLeft := -1
		for row := 0; row < len(grid[0])-1; row++ {
			height := grid[column][row]
			if height > highestFromLeft {
				visibleThrees[[2]int{column, row}] = "L"
				highestFromLeft = height
			}
		}

		// from right
		highestFromRight := -1
		for row := len(grid[0]) - 1; row > 0; row-- {
			height := grid[column][row]
			if height > highestFromRight {
				visibleThrees[[2]int{column, row}] = "R"
				highestFromRight = height
			}
		}
	}

	for column := 1; column < len(grid[0])-1; column++ {

		// from top
		highestFromTop := -1
		for row := 0; row < len(grid)-1; row++ {
			height := grid[row][column]
			if height > highestFromTop {
				visibleThrees[[2]int{row, column}] = "T"
				highestFromTop = height
			}
		}

		// from bottom
		highestFromBottom := -1
		for row := len(grid) - 1; row > 0; row-- {
			height := grid[row][column]
			if height > highestFromBottom {
				visibleThrees[[2]int{row, column}] = "B"
				highestFromBottom = height
			}
		}
	}

	return len(visibleThrees) + 4 // plus 4 for corners
}

func Part2() int {
	grid := parseInput()
	sum := 0
	for row := 1; row < len(grid)-1; row++ {
		for column := 1; column < len(grid[0])-1; column++ {
			score := lineOfSight(grid, row, column, -1, 0)
			score *= lineOfSight(grid, row, column, 1, 0)
			score *= lineOfSight(grid, row, column, 0, -1)
			score *= lineOfSight(grid, row, column, 0, 1)

			if score > sum {
				sum = score
			}
		}
	}

	return sum
}

func lineOfSight(grid [][]int, row, column, dr, dc int) int {
	sum := 0
	startingHeight := grid[row][column]
	row += dr
	column += dc
	for row >= 0 && row < len(grid) && column >= 0 && column < len(grid[0]) {
		height := grid[row][column]
		if height < startingHeight {
			sum++
		} else {
			sum++
			break
		}
		row += dr
		column += dc
	}

	return sum
}

func parseInput() (grid [][]int) {
	for _, line := range util.ReadFile("day8/input.txt") {
		var row []int
		for _, n := range strings.Split(line, "") {
			row = append(row, util.ParseInt(n))
		}
		grid = append(grid, row)
	}
	return grid
}
