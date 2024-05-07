package sem_5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive #1",
			args: args{
				root: BuildTree([]int{11, 8, 16, 2, 9, 6, Poison, 7, Poison, Poison, 9, Poison, Poison, Poison, Poison}, 0),
			},
			want: 3,
		},
		{
			name: "Positive #2",
			args: args{
				root: BuildTree([]int{11}, 0),
			},
			want: 1,
		},
		{
			name: "Positive #3",
			args: args{
				root: BuildTree([]int{11, 1, Poison}, 0),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := MinDepth(tt.args.root)

			require.Equal(t, tt.want, got)
		})
	}
}
