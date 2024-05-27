package sem7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBellmanFord(t *testing.T) {
	type args struct {
		graph map[string]map[string]int
		start string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Positive #1",
			args: args{

				graph: map[string]map[string]int{
					"0": {
						"1": 4, "2": 1,
					},
					"1": {
						"3": 1,
					},
					"2": {
						"1": 2, "3": 5,
					},
					"3": {
						"4": 3,
					},
					"4": {},
				},
				start: "0",
			},
			want: map[string]int{
				"0": 0,
				"1": 3,
				"2": 1,
				"3": 4,
				"4": 7,
			},
		},
		{
			name: "Positive #2",
			args: args{
				graph: map[string]map[string]int{
					"A": {
						"B": 1, "C": 5,
					},
					"B": {
						"A": 1, "C": 2, "D": 3,
					},
					"C": {
						"A": 5, "B": 2, "D": 1,
					},
					"D": {
						"B": 3, "C": 1,
					},
				},
				start: "A",
			},
			want: map[string]int{
				"A": 0,
				"B": 1,
				"C": 3,
				"D": 4,
			},
		},
		{
			name: "Negative cycle #1",
			args: args{
				graph: map[string]map[string]int{
					"0": {
						"1": 5,
					},
					"1": {
						"6": 60,
						"5": 30,
						"2": 20,
					},
					"2": {
						"3": 10,
						"4": 75,
					},
					"3": {
						"2": -15,
					},
					"4": {
						"9": 100,
					},
					"5": {
						"4": 25,
						"8": 50,
						"6": 5,
					},
					"6": {
						"7": -50,
					},
					"7": {
						"8": -10,
					},
					"8": {},
					"9": {},
				},
				start: "0",
			},
			want: map[string]int{
				"0": 0,
				"1": 5,
				"2": -inf,
				"3": -inf,
				"4": -inf,
				"5": 35,
				"6": 40,
				"7": -10,
				"8": -20,
				"9": -inf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BellmanFord(tt.args.graph, tt.args.start)

			require.Equal(t, tt.want, got)
		})
	}
}
