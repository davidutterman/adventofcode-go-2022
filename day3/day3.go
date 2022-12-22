package day3

import (
	"adventofcode-go-2022/util"
	"fmt"
)

func Part1() {
	lines := util.ReadFile("day3/input-ex.txt")

	for _, line := range lines {
		halfIndex := len(line) / 2
		var first, second = line[halfIndex:], line[:halfIndex]

		fmt.Printf("f:%s s:%s\n", first, second)

		var twin string

		for _, fk := range first {
			for _, sk := range second {
				if fk == sk {
					twin = string(fk)
				}
			}
		}
		fmt.Println(twin)
	}
}
