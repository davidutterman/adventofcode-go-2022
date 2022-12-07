package main

import (
	"adventofcode-go-2022/day1"
	"adventofcode-go-2022/day2"
	"adventofcode-go-2022/day5"
	"fmt"
)

func main() {
	//runDay1()
	//runDay2()
	runDay5()
}

func runDay1() {
	fmt.Println("Day 1:")
	fmt.Printf("Part 1 - Elf with most calories: %v\n", day1.Part1())
	fmt.Printf("Part 2 - Sum of top three elves packs: %v\n", day1.Part2())
}

func runDay2() {
	fmt.Println("Day 2:")
	fmt.Printf("Your points (part1): %d\n", day2.Part1())
	fmt.Printf("Your points (part2): %d\n", day2.Part2())
}

func runDay5() {
	fmt.Println("Day 5:")
	fmt.Printf("After moving all boxes the top ones are: %s\n", day5.Part1())
	fmt.Printf("After moving all boxes in chunks the top ones are: %s\n", day5.Part2())
}
