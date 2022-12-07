package day5

import (
	"adventofcode-go-2022/util"
	"strconv"
	"strings"
)

type instruction struct {
	nrOfBoxes int
	from      int
	to        int
}

func Part1() string {
	lines := util.ReadFile("day5/input.txt")

	piles := parsePiles(lines[:8])
	instructions := parseInstructions(lines[10:])

	piles = moveBoxes(piles, instructions)

	return getTopBoxes(piles)
}

func Part2() string {
	lines := util.ReadFile("day5/input.txt")

	piles := parsePiles(lines[:8])
	instructions := parseInstructions(lines[10:])

	piles = moveMultipleBoxes(piles, instructions)

	return getTopBoxes(piles)

}

func getTopBoxes(piles [][]string) string {
	var result string
	for _, pile := range piles {
		result += pile[len(pile)-1]
	}
	return result
}

func moveMultipleBoxes(piles [][]string, instructions []instruction) [][]string {
	for _, inst := range instructions {

		fromPile := piles[inst.from-1]
		nrFrom := len(fromPile)

		boxesToMove := fromPile[nrFrom-inst.nrOfBoxes:]

		piles[inst.from-1] = fromPile[:nrFrom-inst.nrOfBoxes]

		for _, c := range boxesToMove {
			piles[inst.to-1] = append(piles[inst.to-1], c)
		}
	}
	return piles
}

func moveBoxes(piles [][]string, instructions []instruction) [][]string {
	for _, inst := range instructions {
		for i := inst.nrOfBoxes; i > 0; i-- {
			fromPile := piles[inst.from-1]
			nrFrom := len(fromPile) - 1
			var crate string
			crate, piles[inst.from-1] = fromPile[nrFrom], fromPile[:nrFrom]
			piles[inst.to-1] = append(piles[inst.to-1], crate)
		}
	}
	return piles
}

func parseInstructions(lines []string) []instruction {
	var i []instruction

	for _, line := range lines {
		tokens := strings.Fields(line)
		i = append(i, instruction{
			parseInt(tokens[1]),
			parseInt(tokens[3]),
			parseInt(tokens[5]),
		})
	}
	return i
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
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
