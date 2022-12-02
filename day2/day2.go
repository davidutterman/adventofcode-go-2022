package day2

import (
	"adventofcode-go-2022/util"
)

const ( // Response played
	Rock    = 1 // X
	Paper   = 2 // Y
	Scissor = 3 // Z
)

const ( // Results
	Lose = 0
	Draw = 3
	Win  = 6
)

// Opponent A = Rock B = Paper C = Scissors
func Part1() int {
	var points = map[string]int{
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
	return sum(getPlayedHands(), points)
}

func Part2() int {
	var points = map[string]int{
		// X needs to Lose
		"A X": Scissor + Lose, // Rock vs. Scissors
		"B X": Rock + Lose,    // Paper vs. Rock
		"C X": Paper + Lose,   // Scissors vs. Paper

		// Y needs to Draw
		"A Y": Rock + Draw,    // Rock vs. Rock
		"B Y": Paper + Draw,   // Paper vs. Paper
		"C Y": Scissor + Draw, // Scissors vs. Scissors

		//Z needs to Win
		"A Z": Paper + Win,   // Rock vs. Paper
		"B Z": Scissor + Win, // Paper vs. Scissors
		"C Z": Rock + Win,    // Scissors vs. Rock
	}
	return sum(getPlayedHands(), points)
}

func getPlayedHands() []string {
	return util.ReadFile("day2/input.txt")
}

func sum(lines []string, awarded map[string]int) int {
	sum := 0
	for _, line := range lines {
		sum += awarded[line]
	}
	return sum
}
