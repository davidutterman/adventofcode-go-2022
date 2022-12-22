package main

import (
	"adventofcode-go-2022/day1"
	"adventofcode-go-2022/day2"
	"adventofcode-go-2022/day3"
	"adventofcode-go-2022/day4"
	"adventofcode-go-2022/day5"
	"adventofcode-go-2022/day6"
	"adventofcode-go-2022/day7"
	"adventofcode-go-2022/day8"
	"adventofcode-go-2022/day9"
	"fmt"
)

func main() {
	runDay7()
}

func runDay9() {
	fmt.Println("Part 1:", day9.Part1())
	fmt.Println("Part 2:", day9.Part2())
}
func runDay8() {
	fmt.Println("Part 1:", day8.Part1())
	fmt.Println("Part 2:", day8.Part2())
}
func runDay7() {
	fmt.Println("Part 1:", day7.Part1())
	fmt.Println("Part 2:", day7.Part2())
}
func runDay6() {
	fmt.Printf("Part1 %v\n", day6.Part1())
	fmt.Printf("Part2 %v\n", day6.Part2())
}
func runDay5() {
	fmt.Printf("After moving all boxes the top ones are: %s\n", day5.Part1())
	fmt.Printf("After moving all boxes in chunks the top ones are: %s\n", day5.Part2())
}
func runDay4() {
	fmt.Printf("Number of overlaping elfpairs: %d\n", day4.Part1())
	fmt.Printf("Number of some overlaping elfpairs: %d\n", day4.Part2())
}
func runDay3() {
	day3.Part1()
}
func runDay2() {
	fmt.Printf("Your points (part1): %d\n", day2.Part1())
	fmt.Printf("Your points (part2): %d\n", day2.Part2())
}
func runDay1() {
	fmt.Printf("Part 1 - Elf with most calories: %v\n", day1.Part1())
	fmt.Printf("Part 2 - Sum of top three elves packs: %v\n", day1.Part2())
}
