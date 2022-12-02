package main

import (
	"adventofcode-go-2022/util"
	"fmt"
)

func main() {
	parseLines(util.ReadFile("day2/input.txt"))
}

func parseLines(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += pointsAwarded[line]
	}
	fmt.Printf("Your points: %d\n", sum)
}

// Response (points)
const (
	Rock    = 1 // X
	Paper   = 2 // Y
	Scissor = 3 // Z
)

// Results (points)
const (
	Lose = 0
	Draw = 3
	Win  = 6
)

// Opponent A = Rock B = Paper C = Scissors
var pointsAwarded = map[string]int{
	"A X": Rock + Draw,
	"A Y": Paper + Win,
	"A Z": Scissor + Lose,

	"B X": Rock + Lose,
	"B Y": Paper + Draw,
	"B Z": Scissor + Win,

	"C X": Rock + Win,
	"C Y": Paper + Lose,
	"C Z": Scissor + Draw,
}
