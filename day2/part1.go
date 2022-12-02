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

const (
	Lose = 0
	Draw = 3
	Win  = 6
)

const (
	Rock    = 1
	Paper   = 2
	Scissor = 3
)

/*
Opponent		Response (points)	Results (points)
A = Rock		X = Rock(1)			Lose(0)
B = Paper		Y = Paper(2)		Draw(3)
C = Scissors	Z = Scissors(3)		Win(6)
*/
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
