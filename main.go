package main

import (
	"adventofcode-go-2022/day1"
	"adventofcode-go-2022/day2"
	"adventofcode-go-2022/day3"
	"adventofcode-go-2022/day4"
	"adventofcode-go-2022/day5"
	"adventofcode-go-2022/day6"
	"adventofcode-go-2022/day7"
	"fmt"
)

func main() {
	runDay6()
}

func runDay7() {
	fmt.Println("Day 7:")
	day7.Part1()
}

func runDay6() {
	fmt.Println("Day 6:")
	fmt.Printf("Part1 %s\n", day6.Part1())
	fmt.Printf("Part2 %s\n", day6.Part2())
}
func runDay5() {
	fmt.Println("Day 5:")
	fmt.Printf("After moving all boxes the top ones are: %s\n", day5.Part1())
	fmt.Printf("After moving all boxes in chunks the top ones are: %s\n", day5.Part2())
}

func runDay4() {
	fmt.Println("Day 4:")
	fmt.Printf("Number of overlaping elfpairs: %d\n", day4.Part1())
	fmt.Printf("Number of some overlaping elfpairs: %d\n", day4.Part2())
}

func runDay3() {
	fmt.Println("Day 3:")
	day3.Part1()
}

func runDay2() {
	fmt.Println("Day 2:")
	fmt.Printf("Your points (part1): %d\n", day2.Part1())
	fmt.Printf("Your points (part2): %d\n", day2.Part2())
}

func runDay1() {
	fmt.Println("Day 1:")
	fmt.Printf("Part 1 - Elf with most calories: %v\n", day1.Part1())
	fmt.Printf("Part 2 - Sum of top three elves packs: %v\n", day1.Part2())
}
