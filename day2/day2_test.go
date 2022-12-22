package day2

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Points of all played games",
			want: 12645,
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
			name: "Sum of all games played with secret strategy",
			want: 11756,
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
