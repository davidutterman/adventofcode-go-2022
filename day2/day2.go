package day2

import (
	"adventofcode-go-2022/util"
)

const inputFile = "day2/input.txt"

const ( // Response played
	Rock    = 1 // X
	Paper   = 2 // Y
	Scissor = 3 // Z
)

const ( // Points per game
	Lose = 0
	Draw = 3
	Win  = 6
)

// Opponent A = Rock, B = Paper, C = Scissors
func Part1() int {
	var pointMap = map[string]int{
		"A X": Rock + Draw,    // Rock vs. rock
		"A Y": Paper + Win,    // Rock vs. paper
		"A Z": Scissor + Lose, // Rock vs. scissors

		"B X": Rock + Lose,   // Paper vs. rock
		"B Y": Paper + Draw,  // Paper vs. paper
		"B Z": Scissor + Win, // Paper vs. scissors

		"C X": Rock + Win,     // Scissors vs. rock
		"C Y": Paper + Lose,   // Scissors vs. paper
		"C Z": Scissor + Draw, // Scissors vs. scissors
	}
	return sumOfGames(util.ReadFile(inputFile), pointMap)
}

func Part2() int {
	var pointMap = map[string]int{
		// X needs to Lose
		"A X": Scissor + Lose, // Rock vs. scissors
		"B X": Rock + Lose,    // Paper vs. rock
		"C X": Paper + Lose,   // Scissors vs. paper

		// Y needs to Draw
		"A Y": Rock + Draw,    // Rock vs. rock
		"B Y": Paper + Draw,   // Paper vs. paper
		"C Y": Scissor + Draw, // Scissors vs. scissors

		//Z needs to Win
		"A Z": Paper + Win,   // Rock vs. paper
		"B Z": Scissor + Win, // Paper vs. scissors
		"C Z": Rock + Win,    // Scissors vs. rock
	}
	return sumOfGames(util.ReadFile(inputFile), pointMap)
}

func sumOfGames(lines []string, points map[string]int) int {
	sum := 0
	for _, line := range lines {
		sum += points[line]
	}
	return sum
}
