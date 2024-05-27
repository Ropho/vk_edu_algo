package sem7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		graph map[int][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				graph: map[int][]int{
					1: {3},
					3: {1, 5},
					5: {3, 6, 8},
					6: {5, 7},
					7: {6, 8},
					8: {5, 7},
				},
			},
			want: true,
		},
		{
			name: "Negative #1",
			args: args{
				graph: map[int][]int{
					1: {3},
					3: {1, 5},
					5: {3, 6},
					6: {5, 7},
					7: {6, 8},
					8: {7},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasCycle(tt.args.graph)

			require.Equal(t, tt.want, got)
		})
	}
}

func Test_dfsCycle(t *testing.T) {
	type args struct {
		graph   map[int][]int
		vertex  int
		parent  int
		visited map[int]bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfsCycle(tt.args.graph, tt.args.vertex, tt.args.parent, tt.args.visited); got != tt.want {
				t.Errorf("dfsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
