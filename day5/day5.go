package day5

import (
	"adventofcode-go-2022/util"
	"strings"
)

func Part1() string {
	piles := moveCrates(parseInput())
	return readTopCrates(piles)
}

func Part2() string {
	piles := moveMultipleCrates(parseInput())
	return readTopCrates(piles)
}

func parseInput() ([][]string, []instruction) {
	lines := util.ReadFile("day5/input.txt")
	return parsePiles(lines[:8]), parseInstructions(lines[10:])
}

func readTopCrates(piles [][]string) string {
	var labels string
	for _, pile := range piles {
		labels += pile[len(pile)-1]
	}
	return labels
}

func moveMultipleCrates(piles [][]string, instructions []instruction) [][]string {
	for _, i := range instructions {
		fromPile := piles[i.indexFrom]

		indexCrate := len(fromPile) - i.numberOfCrates

		piles[i.indexFrom] = fromPile[:indexCrate]
		piles[i.indexTo] = append(piles[i.indexTo], fromPile[indexCrate:]...)
	}
	return piles
}

func moveCrates(piles [][]string, instructions []instruction) [][]string {
	for _, i := range instructions {
		for j := i.numberOfCrates; j > 0; j-- {
			fromPile := piles[i.indexFrom]
			crateIndex := len(fromPile) - 1
			var crate string
			crate, piles[i.indexFrom] = fromPile[crateIndex], fromPile[:crateIndex]
			piles[i.indexTo] = append(piles[i.indexTo], crate)
		}
	}
	return piles
}

type instruction struct {
	numberOfCrates int
	indexFrom      int
	indexTo        int
}

func parseInstructions(lines []string) []instruction {
	var i []instruction
	for _, line := range lines {
		tokens := strings.Fields(line)
		i = append(i, instruction{
			util.ParseInt(tokens[1]),
			util.ParseInt(tokens[3]) - 1,
			util.ParseInt(tokens[5]) - 1,
		})
	}
	return i
}

func parsePiles(lines []string) [][]string {
	piles := make([][]string, 9)
	for i := 7; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			column := j*4 + 1
			crate := string(lines[i][column])
			if crate != " " {
				piles[j] = append(piles[j], crate)
			}
		}
	}
	return piles
}
