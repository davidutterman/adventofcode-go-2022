package day6

import "adventofcode-go-2022/util"

func getFirstLine() string {
	lines := util.ReadFile("day6/input.txt")
	return lines[0]
}

func Part1() int {
	l := getFirstLine()
	for i := 0; i+4 <= len(l); i++ {
		if findDiff(l[i : i+4]) {
			return i + 4
		}
	}
	return -1
}

func findDiff(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func Part2() int {
	l := getFirstLine()
	for i := 0; i+14 <= len(l); i++ {
		if findDiff(l[i : i+14]) {
			return i + 14
		}
	}
	return -1
}
