package day1

import (
	"adventofcode-go-2022/util"
	"sort"
	"strconv"
)

var inputFile = "day1/calories.txt"

func Part1() int {
	elves := parseLines(util.ReadFile(inputFile))
	return findElfWithMostCalories(elves)
}

func Part2() int {
	elves := parseLines(util.ReadFile(inputFile))
	sortBySumOfCalories(elves)
	return sumTopThree(elves)
}

type elf struct {
	calories int
}

func parseLines(lines []string) []elf {
	var allElves []elf
	var myElf elf
	for _, line := range lines {

		if line == "" {
			allElves = append(allElves, myElf)
			myElf = elf{}
			continue
		}

		calories, _ := strconv.Atoi(line)
		myElf.calories += calories
	}
	return allElves
}

func sortBySumOfCalories(elves []elf) {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
}

func sumTopThree(elves []elf) int {
	sum := 0
	for i := 0; i <= 2; i++ {
		sum += elves[i].calories
	}
	return sum
}

func findElfWithMostCalories(elves []elf) int {
	maxElf := elf{}
	for _, myElf := range elves {
		if maxElf.calories < myElf.calories {
			maxElf = myElf
		}
	}
	return maxElf.calories
}
