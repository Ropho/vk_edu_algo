package sem7

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsTree(t *testing.T) {
	type args struct {
		graph map[int][]int
		start int
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
					2: {1},
					1: {2, 3, 4},
					3: {1},
					4: {1, 5, 6},
					5: {4},
					6: {4},
				},
				start: 1,
			},
			want: true,
		},
		{
			name: "Negative #1",
			args: args{
				graph: map[int][]int{
					2: {1},
					1: {2, 3, 4, 6},
					3: {1},
					4: {5, 6, 1},
					5: {4},
					6: {4, 1},
				},
				start: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("TEST")
			got := IsTree(tt.args.graph, tt.args.start)

			require.Equal(t, tt.want, got)
		})
	}
}
