package day10

import (
	"adventofcode-go-2022/util"
	"fmt"
	"strings"
)

func Part1() int {

	sum := 0
	registerx := 1
	cycle := 1
	for _, l := range util.ReadFile("day10/input.txt") {
		token := strings.Fields(l)
		cmd := token[0]
		addx := 0
		currentCycles := 0

		switch cmd {
		case "noop":
			currentCycles = 1
		case "addx":
			currentCycles = 2
			addx = util.ParseInt(token[1])
		}

		for i := 0; i < currentCycles; i++ {
			if cycle%40 == 20 {
				sum += registerx * cycle
			}
			cycle++
		}
		registerx += addx
	}

	return sum
}

const filledPixel = "#"
const emptyPixel = "."

func Part2() int {
	registerx := 1
	cycle := 1

	crt := initCrt()

	for _, l := range util.ReadFile("day10/input.txt") {
		token := strings.Fields(l)
		cmd := token[0]

		currentCycles := 0
		addx := 0

		switch cmd {
		case "noop":
			currentCycles = 1
		case "addx":
			currentCycles = 2
			addx = util.ParseInt(token[1])
		}

		for i := 0; i < currentCycles; i++ {

			pixelRow := (cycle - 1) / 40
			pixelCol := (cycle - 1) % 40
			left, right := registerx-1, registerx+1

			if left <= pixelCol && right >= pixelCol {
				crt[pixelRow][pixelCol] = filledPixel
			}

			cycle++
		}
		registerx += addx
	}

	printCrt(crt)
	return 0
}

func printCrt(crt [6][40]string) {
	log := ""
	for _, rows := range crt {
		for _, cell := range rows {
			log += cell
		}
		log += "\n"
	}
	fmt.Println(log)
}

func initCrt() [6][40]string {
	crt := [6][40]string{}
	for i, rows := range crt {
		for j := range rows {
			crt[i][j] = emptyPixel
		}
	}
	return crt
}
