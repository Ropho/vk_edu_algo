package sem_6

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"vk_algo/sem_5"
)

func Test_dfs(t *testing.T) {
	type args struct {
		node *sem_5.Node
	}
	tests := []struct {
		name string
		args args
		want *dfsResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMaxPathSum(t *testing.T) {
	type args struct {
		root *sem_5.Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive #1",
			args: args{sem_5.BuildTree(
				[]int{8, 3, 2, 9, 4, sem_5.Poison, sem_5.Poison, sem_5.Poison, sem_5.Poison, sem_5.Poison, 11}, 0,
			)},
			want: []int{11, 4, 3, 8, 2},
		},
		{
			name: "Positive #2",
			args: args{sem_5.BuildTree(
				[]int{8, 3, 2, 9, 4, sem_5.Poison, sem_5.Poison, sem_5.Poison, 7, sem_5.Poison, 11, sem_5.Poison, sem_5.Poison, sem_5.Poison, sem_5.Poison}, 0,
			)},
			want: []int{7, 9, 3, 11, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sem_5.PrintTreePreOrder(tt.args.root)
			got := FindMaxPathSum(tt.args.root)

			require.Equal(t, tt.want, got)
		})
	}
}
