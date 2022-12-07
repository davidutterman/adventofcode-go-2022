package day5

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"Final box letters",
			"HBTMTBSDC",
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
		want string
	}{
		{
			"Final box letters, moving in stacks",
			"PQTJRSHWS",
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
