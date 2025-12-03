package day1_test

import (
	"advent-of-code-2025/day1"
	"testing"
)

func TestCountTimesPointing0(t *testing.T) {
	type args struct {
		startsAt int64
		moves    []string
	}
	{
		tests := []struct {
			name string
			args args
			want int64
		}{
			{
				name: "Sample test",
				args: args{
					startsAt: 50,
					moves:    []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"},
				},
				want: 3,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := day1.CountTimesPointing0(tt.args.startsAt, tt.args.moves)

				if got != tt.want {
					t.Errorf("CountTimesPointing0() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}


func TestCountTimesHoveringAt0(t *testing.T) {
	type args struct {
		startsAt int64
		moves    []string
	}
	{
		tests := []struct {
			name string
			args args
			want int64
		}{
			{
				name: "Sample test",
				args: args{
					startsAt: 50,
					moves:    []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"},
				},
				want: 6,
			},
			{
				name: "hovering 0 3 times and keeps same current position",
				args: args{
					startsAt: 50,
					moves:    []string{"L300"},
				},
				want: 3,
			},
			{
				name: "hovering 0 8 times and keeps same current position",
				args: args{
					startsAt: 50,
					moves:    []string{"L300", "R500"},
				},
				want: 8,
			},

			{
				name: "hovering 0 8 times and keeps same current position",
				args: args{
					startsAt: 50,
					moves:    []string{"R50", "R30", "R150", "L279", "R256", "L100"},
				},
				want: 7,
			},
			{
				name: "hovering 0 10 times and keeps same current position",
				args: args{
					startsAt: 50,
					moves:    []string{"R1000"},
				},
				want: 10,
			},
			{
				name: "starting at 0 hovering 0 times when moving left",
				args: args{
					startsAt: 0,
					moves:    []string{"L5"},
				},
				want: 0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := day1.CountTimesHoveringAt0(tt.args.startsAt, tt.args.moves)

				if got != tt.want {
					t.Errorf("CountTimesHoveringAt0() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
