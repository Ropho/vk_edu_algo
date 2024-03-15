package sem_1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSortedArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Positive 1",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{6, 7, 8, 9, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "Positive 2",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{1, 2, 3, 4, 5, 5, 56},
			},
			want: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 5, 56},
		},
		{
			name: "Positive 3",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Positive 4",
			args: args{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: nil,
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Positive 5",
			args: args{
				arr1: []int{},
				arr2: []int{},
			},
			want: []int{},
		},
		{
			name: "Positive 6",
			args: args{
				arr1: nil,
				arr2: nil,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := MergeSortedArrays(tt.args.arr1, tt.args.arr2)

			require.Equal(t, tt.want, got)
		})
	}
}
