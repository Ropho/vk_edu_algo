package sem_5

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	type args struct {
		arr   []int
		index int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Positive #1",
			args: args{
				arr:   ([]int{8, 9, 11, 7, 16, 3, 1}),
				index: 0,
			},
		},
		{
			name: "Positive #2",
			args: args{
				arr:   ([]int{11, 8, 16, 2, 9, 6, 666, 7, 666, 666, 9, 666, 666, 666, 666}),
				index: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := BuildTree(tt.args.arr, tt.args.index)

			PrintTreePreOrder(got)
		})
	}
}
