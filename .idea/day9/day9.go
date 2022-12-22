package day9

import (
	"adventofcode-go-2022/util"
)

type instructions struct {
	direction string
	steps     int
}

func parseInstructions() []instructions {
	var i []instructions
	for _, line := range util.ReadFile("day9/input.txt") {
		i = append(i, instructions{
			direction: line[:1],
			steps:     util.ParseInt(line[2:]),
		})
	}
	return i
}

type node struct {
	grid [2]int
	next *node
}

type rope struct {
	head, tail *node
}

func (r rope) moveOneStep(direction string) {
	moves := map[string][2]int{
		"R": {0, 1},
		"U": {1, 0},
		"L": {0, -1},
		"D": {-1, 0},
	}

	diff := moves[direction]

	r.head.grid[0] += diff[0]
	r.head.grid[1] += diff[1]

	r.head.moveTail()
}

func (n *node) moveTail() {
	if n.next == nil {
		return
	}

	rowSteps := n.grid[0] - n.next.grid[0]
	colSteps := n.grid[1] - n.next.grid[1]

	if util.AbsInt(rowSteps) > 1 && util.AbsInt(colSteps) > 1 {
		n.next.grid[0] += rowSteps / 2
		n.next.grid[1] += colSteps / 2
	} else if util.AbsInt(rowSteps) > 1 {
		n.next.grid[0] += rowSteps / 2
		n.next.grid[1] += colSteps
	} else if util.AbsInt(colSteps) > 1 {
		n.next.grid[0] += rowSteps
		n.next.grid[1] += colSteps / 2
	} else {
		return
	}

	n.next.moveTail()
}

func newRope(length int) rope {

	head := &node{}
	itr := head

	for i := 1; i < length; i++ {
		itr.next = &node{}
		itr = itr.next
	}

	return rope{
		head: head,
		tail: itr,
	}
}

func Part1() int {
	inst := parseInstructions()

	r := newRope(2)
	visited := map[[2]int]bool{}

	for _, i := range inst {
		for i.steps > 0 {
			r.moveOneStep(i.direction)
			visited[r.tail.grid] = true
			i.steps--
		}
	}
	return len(visited)
}

func Part2() int {
	inst := parseInstructions()

	r := newRope(10)
	visited := map[[2]int]bool{}

	for _, i := range inst {
		for i.steps > 0 {
			r.moveOneStep(i.direction)
			visited[r.tail.grid] = true
			i.steps--
		}
	}
	return len(visited)
}
