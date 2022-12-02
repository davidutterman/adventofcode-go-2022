package main

import (
	"adventofcode-go-2022/util"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	lines := util.ReadFile("day1/calories.txt")
	elves := parseLines(lines)

	findElfWithMostCalories(elves)
	sortBySumOfCalories(elves)
	topThree(elves)
}

type elf struct {
	packedCalories []int
	sumOfCalories  int
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

		calories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("error converting %s to int", line)
			continue
		}
		myElf.packedCalories = append(myElf.packedCalories, calories)
		myElf.sumOfCalories += calories
	}
	return allElves
}

func sortBySumOfCalories(elves []elf) {
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].sumOfCalories > elves[j].sumOfCalories
	})
}

func topThree(elves []elf) {
	sumCalories := 0
	for i := 0; i <= 2; i++ {
		sumCalories += elves[i].sumOfCalories
	}
	fmt.Printf("Sum of top three elves packs: %v\n", sumCalories)
}

func findElfWithMostCalories(elves []elf) {
	maxElf := elf{}
	for _, myElf := range elves {
		if maxElf.sumOfCalories < myElf.sumOfCalories {
			maxElf = myElf
		}
	}
	fmt.Printf("Elf with most calories: %v\n", maxElf)
}
