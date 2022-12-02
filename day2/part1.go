package main

import (
	"adventofcode-go-2022/util"
	"fmt"
)

func main() {
	points := parseLines(util.ReadFile("day2/input.txt"))
	fmt.Printf("Your points: %d\n", points)
}

func parseLines(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += pointsAwarded[line]
	}
	return sum
}

/*
Opponent		Response (points)	Results (points)
A = Rock		X = Rock(1)			Lose(0)
B = Paper		Y = Paper(2)		Draw(3)
C = Scissors	Z = Scissors(3)		Win(6)
*/
var pointsAwarded = map[string]int{
	"A X": 4, // Rock vs. Rock(1) + draw(3)
	"A Y": 8, // Rock vs. Paper(2) + win(6)
	"A Z": 3, // Rock vs. Scissors(3) + lose(0)

	"B X": 1, // Paper vs. Rock(1) + lose(0)
	"B Y": 5, // Paper vs. Paper(2) + draw(3)
	"B Z": 9, // Paper vs. Scissors(3) + win(6)

	"C X": 7, // Scissors vs. Rock(1) + win(6)
	"C Y": 2, // Scissors vs. Paper(2) + lose(0)
	"C Z": 6, // Scissors vs. Scissors(3) + draw(3)
}
