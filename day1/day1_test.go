package day1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"Find the elf with most calories",
			66487,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			"What was the sum of the elves with the most calories?",
			197301,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
