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
Opponent
A = Rock
B = Paper
C = Scissors

Response, (score)
X = Rock (1)
Y = Paper (2)
Z = Scissors (3)
*/
var pointsAwarded = map[string]int{
	"A X": 4, // Rock vs. rock(1) draw(3)
	"A Y": 8, // Rock vs. paper(2) + win(6)
	"A Z": 3, // Rock vs. scissors(3) + lose(0)

	"B X": 1, // Paper vs. rock(1) + lose(0)
	"B Y": 5, // Paper vs. paper(2) + draw(3)
	"B Z": 9, // Paper vs. scissors(3) + win(6)

	"C X": 7, // Scissors vs. rock(1) + win(6)
	"C Y": 2, // Scissors vs. paper(2) + lose(0)
	"C Z": 6, // Scissors vs. scissors(3) + draw(3)
}
