package sem7

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_dfs(t *testing.T) {
	type args struct {
		graph   map[int][]int
		v       int
		visited map[int]int
		color   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dfs(tt.args.graph, tt.args.v, tt.args.visited, tt.args.color)
		})
	}
}

func Test_findConnectedComponents(t *testing.T) {
	type args struct {
		graph map[int][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Positive #1",
			args: args{
				map[int][]int{
					1:  {2, 3},
					2:  {1, 3},
					3:  {1, 2},
					4:  {6, 7},
					5:  {6, 7},
					6:  {4, 5, 7},
					7:  {4, 5, 6},
					8:  {11},
					9:  {10, 11},
					10: {9},
					11: {8, 9},
				},
			},
			want: [][]int{{1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findConnectedComponents(tt.args.graph)
			fmt.Println(got)

			require.Equal(t, len(tt.want), len(got))
		})
	}
}
