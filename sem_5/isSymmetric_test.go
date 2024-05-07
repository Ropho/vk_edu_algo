package sem_5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSymmetricBFS(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				root: BuildTree([]int{3, 8, 8, 9, 6, 6, 9}, 0),
			},
			want: true,
		},
		{
			name: "Negative #1",
			args: args{
				root: BuildTree([]int{3, 7, 8, 666, 9, 6, 9}, 0),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			PrintTreePreOrder(tt.args.root)
			got := IsSymmetricBFS(tt.args.root)

			require.Equal(t, got, tt.want)
		})
	}
}

func TestIsSymmetricDFS(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive #1",
			args: args{
				root: BuildTree([]int{3, 8, 8, 9, 6, 6, 9}, 0),
			},
			want: true,
		},
		{
			name: "Negative #1",
			args: args{
				root: BuildTree([]int{3, 7, 8, 666, 9, 6, 9}, 0),
			},
			want: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := IsSymmetricDFS(tt.args.root)

			require.Equal(t, tt.want, got)
		})
	}
}
