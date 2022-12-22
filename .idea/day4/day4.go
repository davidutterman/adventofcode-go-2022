package day4

import (
	"adventofcode-go-2022/util"
	"strconv"
	"strings"
)

type section struct {
	min int
	max int
}

type elfPair struct {
	left  section
	right section
}

func Part1() int {
	sum := 0
	for _, p := range buildData(util.ReadFile("day4/input.txt")) {
		if fullOverlap(p) {
			sum++
		}
	}
	return sum
}

func Part2() int {
	sum := 0
	for _, p := range buildData(util.ReadFile("day4/input.txt")) {
		if partialOverlap(p) {
			sum++
		}
	}
	return sum
}

func partialOverlap(p elfPair) bool {
	if (p.right.min >= p.left.min && p.right.min <= p.left.max) ||
		(p.left.min >= p.right.min && p.left.min <= p.right.max) {
		return true
	}
	return false
}

func fullOverlap(p elfPair) bool {
	if (p.left.min >= p.right.min && p.left.max <= p.right.max) ||
		(p.right.min >= p.left.min && p.right.max <= p.left.max) {
		return true
	}
	return false
}

func buildData(lines []string) []elfPair {
	var data []elfPair
	for _, l := range lines {
		pair := strings.Split(l, ",")
		firstMin, firstMax := parseMinMax(pair[0])
		secondMin, secondMax := parseMinMax(pair[1])

		data = append(data, elfPair{
			left: section{
				min: firstMin,
				max: firstMax,
			},
			right: section{
				min: secondMin,
				max: secondMax,
			},
		})
	}
	return data
}

func parseMinMax(s string) (int, int) {
	minMax := strings.Split(s, "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])
	return min, max
}
