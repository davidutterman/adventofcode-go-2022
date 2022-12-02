package main

import (
	"adventofcode-go-2022/util"
	"fmt"
)

func main() {
	lines := util.ReadFile("day2/input.txt")
	results := parseLines(lines)
	fmt.Println(results)
}

func parseLines(lines []string) []game {
	var games []game
	sumOfGame := 0
	for _, line := range lines {
		thisRound := game{
			result: line,
			points: pointsAwarded[line],
		}
		games = append(games, thisRound)
		sumOfGame += thisRound.points
	}
	fmt.Printf("Your points: %d after %d games.\n", sumOfGame, len(games))
	return games
}

type game struct {
	result string
	points int
}

/*
Opponent, (score)
A = Rock (1)
B = Paper (2)
C = Scissors (3)

Response ?
X = ?
Y = ?
Z = ?
*/
var pointsAwarded = map[string]int{
	"A X": 3, // Need to loose(0) Rock vs. Scissors(3)
	"B X": 1, // Need to loose(0) Paper vs. Rock(1)
	"C X": 2, // Need to loose(0) Scissors vs. Paper(2)

	"A Y": 4, // Need to draw(3) Rock vs. Rock(1)
	"B Y": 5, // Need to draw(3) Paper vs. Paper(2)
	"C Y": 6, // Need to draw(3) Scissors vs. Scissors(3)

	"A Z": 8, // Need to win(6) Rock vs. Paper(2)
	"B Z": 9, // Need to win(6) Paper vs. Scissors(3)
	"C Z": 7, // Need to win(6) Scissors vs. Rock(1)
}
