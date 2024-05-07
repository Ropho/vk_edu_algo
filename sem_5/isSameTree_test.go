package sem_5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSameTree(t *testing.T) {
	type args struct {
		first  *Node
		second *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				first:  BuildTree([]int{8, 9, 11, 7, 16, 3, 1}, 0),
				second: BuildTree([]int{8, 9, 11, 7, 16, 3, 1}, 0),
			},
			want: true,
		},

		{
			name: "Positive #2",
			args: args{
				first:  BuildTree([]int{8, 9, 11, 7, 16, 3, 2}, 0),
				second: BuildTree([]int{8, 9, 11, 7}, 0),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSameTree(tt.args.first, tt.args.second)

			require.Equal(t, got, tt.want)
		})
	}
}
