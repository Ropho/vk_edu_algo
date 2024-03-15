package sem_1

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeArrInPlace(t *testing.T) {
	type args struct {
		arr1        []int
		arr2        []int
		lengthFirst int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name: "Postitive 1",
			args: args{
				arr1:        make([]int, 12),
				arr2:        []int{1, 2, 3, 4, 5, 6},
				lengthFirst: 5,
			},
			want:    []int{0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 0},
			wantErr: nil,
		},
		{
			name: "Postitive 2",
			args: args{
				arr1:        []int{0, 1, 2, 3, 4, -1, -1, -1, -1, -1, -1},
				arr2:        []int{1, 2, 3, 4, 5},
				lengthFirst: 5,
			},
			want:    []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, -1},
			wantErr: nil,
		},
		{
			name: "Negative: not enough capacity",
			args: args{
				arr1:        make([]int, 5),
				arr2:        []int{1, 2, 3, 4, 5, 6},
				lengthFirst: 5,
			},
			want:    nil,
			wantErr: errors.New("validation failed for arr1 length"),
		},
		{
			name: "Negative: nil array2",
			args: args{
				arr1:        make([]int, 5),
				arr2:        nil,
				lengthFirst: 0,
			},
			want:    nil,
			wantErr: errors.New("validation failed for length of array 2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MergeArrInPlace(tt.args.arr1, tt.args.arr2, tt.args.lengthFirst)

			require.Equal(t, tt.wantErr, err)
			require.Equal(t, tt.want, got)
		})
	}
}
